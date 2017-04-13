package main

import (
	"flag"
	"fmt"

	"github.com/joelbraga/aztek"
	"github.com/andrepinto/goway/action"
	"github.com/andrepinto/goway/domain"
	"github.com/andrepinto/goway/api"
	"github.com/andrepinto/goway/server"
	"github.com/andrepinto/goway/service"
)

var ch chan *action.Action
var portRPC = flag.Int("portRPC", 10000, "The RPC server port")

func init() {
	flag.String("connectionString", "host=localhost user=postgres dbname=goway sslmode=disable", "--connectionString=\"host=localhost user=postgres dbname=goway sslmode=disable\"")
}

func main() {

	flag.Parse()

	repo := aztek.NewPostgesCoreRepo(flag.Lookup("connectionString").Value.String())

	var interfaceSlice []interface{} = make([]interface{}, 4)
	interfaceSlice[0] = domain.Product{}
	interfaceSlice[1] = domain.Route{}
	interfaceSlice[2] = domain.Inject{}
	interfaceSlice[3] = domain.Client{}

	repo.Migrations(interfaceSlice)

	go getEvents()
	actionEvent := action.NewActionEvent()
	ch = actionEvent.Data

	apiResource := api.NewApiResource(&api.ApiOptions{
		repo,
		actionEvent,
	})

	rpc := server.NewRPCServer(*portRPC, service.NewGowayImpl(apiResource))
	rpc.Serve()
}

func getEvents() {
	for {
		select {
		case msg := <-ch:
			fmt.Println(msg.Type)
			fmt.Println(msg.Payload)
		}
	}
}
