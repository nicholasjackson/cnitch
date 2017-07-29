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
var statsDServer = flag.String("statsd-server", "", "hostname or ip address of the statsD collector")
var checkInterval = flag.String("check", "10s", "check interval, 1m, 10s, 1h, etc.")

func main() {
	flag.Parse()

	duration, err := time.ParseDuration(*checkInterval)
	if err != nil {
		log.Fatalln("invalid duration, specify a duration such as 10s = 10 seconds")
	}

	log.Println("Starting Cnitch: Monitoring Docker Processes at:", os.Getenv("DOCKER_HOST"))
	log.Println("Checking for root processes every:", duration.String())

	c := cnitch.New(duration, *hostName)

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
