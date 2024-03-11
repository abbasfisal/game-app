package presenceserver

import (
	"fmt"
	"github.com/abbasfisal/game-app/contract/golang/presence"
	"github.com/abbasfisal/game-app/service/presenceservice"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	presence.UnimplementedPresenceServiceServer
	svc presenceservice.Service
}

func (s Server) Start() {
	listener, err := net.Listen("tcp", ":8086")
	if err != nil {
		panic(err)
	}

	//presence server
	presenceSvcServer := Server{}

	//start server
	grpcServer := grpc.NewServer()

	presence.RegisterPresenceServiceServer(grpcServer, &presenceSvcServer)

	fmt.Println("grpc server is started ")

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}

//func (s Server) GetPresence(ctx context.Context, req *presence.GetPresenceRequest) (*presence.GetPresenceResponse, error) {
////	s.svc.GetPresence(ctx)
//
//}
