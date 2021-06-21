package pds

import (
	"context"
	"io"
	"micro_services/api/v1/port"
	"micro_services/pkg/pds/repository"
	"time"
)

type PortDomainService struct {
	repository.PortRepository
}

func (p PortDomainService) RecordPort(server port.PDService_RecordPortServer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var portCount int32
	for {
		po, err := server.Recv()
		if err == io.EOF {
			return server.SendAndClose(&port.PortSummary{PortCount: portCount})
		}
		if err != nil {
			return err
		}
		err = p.UpsertPort(ctx, *po)
		if err != nil {
			return err
		}
		portCount++
	}

}

func (p PortDomainService) GetByID(ctx context.Context, msg *port.IDMsg) (*port.Port, error) {
	return p.PortRepository.GetByID(ctx, msg.GetID())
}
