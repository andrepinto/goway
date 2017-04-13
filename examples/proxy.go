package main

import (
	"fmt"
	"flag"
	"github.com/golang/glog"
	"github.com/andrepinto/goway/server"
)

var portREST = flag.Int("portREST", 8080, "The REST server port")

func main() {
	rest := server.NewRESTServer("localhost:10000", *portREST)
	err := rest.Serve()

	if err != nil {
		glog.Fatal(err)
	} else {
		fmt.Println("REST Server Listening on Port: %d", *portREST)
	}
}
