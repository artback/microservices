package main

import (
	"context"
	"errors"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"io"
	"log"
	"micro_services/api/v1/api"
	"micro_services/api/v1/port"
	"micro_services/pkg/client/parser"
	"net/http"
	"os"
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
	s := Service{file}
	sum, err := s.Run(context.Background())
	if err != nil {
		log.Fatal(sum)
	}
	a := Api{}
	a.Run()
}

type Service struct {
	r io.Reader
}

func (s Service) Run(ctx context.Context) (*port.PortSummary, error) {
	dial, err := grpc.Dial(":4040", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	client := port.NewPDServiceClient(dial)

	recorder, err := client.RecordPort(ctx)
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

type Api struct {
	conn *grpc.ClientConn
}

func (a Api) Run() error {
	router := runtime.NewServeMux()
	err := api.RegisterPDServiceHandler(context.Background(), router, a.conn)
	if err != nil {
		return err
	}
	return http.ListenAndServe(":8080", router)
}
