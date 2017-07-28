package entities

// Exception defines any exceptions which have been found while processing
// the current rule
type Exception struct {
	// Message represents a human readable message relating to the exception
	Message string

	// Code represents an error code associated with the error
	Code int

	// Tag is the tag associated with the error which would be sent to the backend
	// such as statsD e.g. exception.root_process
	Tag string

	// Info is the raw information associated with the exception
	Info interface{}
}
