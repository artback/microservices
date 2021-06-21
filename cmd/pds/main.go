package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"micro_services/api/v1/port"
	"micro_services/pkg/pds"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()

	port.RegisterPDServiceServer(srv, &pds.PortDomainService{})

	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}
