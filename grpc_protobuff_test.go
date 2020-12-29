package benchmark

import (
	"context"
	"testing"
	"time"

	"github.com/xans-me/protobuff-json-benchmark/grpcProtobuff"
	"github.com/xans-me/protobuff-json-benchmark/grpcProtobuff/proto"
	"google.golang.org/grpc"
)

func init() {
	go grpcProtobuff.Start()
	time.Sleep(time.Second)
}

func BenchmarkGRPCProtobuf(b *testing.B) {
	conn, err := grpc.Dial("127.0.0.1:90000", grpc.WithInsecure())
	if err != nil {
		b.Fatalf("grpc connection failed: %v", err)
	}

	client := proto.NewAPIClient(conn)

	for n := 0; n < b.N; n++ {
		doGRPC(client, b)
	}
}

func doGRPC(client proto.APIClient, b *testing.B) {
	resp, err := client.CreateUser(context.Background(), &proto.User{
		Email:    "mulia.ichsan@amarbank.co.id",
		Name:     "Mulia Ichsan",
		Password: "inipassword00",
	})

	if err != nil {
		b.Fatalf("grpc request failed: %v", err)
	}

	if resp == nil || resp.Code != 200 || resp.User == nil || resp.User.Id != "1000000" {
		b.Fatalf("grpc response is wrong: %v", resp)
	}
}
