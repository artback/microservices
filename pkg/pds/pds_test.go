package pds

import (
	"container/list"
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
	"io"
	"micro_services/api/v1/port"
	"micro_services/pkg/pds/repository"
	"reflect"
	"testing"
)

type MockRepository struct {
	*port.Port
	err error
}

func (m MockRepository) UpsertPort(ctx context.Context, port port.Port) error {
	return m.err
}

func (m MockRepository) GetByID(ctx context.Context, id string) (*port.Port, error) {
	return m.Port, m.err
}

type mockServer struct {
	*list.List
	err error
}

func newMockServer(ports []*port.Port, err error) mockServer {
	queue := list.New()
	for _, r := range ports {
		queue.PushBack(r)
	}
	return mockServer{
		queue,
		err,
	}
}
func (m mockServer) SendAndClose(summary *port.PortSummary) error {
	return m.err
}

func (m mockServer) Recv() (*port.Port, error) {
	if m.err != nil {
		return nil, m.err
	}
	e := m.Front()
	if e == nil {
		return nil, io.EOF
	}
	m.Remove(e)
	return e.Value.(*port.Port), nil
}

func (m mockServer) SetHeader(md metadata.MD) error {
	panic("implement me")
}

func (m mockServer) SendHeader(md metadata.MD) error {
	panic("implement me")
}

func (m mockServer) SetTrailer(md metadata.MD) {
	panic("implement me")
}

func (m mockServer) Context() context.Context {
	panic("implement me")
}

func (m mockServer) SendMsg(_ interface{}) error {
	panic("implement me")
}

func (m mockServer) RecvMsg(_ interface{}) error {
	panic("implement me")
}

func TestPortDomainService_RecordPort(t *testing.T) {
	type fields struct {
		PortRepository repository.PortRepository
	}
	type args struct {
		server port.PDService_RecordPortServer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Successful stream",
			fields:  fields{MockRepository{err: nil}},
			args:    args{server: newMockServer([]*port.Port{{ID: "1"}}, nil)},
			wantErr: false,
		},
		{
			name:    "UnSuccessful stream repository error",
			fields:  fields{MockRepository{err: errors.New("something happened")}},
			args:    args{server: newMockServer([]*port.Port{{ID: "1"}}, nil)},
			wantErr: true,
		},
		{
			name:    "UnSuccessful stream recv error",
			fields:  fields{MockRepository{}},
			args:    args{server: newMockServer([]*port.Port{{ID: "1"}}, errors.New("something"))},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PortDomainService{
				PortRepository: tt.fields.PortRepository,
			}
			if err := p.RecordPort(tt.args.server); (err != nil) != tt.wantErr {
				t.Errorf("RecordPort() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPortDomainService_GetByID(t *testing.T) {
	type fields struct {
		PortRepository repository.PortRepository
	}
	type args struct {
		ctx context.Context
		msg *port.IDMsg
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *port.Port
		wantErr bool
	}{
		{
			name:   "getByID return data",
			fields: fields{MockRepository{Port: &port.Port{ID: "4"}}},
			args:   args{ctx: context.Background(), msg: &port.IDMsg{ID: "4"}},
			want:   &port.Port{ID: "4"},
		},
		{
			name:    "getByID return error",
			fields:  fields{MockRepository{err: errors.New("something happened")}},
			args:    args{ctx: context.Background(), msg: &port.IDMsg{ID: "4"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PortDomainService{
				PortRepository: tt.fields.PortRepository,
			}
			got, err := p.GetByID(tt.args.ctx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
