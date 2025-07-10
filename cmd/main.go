package main

import (
	"log"
	"reverse-proxy/internal/server"
)


func main(){

	if err:= server.Run();err != nil{
		log.Fatal("Error :: Could Not start the Server: %v",err)
	}
}
