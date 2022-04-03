package main

import (
	server "HubbleWeb/server"
	service "HubbleWeb/service"
	"time"
)

func main() {
	go server.ApiServer()
	time.Sleep(1 * time.Second)
	service.UserInteraction()

}
