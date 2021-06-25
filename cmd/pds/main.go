package main

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"micro_services/api/v1/port"
	"micro_services/pkg/pds"
	"micro_services/pkg/pds/mongodb"
	"net"
	"os"
)

func main() {
	pdsHost, ok := os.LookupEnv("GRPC_ADDRESS")
	if !ok {
		log.Fatal(errors.New("GRPC_ADDRESS not set"))
	}
	listener, err := net.Listen("tcp", pdsHost)

	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	// Set client options
	mongodbHost, ok := os.LookupEnv("MONGO_DB_HOST")
	if !ok {
		log.Fatal(errors.New("MONGO_DB_HOST not set"))
	}
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s", mongodbHost))

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
