package reporting

import (
	"io"
	"log"

	"github.com/nicholasjackson/cnitch/entities"
)

// Logger implements a simple StdOut backend
type Logger struct {
	logger *log.Logger
}

// Report sends information to the backend
func (l *Logger) Report(host entities.Host, infos []entities.Info) error {
	for _, info := range infos {
		l.logger.Printf("Checking image: %s, id: %s\n", info.ContainerImage, info.ContainerID)

		for _, exception := range info.Exceptions {
			l.logger.Println(exception.Message)
		}
		l.logger.Println("")
	}

	return nil
}

// NewLogger creates a logger which will write to the given writer
func NewLogger(writer io.Writer) *Logger {
	return &Logger{logger: log.New(writer, "", 0)}
}
