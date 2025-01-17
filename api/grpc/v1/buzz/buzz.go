package buzz

import (
	context "context"
	"log"

	grpc "google.golang.org/grpc"
)

// server is used to implement buzz.SearchServiceServer.
type server struct {
	UnimplementedSearchServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Search(_ context.Context, in *SearchRequest) (*SearchResponse, error) {
	log.Printf("Received: %v", in.GetQuery())
	return &SearchResponse{Results: []string{in.GetQuery()}}, nil
}

func RegisterServices(s *grpc.Server) {
	RegisterSearchServiceServer(s, &server{})
}
