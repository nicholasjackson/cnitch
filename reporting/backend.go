package reporting

import "github.com/nicholasjackson/cnitch/rules"

type Backend interface {
	Report([]rules.Info) error
}
