package entities

// Host defines information about the current Docker Engine
type Host struct {
	// Name is the name of the Docker Engine
	Name string
	// IPAddress of the Docker Engine
	IPAddress string
	// DockerVersion is the current Docker Engine version
	DockerVersion string
}
