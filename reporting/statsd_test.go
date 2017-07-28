package reporting

import (
	"testing"

	"github.com/nicholasjackson/cnitch/entities"
	"github.com/nicholasjackson/cnitch/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var infos = []entities.Info{
	entities.Info{
		ContainerImage: "rootcontainer:latest",
		Exceptions: []entities.Exception{
			entities.Exception{
				Tag: "exception.root_process",
			},
		},
	},
}

func setupStatsD() (*StatsD, *mocks.MockStatsD) {
	mockStatsD := &mocks.MockStatsD{}
	mockStatsD.On("Incr", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	s := &StatsD{client: mockStatsD}

	return s, mockStatsD
}

func TestIncrementsContainerStats(t *testing.T) {
	r, m := setupStatsD()
	host := entities.Host{}

	r.Report(host, infos)

	m.AssertCalled(t, "Incr", mock.Anything, mock.Anything, mock.Anything)
}

func TestSetsCorrectMetricName(t *testing.T) {
	r, m := setupStatsD()
	host := entities.Host{}

	r.Report(host, infos)

	m.AssertCalled(t, "Incr", "cnitch.exception.root_process", mock.Anything, mock.Anything)
}

func TestSetsCorrectMetricTags(t *testing.T) {
	r, m := setupStatsD()
	host := entities.Host{HostName: "127.0.0.1"}

	r.Report(host, infos)

	tags := []string{"host:127.0.0.1", "container:rootcontainer:latest"}

	arg := m.Calls[0].Arguments.Get(1)

	assert.Equal(t, tags, arg)
	m.AssertCalled(t, "Incr", "cnitch.exception.root_process", tags, mock.Anything)
}
