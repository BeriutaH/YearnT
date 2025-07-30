package main

import (
	"Yearn-go/config"
	"Yearn-go/routers"
)

func main() {
	config.InitDB()
	r := routers.SetupRouter()
	_ = r.Run(":8080")
}
