package main

import (
	"github.com/baoer/im_sys/router"
)

func main() {
	r := router.Router()

	r.Run(":8080")
	// str, _ := util.Gettoken("sdasdas", "dasd")
	// fmt.Println(str)
	// myclaims, _ := util.Parsetoken(str)
	// fmt.Println(myclaims)
}
