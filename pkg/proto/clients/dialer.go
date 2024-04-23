package clients

import (
	"context"

	"google.golang.org/grpc"
)

type GrpcDialer func(ctx context.Context, target string, options ...grpc.DialOption) (*grpc.ClientConn, error)
