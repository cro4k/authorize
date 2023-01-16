package rpc_test

import (
	"context"
	"github.com/cro4k/authorize/rpc/authrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func TestRPC(t *testing.T) {
	//go rpc.NewServer().Run()

	conn, err := grpc.DialContext(context.Background(), "127.0.0.1:8089", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.Close()
	client := authrpc.NewAuthServiceClient(conn)
	resp, err := client.VerifyToken(context.Background(), &authrpc.VerifyTokenRequest{
		Cid:   "21f20ab6-01cb-4d21-b5bc-dc7eff92b939",
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiJiZjc3OWZkYS05M2FlLTQ4ZGYtOTBjYy0xMmU2MmRjNjhhZjUiLCJjaWQiOiIyMWYyMGFiNi0wMWNiLTRkMjEtYjViYy1kYzdlZmY5MmI5MzkiLCJ1c2VybmFtZSI6InNpbmdzZW4iLCJjcyI6MH0.wLjOblDLQ6Yu40KwIoHxDdnGVHKOnOoHLTKzaHNAwcA",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)

}
