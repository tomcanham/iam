package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tomcanham/iam/cli"
	"github.com/tomcanham/iam/cmd/client"
	"github.com/tomcanham/iam/cmd/serve"
)

const (
	APP_NAME = "ias_clone"
)

func New(cli *cli.CLI) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   APP_NAME,
		Short: fmt.Sprintf("%s is a cli tool for performing simple grpc operations", APP_NAME),
	}

	rootCmd.PersistentFlags().IntVarP(&cli.Config.Port, "port", "p", cli.Config.Port, "port to use")
	rootCmd.AddCommand(serve.New(cli))
	rootCmd.AddCommand(client.New(cli))

	return rootCmd
}

func Execute(cli *cli.CLI) error {
	viper.AutomaticEnv()

	return New(cli).Execute()
}
