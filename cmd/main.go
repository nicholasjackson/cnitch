package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/nicholasjackson/cnitch"
	"github.com/nicholasjackson/cnitch/reporting"
)

var hostName = flag.String("hostname", "localhost", "hostname or fully qualified domain for the docker host")
var statsDServer = flag.String("statsD-server", "", "hostname or ip address of the statsD collector")

func main() {
	log.Println("Starting Cnitch: Monitoring Docker Processes at:", os.Getenv("DOCKER_HOST"))
	log.Println("")

	c := cnitch.New(10*time.Second, *hostName)

	c.AddReporting(reporting.NewLogger(os.Stdout))

	if *statsDServer != "" {
		statsD, err := reporting.NewStatsD(*statsDServer)
		if err != nil {
			panic(err)
		}

		c.AddReporting(statsD)
	}

	c.Run()
}
