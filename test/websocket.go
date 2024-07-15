package test

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var wsc = make(map[*websocket.Conn]struct{}, 0)

func HandlerConnecrtions(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer ws.Close()
	wsc[ws] = struct{}{}
	for {
		mt, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			break
		}

		for c := range wsc {
			err = c.WriteMessage(mt, msg)
			if err != nil {
				log.Println(err.Error())
				break
			}
		}

	}
}
