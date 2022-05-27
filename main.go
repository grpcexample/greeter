package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpcexample/greeter/v1"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

var (
	grpcPort = os.Getenv("GRPC_PORT")
)

func init() {
	if grpcPort == "" {
		grpcPort = "9090"
	}
}

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	grpcAddr := fmt.Sprintf(":%s", grpcPort)

	log := log.With().Str("port", grpcPort).Logger()

	log.Info().Msg("starting gRPC server")

	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()

	greeterServer := greeter.NewGreeterServer()

	greeter.RegisterGreeterServer(grpcServer, greeterServer)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			panic(err)
		}
	}()

	log.Info().Msg("gRPC server started")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt

	log.Info().Msg("stopping gRPC server")
	grpcServer.GracefulStop()
	log.Info().Msg("gRPC server stopped")
}
