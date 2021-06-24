package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"micro_services/api/v1/port"
	"micro_services/pkg/pds"
	"micro_services/pkg/pds/mongodb"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	port.RegisterPDServiceServer(srv, &pds.PortDomainService{mongodb.PortRepository{client}})

	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}
