package config

import (
	"fmt"
	"os"
)

func LoadGrpcAddress() string {
	return fmt.Sprintf("%s:%s", os.Getenv("GRPC_HOST"), os.Getenv("GRPC_PORT"))
}
