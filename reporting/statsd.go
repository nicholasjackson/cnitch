package reporting

import (
	dogstatsd "github.com/DataDog/datadog-go/statsd"
	"github.com/nicholasjackson/cnitch/entities"
	"github.com/nicholasjackson/cnitch/internal/statsd"
)

// StatsD implements a StatsD backend
type StatsD struct {
	client statsd.Client
}

// Report sends information to the backend
func (l *StatsD) Report(host entities.Host, infos []entities.Info) error {
	rootTags := []string{"host:" + host.HostName}

	for _, info := range infos {
		tags := append(rootTags, "container:"+info.ContainerImage)

		for _, exception := range info.Exceptions {
			l.client.Incr("cnitch."+exception.Tag, tags, 1)
		}
	}

	return nil
}

// NewStatsD creates a new instance of the statsd backend
func NewStatsD(endpoint string) (*StatsD, error) {
	c, err := dogstatsd.New(endpoint)
	if err != nil {
		return nil, err
	}

	return &StatsD{c}, nil
}
