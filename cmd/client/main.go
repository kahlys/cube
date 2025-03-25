package main

import (
	"context"
	"flag"
	"log/slog"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/kahlys/cube/internal/proto"
)

func main() {
	name := flag.String("name", "Alice", "The name to greet")
	flag.Parse()

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("InitClientFailed", "error", err)
		os.Exit(1)
	}
	defer conn.Close()

	clt := pb.NewHelloServiceClient(conn)
	res, err := clt.Hello(context.Background(), &pb.HelloRequest{Name: *name})
	if err != nil {
		slog.Error("HelloFailed", "error", err)
		os.Exit(1)
	}
	slog.Info("HelloResponse", "message", res.Message, "id", res.Id)
}
