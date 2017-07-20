package rules

import "github.com/nicholasjackson/cnitch/entities"

// Rule defines an interface which should be implemented to check for problems
// with a running container
type Rule interface {
	// Execute the rule
	Execute(containerID string) ([]entities.Exception, error)
}
