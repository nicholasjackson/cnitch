package rules

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// RootProcess identifies any root processes running in a container
type RootProcess struct {
	cli client.ContainerAPIClient
}

func NewRootProcess(cli client.ContainerAPIClient) *RootProcess {
	return &RootProcess{cli: cli}
}

func (r *RootProcess) Execute() ([]Info, error) {
	containers, err := r.cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	var infos []Info

	for _, container := range containers {
		info := Info{
			ContainerImage: container.Image,
			ContainerID:    container.ID,
			Exceptions:     make([]Exception, 0),
		}

		procs, err := r.cli.ContainerTop(context.Background(), container.ID, nil)
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(procs.Processes); i++ {
			// lookup columns from procs.Titles
			if procs.Processes[i][0] == "root" || procs.Processes[i][0] == "0" {
				message := fmt.Sprintf("WARNING: found process running as root: %s pid: %s", procs.Processes[i][7], procs.Processes[i][1])
				info.Exceptions = append(info.Exceptions, Exception{Message: message})
			}
		}

		infos = append(infos, info)
	}

	return infos, nil
}
