package main

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/mchmarny/gcp-metrics/pkg/metric"
	gha "github.com/sethvargo/go-githubactions"
)

var (
	version   = "v0.0.1-default"
	int64Base = 10
)

func main() {
	// init action with the version and build time
	a := gha.WithFieldsMap(map[string]string{
		"version": version,
		"build":   time.Now().UTC().Format(time.RFC3339),
	})

	// log start and end
	a.Infof("starting action")
	defer a.Infof("action completed")

	// get inputs
	projectID := a.GetInput("project")
	if projectID == "" {
		a.Fatalf("project is required")
	}

	metricName := a.GetInput("metric")
	if metricName == "" {
		a.Fatalf("metric is required")
	}

	valueStr := a.GetInput("value")
	if valueStr == "" {
		a.Fatalf("count is required")
	}
	metricVal, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		a.Fatalf("error parsing count: %s", err)
	}

	// get workflow context
	ctx, err := a.Context()
	if err != nil {
		a.Fatalf("error getting context: %s", err)
	}

	// create labels from the interesting bits in context
	labels := map[string]string{
		"action": ctx.Action,
		"actor":  ctx.Actor,
		"run":    strconv.FormatInt(ctx.RunID, int64Base),
		"repo":   ctx.Repository,
		"ref":    ctx.Ref,
		"sha":    ctx.SHA,
	}

	a.Infof("counting metric %s with value %d...", metricName, metricVal)

	// create counter
	counter, err := getCounter(projectID)
	if err != nil {
		a.Fatalf("error creating metric: %s", err)
	}

	metricType := metric.MakeMetricType(metricName)

	// count
	if err := counter.Count(context.Background(), metricType, metricVal, labels); err != nil {
		a.Fatalf("error counting metric: %s", err)
	}

	// set output
	a.SetOutput("metric", metricType)
	a.SetOutput("value", strconv.FormatInt(metricVal, int64Base))
}

func getCounter(projectID string) (metric.Counter, error) {
	if projectID == "" {
		return nil, errors.New("project is required")
	}
	if projectID == "test" {
		return metric.NewConsoleCounter(), nil
	}
	return metric.New(projectID)
}
