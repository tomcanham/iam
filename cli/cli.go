package cli

import (
	"fmt"

	"github.com/tomcanham/iam/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	APP_NAME = "ias_clone"
)

type CLI struct {
	config.Config
}

func New() *CLI {
	config := config.New()

	cli := &CLI{
		Config: config,
	}

	return cli
}

func (c *CLI) GetEndpoint() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *CLI) NewConn() (*grpc.ClientConn, error) {
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
