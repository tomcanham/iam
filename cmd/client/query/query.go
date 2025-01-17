package query

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/tomcanham/iam/api/grpc/v1/buzz"

	"github.com/tomcanham/iam/cli"
)

func New(cli *cli.CLI) *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "query [flags] query_string",
		Short:   "Perform a query",
		Aliases: []string{"q"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			addr := fmt.Sprintf("%s:%d", cli.Host, cli.Port)
			log.Printf("performing query (against %s): %s", addr, args[0])

			conn, err := cli.NewConn()
			if err == nil {
				defer conn.Close()
				c := buzz.NewSearchServiceClient(conn)

				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				r, err := c.Search(ctx, &buzz.SearchRequest{Query: args[0]})
				if err == nil {
					log.Printf("Results: %v", r.GetResults())
				}
			}

			return
		},
	}

	return cmd
}
