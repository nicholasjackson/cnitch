package entities

// Exception defines any exceptions which have been found while processing
// the current rule
type Exception struct {
	Message string
	Info    interface{}
}
