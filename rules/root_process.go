package rules

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
	"github.com/nicholasjackson/cnitch/entities"
)

// RootProcess identifies any root processes running in a container
type RootProcess struct {
	cli client.ContainerAPIClient
}

// NewRootProcess creates a new RootProcess
func NewRootProcess(cli client.ContainerAPIClient) *RootProcess {
	return &RootProcess{cli: cli}
}

// Execute the rule and return any exceptions
func (r *RootProcess) Execute(containerID string) ([]entities.Exception, error) {
	var exceptions []entities.Exception

	procs, err := r.cli.ContainerTop(context.Background(), containerID, nil)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(procs.Processes); i++ {
		// lookup columns from procs.Titles
		if procs.Processes[i][0] == "root" || procs.Processes[i][0] == "0" {
			message := fmt.Sprintf("WARNING: found process running as root: %s pid: %s", procs.Processes[i][7], procs.Processes[i][1])
			exceptions = append(exceptions, entities.Exception{Message: message})
		}
	}

	return exceptions, nil
}
