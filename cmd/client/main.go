package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"micro_services/api/v1/port"
	"micro_services/pkg/client/parser"
	"os"
)

func main() {
	dial, err := grpc.Dial(":4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer dial.Close()

	file, err := os.Open(os.Args[1])
	defer file.Close()
	if err != nil {
		panic(err)
	}

	client := port.NewPDServiceClient(dial)
	recorder, err := client.RecordPort(context.Background())
	if err != nil {
		panic(err)
	}

	ch := make(chan parser.Result)
	go parser.ReadPorts(file, ch)
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
		log.Println(err)
	}
	log.Println(res)
}
