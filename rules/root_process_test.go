package rules

import (
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/nicholasjackson/cnitch/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setup() (*mocks.DockerAPI, *RootProcess) {
	dockerAPI := &mocks.DockerAPI{}

	return dockerAPI, &RootProcess{cli: dockerAPI}
}

func TestExecuteReturnsNothingWhenNoRunningProcesses(t *testing.T) {
	dockerAPI, rule := setup()
	dockerAPI.On("ContainerTop", mock.Anything, mock.Anything, mock.Anything).Return(types.ContainerProcessList{}, nil)

	exceptions, err := rule.Execute("something")

	assert.Nil(t, err)
	assert.Nil(t, exceptions)

}

func TestExecuteReturnsExceptionWhenRunningProcessesIsRoot(t *testing.T) {
	processes := types.ContainerProcessList{}
	processes.Processes = make([][]string, 1)
	processes.Processes[0] = []string{"root", "123", "234", "", "", "", "", "/rootprocess"}

	dockerAPI, rule := setup()
	dockerAPI.On("ContainerTop", mock.Anything, mock.Anything, mock.Anything).Return(processes, nil)

	exceptions, err := rule.Execute("something")

	assert.Nil(t, err)
	assert.Equal(t, 1, len(exceptions))
}

func TestExecuteReturnsExceptionWhenRunningProcessesIsRootId(t *testing.T) {
	processes := types.ContainerProcessList{}
	processes.Processes = make([][]string, 1)
	processes.Processes[0] = []string{"0", "123", "234", "", "", "", "", "/rootprocess"}

	dockerAPI, rule := setup()
	dockerAPI.On("ContainerTop", mock.Anything, mock.Anything, mock.Anything).Return(processes, nil)

	exceptions, err := rule.Execute("something")

	assert.Nil(t, err)
	assert.Equal(t, 1, len(exceptions))
}

func TestExecuteReturns1ExceptionWhen2RunningProcessesAnd1IsRoot(t *testing.T) {
	processes := types.ContainerProcessList{}
	processes.Processes = make([][]string, 2)
	processes.Processes[0] = []string{"root", "123", "234", "", "", "", "", "/rootprocess"}
	processes.Processes[1] = []string{"2323", "123", "234", "", "", "", "", "/userprocess"}

	dockerAPI, rule := setup()
	dockerAPI.On("ContainerTop", mock.Anything, mock.Anything, mock.Anything).Return(processes, nil)

	exceptions, err := rule.Execute("something")

	assert.Nil(t, err)
	assert.Equal(t, 1, len(exceptions))
}

func TestExecuteReturns2ExceptionWhen2RunningProcessesAndBothIsRoot(t *testing.T) {
	processes := types.ContainerProcessList{}
	processes.Processes = make([][]string, 2)
	processes.Processes[0] = []string{"root", "123", "234", "", "", "", "", "/rootprocess"}
	processes.Processes[1] = []string{"root", "233", "2344", "", "", "", "", "/rootprocess"}

	dockerAPI, rule := setup()
	dockerAPI.On("ContainerTop", mock.Anything, mock.Anything, mock.Anything).Return(processes, nil)

	exceptions, err := rule.Execute("something")

	assert.Nil(t, err)
	assert.Equal(t, 2, len(exceptions))
}
