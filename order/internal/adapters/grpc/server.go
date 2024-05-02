package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/huseyinbabal/microservices-proto/golang/order"
	"github.com/psbrar99/go-grpc/order/internal/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	api  ports.APIPort
	port int
	order.UnimplementedOrderServer
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
	order.RegisterOrderServer(grpcserver, a)
}
