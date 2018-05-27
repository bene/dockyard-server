package server

import "github.com/docker/docker/client"

func NewDockerClient(apiVersion string) (*client.Client, error) {

	cli, err := client.NewClientWithOpts(client.WithVersion(apiVersion))
	if err != nil {
		return nil, err
	}

	return cli, nil
}
