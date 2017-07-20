package reporting

import (
	"log"

	"github.com/nicholasjackson/cnitch/entities"
)

// Logger implements a simple StdOut backend
type Logger struct {
}

// Report sends information to the backend
func (l *Logger) Report(host entities.Host, infos []entities.Info) error {
	for _, info := range infos {
		log.Printf("Checking image: %s, id: %s\n", info.ContainerImage, info.ContainerID)

		for _, exception := range info.Exceptions {
			log.Println(exception.Message)
		}
		log.Println("")
	}

	return nil
}
