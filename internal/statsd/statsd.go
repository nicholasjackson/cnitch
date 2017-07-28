package statsd

// Client is an interface implementing core features of DataDogs statsD client
type Client interface {
	Incr(name string, tags []string, rate float64) error
	Gauge(name string, value float64, tags []string, rate float64) error
}
