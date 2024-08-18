package collectors

import (
	"log/slog"
	"math/rand"

	"github.com/prometheus/client_golang/prometheus"
)

// Namespace constant value prefixed on metrics boilerplate_
const (
	namespace = "boilerplate"
)

// exporter settings, more settings can be added if needed
type Settings struct {
	LogLevel    string
	LogFormat   string
	MetricsPath string
	ListenPort  string
	Address     string
}

type metrics struct {
	sampleMetricA *prometheus.Desc
	sampleMetricB *prometheus.Desc
}

type Exporter struct {
	metrics  *metrics
	Settings *Settings
	Logger   *slog.Logger
}

// Describes all metrics
// implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.metrics.sampleMetricA
	ch <- e.metrics.sampleMetricB
}

// Collects metrics configured and returns them as prometheus metrics
// implements prometheus.Collector
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {

	ch <- prometheus.MustNewConstMetric(
		e.metrics.sampleMetricA,
		prometheus.GaugeValue,
		e.sampleMetricA(),
		"label1Value",
	)

	ch <- prometheus.MustNewConstMetric(
		e.metrics.sampleMetricB,
		prometheus.GaugeValue,
		e.sampleMetricB(),
		"label1Value",
	)
}


// Functions that gather and return the metric values
func (e *Exporter) sampleMetricA() float64 {
	return rand.Float64()
}

func (e *Exporter) sampleMetricB() float64 {
	return rand.Float64()
}

// Initializes the metrics
func newMetrics() *metrics {
	return &metrics{
		sampleMetricA: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "sampleMetricA"),
			"Sample metric A description",
			[]string{"Label1"}, nil,
		),

		sampleMetricB: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "sampleMetricB"),
			"Sample metric B description",
			[]string{"Label1"}, nil,
		),
	}
}

// Initializes the exporter
func NewExporter(s *Settings, logger *slog.Logger) (*Exporter, error) {
	metrics := newMetrics()

	exporter := &Exporter{
		metrics:  metrics,
		Settings: s,
		Logger:   logger,
	}

	return exporter, nil
}
