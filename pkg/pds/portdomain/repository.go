package portdomain

import (
	"context"
	"micro_services/api/v1/port"
)

type Repository interface {
	UpsertPort(ctx context.Context, port port.Port) error
	GetByID(ctx context.Context, id string) (*port.Port, error)
}
