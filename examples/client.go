package main

import (
	"flag"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"github.com/andrepinto/goway/domain"
)

var serverAddr = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")

func main() {
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := domain.NewGowayClient(conn)


	grpclog.Printf("Getting Version")
	res, err := client.Version(context.Background(), &domain.VersionRequest{})
	if err != nil {
		grpclog.Fatalf("%v.Version(_) = _, %v: ", client, err)
	}
	grpclog.Println(res)
}
