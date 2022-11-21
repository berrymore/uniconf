package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"strings"
	"uniconf/internal/container"
)

type Docker interface {
	GetContainers() (map[string]*container.Container, error)
}

type BasicDocker struct {
	client client.APIClient
	ctx    context.Context
}

func (d *BasicDocker) GetContainers() (map[string]*container.Container, error) {
	raw, err := d.client.ContainerList(d.ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}

	containers := map[string]*container.Container{}

	for _, c := range raw {
		if l, ok := c.Labels["com.dockerconf.component"]; ok {
			containers[l] = &container.Container{
				Name:    strings.Trim(c.Names[0], "/"),
				Address: c.NetworkSettings.Networks["bridge"].IPAddress,
			}
		}
	}

	return containers, nil
}

func CreateDocker(client client.APIClient) Docker {
	return &BasicDocker{
		client: client,
		ctx:    context.Background(),
	}
}
