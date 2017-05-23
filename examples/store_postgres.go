package main

import (
	"github.com/andrepinto/goway/store"
	"flag"
	"log"
)

func init() {
	flag.String("connectionString", "host=localhost user=postgres dbname=goway sslmode=disable", "--connectionString=\"host=localhost user=postgres dbname=goway sslmode=disable\"")
}

func main() {
	flag.Parse()

	repo := store.NewPostgresRepository(&store.PostgresRepositoryOptions{
		ConnectionString: flag.Lookup("connectionString").Value.String(),
	})

	for _, v := range repo.GetAllClients(){
		log.Println(v)
	}

	for _, v := range repo.GetAllProducts(){
		log.Println(v)
	}
}
