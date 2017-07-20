package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var cli *client.Client

func main() {
	log.Println("Starting Cnitch: Monitoring Docker Processes at:", os.Getenv("DOCKER_HOST"))
	log.Println("")

	var err error
	cli, err = client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	for {
		report()
		time.Sleep(5000 * time.Millisecond)
	}
}

func report() {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		log.Println("Inspecting processes for:", container.Image)

		procs, err := cli.ContainerTop(context.Background(), container.ID, nil)
		if err != nil {
			panic(err)
		}

		for i := 0; i < len(procs.Processes); i++ {
			// lookup columns from procs.Titles
			if procs.Processes[i][0] == "root" || procs.Processes[i][0] == "0" {
				log.Println("WARNING: found process running as root: ", procs.Processes[i][7], "pid:", procs.Processes[i][1])
			}
		}

		log.Println("")
	}
}
