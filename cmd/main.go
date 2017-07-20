package main

import (
	"log"
	"os"
	"time"

	"github.com/docker/docker/client"
	"github.com/nicholasjackson/cnitch"
	"github.com/nicholasjackson/cnitch/reporting"
)

func main() {
	log.Println("Starting Cnitch: Monitoring Docker Processes at:", os.Getenv("DOCKER_HOST"))
	log.Println("")

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	c := cnitch.New(10*time.Second, cli)
	c.AddReporting(&reporting.Logger{})
	c.Run()
}
