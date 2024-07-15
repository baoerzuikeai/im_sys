package service

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/baoer/im_sys/models"
	"github.com/baoer/im_sys/util"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type PrivateMessage struct {
	NickName string `json:"nickname"`
	Uid      string `json:"uid"`
	Msg      string `json:"msg"`
	ToUserID string `json:"touserid"`
}

type Message struct {
	NickName     string `json:"nickname"`
	Uid          string `json:"uid"`
	Msg          string `json:"msg"`
	RoomIdentity string `bson:"room_identity"`
}

var upgrader = websocket.Upgrader{}
var wsc = make(map[string]*websocket.Conn, 0)
var mu sync.Mutex // 用于保护wsc的互斥锁

func WebSocketsendPrivateMessage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		wsn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer wsn.Close()

		myclaim := ctx.MustGet("user_claims").(*util.Myclaims)
		mu.Lock()
		wsc[myclaim.UserId] = wsn
		mu.Unlock()
		defer func() {
			mu.Lock()
			delete(wsc, myclaim.UserId)
			mu.Unlock()
		}()

		pmessage := new(PrivateMessage)
		for {
			err := wsn.ReadJSON(pmessage)
			if err != nil {
				log.Println(err.Error())
				break
			}

			mu.Lock()
			if v, ok := wsc[pmessage.ToUserID]; ok {
				pmessage.Uid = myclaim.UserId
				err = v.WriteJSON(pmessage)
				if err != nil {
					log.Println(err.Error())
					break
				}
				err = wsn.WriteJSON(pmessage)
				if err != nil {
					log.Println(err.Error())
					break
				}
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"code": -1,
					"msg":  "用户未在线",
				})
			}
			mu.Unlock()

		}
	}
}

func WebSocketsendChannelMessage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		wsn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer wsn.Close()

		myclaim := ctx.MustGet("user_claims").(*util.Myclaims)
		mu.Lock()
		wsc[myclaim.UserId] = wsn
		mu.Unlock()
		defer func() {
			mu.Lock()
			delete(wsc, myclaim.UserId)
			mu.Unlock()
		}()

		message := new(Message)
		for {
			err := wsn.ReadJSON(message)
			if err != nil {
				log.Println(err.Error())
				break
			}
			//查询同一房间的用户
			users, err := models.GetUsersByRoomIdentity(message.RoomIdentity)
			if err != nil {
				log.Println(err.Error())
				break
			}
			mu.Lock()
			for _, uid := range users {
				fmt.Println(uid)
				if v, ok := wsc[uid]; ok {
					message.Uid = myclaim.UserId
					err = v.WriteJSON(message)
					if err != nil {
						log.Println(err.Error())
						break
					}
				}
			}
			mu.Unlock()
		}
	}
}
