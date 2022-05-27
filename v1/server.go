package greeter

import (
	context "context"

	"github.com/rs/zerolog/log"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type greeterServer struct {
	UnimplementedGreeterServer
}

func NewGreeterServer() GreeterServer {
	return &greeterServer{}
}

func (greeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	defer log.Info().Msg("SayHello called")

	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}

	return &HelloReply{Message: "Hello " + req.Name}, nil
}
