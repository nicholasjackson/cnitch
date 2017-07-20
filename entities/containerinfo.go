package entities

// Info defines the information about discovered problems with the current
// containers running on the engine
type Info struct {
	// ContainerImage is the image name corresponding to the running container
	ContainerImage string

	// ContainerID is the ID of the running container
	ContainerID string

	// Exceptions is a slice of Exception that have been found by running
	// the rules
	Exceptions []Exception
}
