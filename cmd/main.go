package main

import (
	"log"
	"os"
	"time"

	"github.com/nicholasjackson/cnitch"
	"github.com/nicholasjackson/cnitch/reporting"
)

func main() {
	log.Println("Starting Cnitch: Monitoring Docker Processes at:", os.Getenv("DOCKER_HOST"))
	log.Println("")

	c := cnitch.New(10 * time.Second)
	c.AddReporting(reporting.NewLogger(os.Stdout))
	c.Run()
}
