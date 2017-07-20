package cnitch

import (
	"time"

	"github.com/docker/docker/client"
	"github.com/nicholasjackson/cnitch/reporting"
	"github.com/nicholasjackson/cnitch/rules"
)

type Cnitch struct {
	interval  time.Duration
	rules     []rules.Rule
	reporting []reporting.Backend
}

func (c *Cnitch) AddReporting(r reporting.Backend) {
	c.reporting = append(c.reporting, r)
}

func (c *Cnitch) Run() {
	for {
		for _, rule := range c.rules {
			infos, _ := rule.Execute()
			c.report(infos)
		}

		time.Sleep(c.interval)
	}
}

func (c *Cnitch) report(infos []rules.Info) {
	for _, backend := range c.reporting {
		backend.Report(infos)
	}
}

func New(checkInterval time.Duration, cli client.APIClient) *Cnitch {
	cn := Cnitch{interval: checkInterval}

	cn.reporting = make([]reporting.Backend, 0)

	cn.rules = make([]rules.Rule, 0)
	cn.rules = append(cn.rules, rules.NewRootProcess(cli))

	return &cn
}
