package grpc

import (
	"log"
	"net"

	"github.com/Nidal-Bakir/hex/internal/adapters/framwork/left/grpc/pb"
	"github.com/Nidal-Bakir/hex/internal/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.ApiPort
	pb.UnimplementedArithmaticServiceServer
}

func NewAdapter(api ports.ApiPort) *Adapter {
	return &Adapter{api: api}
}

func (ad Adapter) Run() {
	listen, err := net.Listen("tcp", "9000")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterArithmaticServiceServer(grpcServer, ad)
	if err = grpcServer.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
