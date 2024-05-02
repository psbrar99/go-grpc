package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/psbrar99/go-grpc/order/config"
	"github.com/psbrar99/go-grpc/order/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api  ports.APIPort
	port int
	apis.RegisterOrderServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}

}

func (a Adapter) Run() {

	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":d", a.port))
	if err != nil {
		log.Fatalf("Unable to listen on port: %d, error: %v", a.port, err)
	}
	grpcserver := grpc.NewServer()
	apis.RegisterOrderServer(grpcserver, a)
	if config.GetEnv() == "dev" {
		reflection.Register(grpcserver)
	}
	if err := grpcserver.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc")
	}
}
