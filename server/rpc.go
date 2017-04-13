package server

import (
	"net"
	"fmt"
	"log"
	"google.golang.org/grpc"
	"github.com/andrepinto/goway/domain"
)

type RPCServer struct {
	port int
	server domain.GowayServer
}

func(rpc *RPCServer) Serve() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", rpc.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	domain.RegisterGowayServer(grpcServer, rpc.server)
	return grpcServer.Serve(lis)
}

func NewRPCServer(port int, server domain.GowayServer) *RPCServer {
	return &RPCServer{
		port,
		server,
	}
}
