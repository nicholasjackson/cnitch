package reporting

import "github.com/nicholasjackson/cnitch/entities"

// Backend defines an interface for a reporting backend
type Backend interface {
	// Report sends exception information to the backend
	Report(entities.Host, []entities.Info) error
}
