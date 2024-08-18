package collectors

import (
	"strings"
	"testing"
	"github.com/coolapso/prometheus-exporter-boilerplate/internal/slogLogger"
)

func TestNewMetrics(t *testing.T) {
	metrics := newMetrics()

	t.Run("Test Sample metric A", func(t *testing.T) {
		expected := `Desc{fqName: "boilerplate_sampleMetricA", help: "Sample metric A description", constLabels: {}, variableLabels: {Label1}}`
		got := metrics.sampleMetricA.String()
		if !strings.Contains(got, expected) {
			t.Fatalf("Metric does not contain expected fqName, expected: %v, got %v", expected, got) 
		}
	})

	t.Run("Test Sample metric B", func(t *testing.T) {
		expected := `Desc{fqName: "boilerplate_sampleMetricB", help: "Sample metric B description", constLabels: {}, variableLabels: {Label1}}`
		got := metrics.sampleMetricB.String()
		if !strings.Contains(got, expected) {
			t.Fatalf("Metric does not contain expected fqName, expected: %v, got %v", expected, got) 
		}

	})
}

func TestNewExporter(t *testing.T) {
	logger, _ := slogLogger.NewLogger("info", "text")
	settings := &Settings{
		LogLevel:    "debug",
		LogFormat:   "text",
		ListenPort:  "8080",
		MetricsPath: "/metrics",
		Address:     "localhost",
	}

	exporter, err := NewExporter(settings, logger)
	if err != nil { 
		t.Fatal("Failed to create new exporter:", err)
	}

	t.Run("Test sampleMetricA", func(t *testing.T) {
		got := exporter.sampleMetricA()
		if got <= 0.0 || got > 1.0  { 
			t.Fatalf("Got wrong value, expected value between 0 and 1, got: %v", got)
		}

	})

	t.Run("Test sampleMetricB", func(t *testing.T) {
		got := exporter.sampleMetricB()
		if got <= 0.0 || got > 1.0  { 
			t.Fatalf("Got wrong value, expected value between 0 and 1, got: %v", got)
		}
	})
}
