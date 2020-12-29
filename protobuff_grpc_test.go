package benchmarks

import (
	"testing"
	"time"

	pbGrpc "github.com/xans-me/protobuff-json-benchmark/protobuff-grpc"
	"github.com/xans-me/protobuff-json-benchmark/protobuff-grpc/proto"
	"golang.org/x/net/context"
	g "google.golang.org/grpc"
)

func init() {
	go pbGrpc.Start()
	time.Sleep(time.Second)
}

func BenchmarkProtobuffGRPC(b *testing.B) {
	conn, err := g.Dial("127.0.0.1:60000", g.WithInsecure())
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
		Email:    "foo@bar.com",
		Name:     "Bench",
		Password: "bench",
	})

	if err != nil {
		b.Fatalf("grpc request failed: %v", err)
	}

	if resp == nil || resp.Code != 200 || resp.User == nil || resp.User.Id != "1000000" {
		b.Fatalf("grpc response is wrong: %v", resp)
	}
}
