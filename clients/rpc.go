package clients

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"

	"github.com/cro4k/authorize/clients/authrpc"
)

type rpcClient struct {
	host string
	mu   *sync.RWMutex
	conn *grpc.ClientConn

	options []grpc.DialOption
}

func (c *rpcClient) checkConn() (*grpc.ClientConn, error) {

	if conn, err := c.getConn(); err == nil {
		return conn, nil
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	conn, err := c.dial(context.Background())
	if err != nil {
		return nil, err
	}
	c.conn.Close()
	c.conn = conn
	return conn, nil
}

func (c *rpcClient) getConn() (*grpc.ClientConn, error) {
	c.mu.RLocker()
	defer c.mu.RUnlock()
	if c.conn == nil {
		return nil, errors.New("not connected")
	}
	if state := c.conn.GetState(); state != connectivity.TransientFailure && state != connectivity.Shutdown {
		return c.conn, nil
	} else {
		return nil, fmt.Errorf("unexpected connection state: %v", state)
	}
}

func (c *rpcClient) dial(ctx context.Context) (*grpc.ClientConn, error) {
	return grpc.DialContext(ctx, c.host, c.options...)
}

func (c *rpcClient) AuthClient() (authrpc.AuthServiceClient, error) {
	cc, err := c.checkConn()
	if err != nil {
		return nil, err
	}
	return authrpc.NewAuthServiceClient(cc), nil
}

type RPCClient interface {
	AuthClient() (authrpc.AuthServiceClient, error)
}

func NewRPCClient(host string, options ...grpc.DialOption) RPCClient {
	c := &rpcClient{
		host:    host,
		mu:      new(sync.RWMutex),
		options: options,
	}
	return c
}
