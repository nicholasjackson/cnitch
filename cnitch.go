package cnitch

import (
	"context"
	"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/nicholasjackson/cnitch/entities"
	"github.com/nicholasjackson/cnitch/reporting"
	"github.com/nicholasjackson/cnitch/rules"
)

// Cnitch is a library which implements rules to check for security problems
// in running docker containers
type Cnitch struct {
	interval  time.Duration
	rules     []rules.Rule
	reporting []reporting.Backend
	host      entities.Host
	cli       *client.Client
}

// AddReporting allows the addition of reporting backends to which any
// discovered vulnerabilities will be sent
func (c *Cnitch) AddReporting(r reporting.Backend) {
	c.reporting = append(c.reporting, r)
}

// Run cnitch and process any rules, this is a blocking call
func (c *Cnitch) Run() {
	for {

		containers, err := c.cli.ContainerList(context.Background(), types.ContainerListOptions{})
		if err != nil {
			log.Println("Error getting list of containers:", err)
		}

		var infos []entities.Info

		for _, container := range containers {
			info := entities.Info{
				ContainerImage: container.Image,
				ContainerID:    container.ID,
				Exceptions:     make([]entities.Exception, 0),
			}

			for _, rule := range c.rules {
				exceptions, _ := rule.Execute(container.ID)
				info.Exceptions = append(info.Exceptions, exceptions...)
			}

			infos = append(infos, info)
		}

		c.report(infos)
		time.Sleep(c.interval)
	}
}

func (c *Cnitch) report(infos []entities.Info) {

	for _, backend := range c.reporting {
		backend.Report(c.host, infos)
	}
}

// New creates a new Cnitch with the given checkInterval
func New(checkInterval time.Duration, hostName string) *Cnitch {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	cn := Cnitch{interval: checkInterval, cli: cli}

	dockerInfo, err := cli.Info(context.Background())
	if err != nil {
		panic(err)
	}

	cn.host.HostName = hostName
	cn.host.Name = dockerInfo.Name
	cn.host.DockerVersion = dockerInfo.ServerVersion

	cn.reporting = make([]reporting.Backend, 0)

	cn.rules = make([]rules.Rule, 0)
	cn.rules = append(cn.rules, rules.NewRootProcess(cli))

	return &cn
}
