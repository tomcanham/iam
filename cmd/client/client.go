package client

import (
	"github.com/spf13/cobra"
	"github.com/tomcanham/iam/cli"
	"github.com/tomcanham/iam/cmd/client/query"
)

func New(cli *cli.CLI) *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "client",
		Short:   "Client commands",
		Aliases: []string{"c"},
	}

	cmd.PersistentFlags().StringVar(&cli.Host, "hostname", cli.Host, "hostname to use")
	cmd.AddCommand(query.New(cli))

	return cmd
}
