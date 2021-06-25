package client

import (
	"context"
	"micro_services/api/v1/api"
	"micro_services/api/v1/port"
)

type ApiClient struct {
	port.PDServiceClient
}

func (c ApiClient) GetByID(ctx context.Context, msg *api.IDMsg) (*port.Port, error) {
	return c.PDServiceClient.GetByID(ctx, &port.IDMsg{ID: msg.ID})
}
