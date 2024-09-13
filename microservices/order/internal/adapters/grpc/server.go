package grpc

import (
	"fmt"
	"net"

	"github.com/SergioVenicio/microservices/order/config"
	"github.com/SergioVenicio/microservices/order/internal/ports"
	"github.com/huseyinbabal/microservices-proto/golang/order"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api  ports.APIPort
	port int
	order.UnimplementedOrderServer
}

func (a Adapter) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	order.RegisterOrderServer(grpcServer, a)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	log.Printf("starting payment service on port %d ...", a.port)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to server grpc on port %d", a.port)
	}

}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}
