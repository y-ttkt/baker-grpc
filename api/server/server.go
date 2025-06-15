// APIサーバーのエントリポイントになるので、パッケージをmainにする
package main

import (
	"fmt"
	"github.com/y-ttkt/baker/gen/api"
	"github.com/y-ttkt/baker/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	api.RegisterPancakeBakerServiceServer(
		server,
		handler.NewBakerHandler(),
	)
	reflection.Register(server)
	go func() {
		log.Printf("start gRPC server port: %v", port)
		server.Serve(lis)
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	server.GracefulStop()
}
