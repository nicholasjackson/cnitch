package reporting

import (
	"bytes"
	"strings"
	"testing"

	"github.com/nicholasjackson/cnitch/entities"
	"github.com/stretchr/testify/assert"
)

func setupLogger() (*bytes.Buffer, *Logger) {
	stringWriter := bytes.Buffer{}
	logger := NewLogger(&stringWriter)

	return &stringWriter, logger
}

func TestLogsCorrectHeaderDetails(t *testing.T) {
	writer, logger := setupLogger()

	infos := make([]entities.Info, 1)
	infos[0].ContainerImage = "fakeimage"
	infos[0].ContainerID = "abc123"

	logger.Report(entities.Host{}, infos)

	lines := strings.Split(writer.String(), "\n")

	assert.Equal(t, "Checking image: fakeimage, id: abc123", lines[0][20:len(lines[0])])
}

func TestLogsCorrectExceptionDetails(t *testing.T) {
	writer, logger := setupLogger()

	infos := make([]entities.Info, 1)
	infos[0].ContainerImage = "fakeimage"
	infos[0].ContainerID = "abc123"
	infos[0].Exceptions = []entities.Exception{entities.Exception{Message: "ooooh root"}}

	logger.Report(entities.Host{}, infos)

	lines := strings.Split(writer.String(), "\n")

	assert.Equal(t, ">> ooooh root", lines[1][20:len(lines[1])])
}

func TestLogsWhenNoException(t *testing.T) {
	writer, logger := setupLogger()

	infos := make([]entities.Info, 1)
	infos[0].ContainerImage = "fakeimage"
	infos[0].ContainerID = "abc123"

	logger.Report(entities.Host{}, infos)

	lines := strings.Split(writer.String(), "\n")

	assert.Equal(t, ">> No root processes found", lines[1][20:len(lines[1])])
}
