package serve

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"github.com/tomcanham/iam/api/grpc/v1/buzz"
	"github.com/tomcanham/iam/cli"
	"google.golang.org/grpc"
)

func registerServices(s *grpc.Server) {
	buzz.RegisterServices(s)
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

			s, err := cli.Server()
			if err != nil {
				log.Fatalf("failed to create server: %v", err)
			}

			registerServices(s)

			log.Printf("server listening at %v", lis.Addr())
			if err := s.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}

			return nil
		},
	}

	return serveCmd
}
