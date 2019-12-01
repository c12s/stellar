package server

import (
	"fmt"
	sPb "github.com/c12s/scheme/stellar"
	"github.com/c12s/stellar/storage"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	db storage.DB
}

func (s *Server) List(ctx context.Context, req *sPb.ListReq) (*sPb.ListResp, error) {
	return s.db.List(ctx, req)
}

func (s *Server) Get(ctx context.Context, req *sPb.GetReq) (*sPb.GetResp, error) {
	return s.db.Get(ctx, req)
}

func Run(address string, db storage.DB) {
	lis, err := net.Listen("tcp", address)
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
