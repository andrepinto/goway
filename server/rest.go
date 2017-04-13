package server

import (
	"net/http"
	"fmt"

	"google.golang.org/grpc"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/andrepinto/goway/domain"
)

type RESTServer struct {
	rpcEndpoint string
	port int
}

func (srv *RESTServer) run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := domain.RegisterGowayHandlerFromEndpoint(ctx, mux, srv.rpcEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", srv.port), mux)
}

func (srv *RESTServer) Serve() error {
	defer glog.Flush()
	return srv.run()
}

func NewRESTServer(rpcEndpoint string, port int) *RESTServer {
	return &RESTServer{
		rpcEndpoint,
		port,
	}
}
