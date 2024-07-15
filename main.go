package main

import (
	"errors"
	"log"

	"github.com/baoer/im_sys/router"
)

func main() {
	r := router.Router()

	r.Run(":8080")
	// str, _ := util.Gettoken("sdasdas", "dasd")
	// fmt.Println(str)
	// myclaims, _ := util.Parsetoken(str)
	// fmt.Println(myclaims)
	// http.HandleFunc("/ws", test.HandlerConnecrtions)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	err := errors.New("something is wrong!!")
	log.Println(err)
}
