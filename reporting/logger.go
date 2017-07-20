package reporting

import (
	"log"

	"github.com/nicholasjackson/cnitch/rules"
)

type Logger struct {
}

func (l *Logger) Report(infos []rules.Info) error {
	for _, info := range infos {
		log.Printf("Checking image: %s, id: %s\n", info.ContainerImage, info.ContainerID)

		for _, exception := range info.Exceptions {
			log.Println(exception.Message)
		}
		log.Println("")
	}

	return nil
}
