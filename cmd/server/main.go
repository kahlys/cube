package main

import (
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	"log/slog"
	"math/rand"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	pb "github.com/kahlys/cube/internal/proto"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	response := &pb.HelloResponse{
		Message: fmt.Sprintf("Hello %s", req.Name),
		Id:      base64.RawStdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", rand.Intn(1000000)))),
	}
	return response, nil
}

func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	slog.Info("RequestReceived", "method", info.FullMethod, "request", req)

	resp, err := handler(ctx, req)

	if err != nil {
		slog.Error("RequestFailed", "method", info.FullMethod, "error", status.Convert(err).Message())
	} else {
		slog.Info("RequestCompleted", "method", info.FullMethod, "response", resp)
	}

	return resp, err
}

func main() {
	port := flag.String("port", "50051", "The server port")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		slog.Error("InitListenerFailed", "error", err)
		os.Exit(1)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(loggingInterceptor),
	)
	pb.RegisterHelloServiceServer(s, &server{})
	reflection.Register(s)

	slog.Info("ServerStart", "port", *port)
	if err := s.Serve(lis); err != nil {
		slog.Error("ServeFailed", "error", err)
		os.Exit(1)
	}
}
