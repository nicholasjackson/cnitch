package rules

type Info struct {
	ContainerImage string
	ContainerID    string
	Exceptions     []Exception
}

type Exception struct {
	Message string
	Info    interface{}
}

type Rule interface {
	Execute() ([]Info, error)
}
