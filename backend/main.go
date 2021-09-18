package main

import (
	"chatapp/router"
)

func main() {
	r := router.Router()

	// fmt.Println("starting server")
	r.Run()
}
