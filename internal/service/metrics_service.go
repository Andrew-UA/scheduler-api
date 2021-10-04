package service

import (
	"sync"
	"time"
)

type Metric struct {
	Method   string
	Route    string
	Duration time.Duration
}

type MetricsService interface {
	Add(metric Metric)
	List() []Metric
}

var instance *metricsService = nil
var once sync.Once

type metricsService struct {
	metrics []Metric
	mu      sync.RWMutex
}

func GetMetricsServiceInstance() MetricsService {
	once.Do(func() {
		instance = new(metricsService)
	})

	return instance
}

func (m *metricsService) Add(metric Metric) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.metrics = append(m.metrics, metric)
}

func (m *metricsService) List() []Metric {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.metrics
}
