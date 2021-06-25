package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"io"
	"log"
	"micro_services/api/v1/api"
	"micro_services/api/v1/port"
	"micro_services/pkg/client"
	"micro_services/pkg/client/parser"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal(errors.New("missing file argument"))
	}
	var fileName = os.Args[1]
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	pdsHost, ok := os.LookupEnv("PDS_HOST")
	if !ok {
		log.Fatal(errors.New("PDS_HOST not set"))
	}

	dial, err := grpc.Dial(pdsHost, grpc.WithInsecure())
	client := port.NewPDServiceClient(dial)
	if err != nil {
		log.Fatal(err)
	}
	s := Import{client, file}
	_, err = s.Run(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	httpHost, ok := os.LookupEnv("HTTP_ADDRESS")
	if !ok {
		log.Fatal(errors.New("HTTP_ADDRESS not set"))
	}
	grpcHost, ok := os.LookupEnv("GRPC_ADDRESS")
	if !ok {
		log.Fatal(errors.New("GRPC_ADDRESS not set"))
	}
	go startGRPCServer(grpcHost, client)
	log.Fatal(RunHTTPServer(context.Background(), grpcHost, httpHost))
}

type Import struct {
	port.PDServiceClient
	r io.Reader
}

func (s Import) Run(ctx context.Context) (*port.PortSummary, error) {
	recorder, err := s.RecordPort(ctx)
	if err != nil {
		return nil, err
	}
	ch := make(chan parser.Result)
	go parser.ReadPorts(s.r, ch)
	for result := range ch {
		if result.Err != nil {
			log.Println(result.Err)
		}
		err := recorder.Send(result.Port)
		if err != nil {
			log.Println(err)
		}
	}
	res, err := recorder.CloseAndRecv()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func startGRPCServer(address string, serviceClient port.PDServiceClient) error {
	// create a listener on TCP port
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	// create a server instance

	c := client.ApiClient{serviceClient}
	// create a gRPC server object
	grpcServer := grpc.NewServer()
	// attach the Ping service to the server
	api.RegisterAPIServiceServer(grpcServer, &c)
	// start the server
	log.Printf("starting HTTP/2 gRPC server on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %s", err)
	}
	return nil
}
func RunHTTPServer(ctx context.Context, grpcHost, httpHost string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := api.RegisterAPIServiceHandlerFromEndpoint(ctx, mux, grpcHost, opts); err != nil {
		log.Fatalf("failed to start HTTP gateway: %v", err)
	}

	srv := &http.Server{
		Addr:    httpHost,
		Handler: mux,
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	log.Println("starting HTTP/REST gateway...")
	return srv.ListenAndServe()
}
