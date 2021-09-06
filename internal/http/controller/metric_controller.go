package controller

import (
	"encoding/json"
	"net/http"
	"net/url"
	"scheduler/internal/service"
	"scheduler/pkg/logger"
)

type MetricController struct {
	Logger logger.Logger
}

func NewMetricController(logger logger.Logger) *MetricController {
	return &MetricController{
		Logger: logger,
	}
}

func (c *MetricController) List(w http.ResponseWriter, r *http.Request, p *url.Values) {
	c.Logger.Debugf("MetricController:List")
	metricService := service.GetMetricsServiceInstance()

	response, err := json.Marshal(metricService.List())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
