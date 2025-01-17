package serve

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"github.com/tomcanham/iam/api/grpc/v1/buzz"
	"github.com/tomcanham/iam/cli"
	"google.golang.org/grpc"
)

// server is used to implement buzz.SearchServiceServer.
type server struct {
	buzz.UnimplementedSearchServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Search(_ context.Context, in *buzz.SearchRequest) (*buzz.SearchResponse, error) {
	log.Printf("Received: %v", in.GetQuery())
	return &buzz.SearchResponse{Results: []string{in.GetQuery()}}, nil
}

func New(cli *cli.CLI) *cobra.Command {
	var serveCmd = &cobra.Command{
		Use:     "serve",
		Short:   "Start the server",
		Aliases: []string{"s"},
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Printf("Starting server on port %d", cli.Port)

			lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cli.Port))
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}

			s := grpc.NewServer()
			buzz.RegisterSearchServiceServer(s, &server{})
			log.Printf("server listening at %v", lis.Addr())
			if err := s.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
			return nil
		},
	}

	return serveCmd
}
