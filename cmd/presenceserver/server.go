package main

import "github.com/abbasfisal/game-app/delivery/httpserver/grpcserver/presenceserver"

func main() {
	server := presenceserver.Server{}
	server.Start()
}
