package main

import (
	"github.com/baoer/im_sys/router"
)

func main() {
	r := router.Router()
	r.Run(":8080")

}
