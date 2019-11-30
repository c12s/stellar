package server

import (
	"fmt"
	sPb "github.com/c12s/scheme/stellar"
	"github.com/c12s/stellar/model"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct{}

func (s *Server) List(context.Context, *sPb.ListReq) (*sPb.ListResp, error) {
	return nil, nil
}

func (s *Server) Get(context.Context, *sPb.GetReq) (*sPb.GetResp, error) {
	return nil, nil
}

func Run(conf *model.Config) {
	lis, err := net.Listen("tcp", conf.Address)
	if err != nil {
		log.Fatalf("failed to initializa TCP listen: %v", err)
	}
	defer lis.Close()

	server := grpc.NewServer()
	stellarServer := &Server{}

	fmt.Println("StellarService RPC Started")
	sPb.RegisterStellarServiceServer(server, stellarServer)
	server.Serve(lis)
}
