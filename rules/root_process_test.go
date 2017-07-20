package rules

import (
	"testing"

	"github.com/nicholasjackson/cnitch/internal/mocks"
)

func setup() (*mocks.DockerAPI, *RootProcess) {
	dockerAPI := &mocks.DockerAPI{}

	return dockerAPI, &RootProcess{cli: dockerAPI}
}

func TestExecuteReturnsExceptions(t *testing.T) {
	dockerAPI, rule := setup()

	exceptions, err := rule.Execute("something")
}
