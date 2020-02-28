package main

import (
	"gin-demo/router"
	"log"
)

func main() {
	router := router.InitRouter()
	var err = router.Run(":8000")
	if err != nil {
		log.Fatalln(err)
	}
}
