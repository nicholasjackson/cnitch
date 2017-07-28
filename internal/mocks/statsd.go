package mocks

import "github.com/stretchr/testify/mock"

type MockStatsD struct {
	mock.Mock
}

func (s *MockStatsD) Incr(name string, tags []string, rate float64) error {
	args := s.Called(name, tags, rate)
	return args.Error(0)
}

func (s *MockStatsD) Gauge(name string, value float64, tags []string, rate float64) error {
	args := s.Called(name, value, tags, rate)
	return args.Error(0)
}
