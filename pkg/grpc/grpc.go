package grpc

import (
	"github.com/yogarn/arten/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClientConn() (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(config.LoadGrpcAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
