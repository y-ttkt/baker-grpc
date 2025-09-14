// APIサーバーのエントリポイントになるので、パッケージをmainにする
package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/y-ttkt/baker/gen/api"
	"github.com/y-ttkt/baker/handler"
	"go.uber.org/zap"
	"go.uber.org/zap/zapgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
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

	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to init zap logger: %v", err)
	}

	grpclog.SetLoggerV2(zapgrpc.NewLogger(zapLogger))

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(zapLoggerAdapter(zapLogger)),
		),
	)
	api.RegisterPancakeBakerServiceServer(
		server,
		handler.NewBakerHandler(),
	)
	api.RegisterImageUploadServiceServer(
		server,
		handler.NewImageUploadHandler(),
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

func zapLoggerAdapter(l *zap.Logger) logging.Logger {
	// logging.Logger は関数型 LoggerFunc でも実装できます
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		// Interceptor が渡してくる fields と、Context に入っているフィールドを統合
		all := append(logging.ExtractFields(ctx), fields...)

		// []any（key, val, key, val, ...）→ []zap.Field に変換
		zfs := make([]zap.Field, 0, len(all)/2)
		for i := 0; i+1 < len(all); i += 2 {
			k, ok := all[i].(string)
			if !ok {
				continue
			}
			zfs = append(zfs, zap.Any(k, all[i+1]))
		}

		switch lvl {
		case logging.LevelDebug:
			l.Debug(msg, zfs...)
		case logging.LevelInfo:
			l.Info(msg, zfs...)
		case logging.LevelWarn:
			l.Warn(msg, zfs...)
		case logging.LevelError:
			l.Error(msg, zfs...)
		default:
			l.Info(msg, zfs...)
		}
	})
}
