package main

import (
	"context"
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

	projectID := a.GetInput("project")
	if projectID == "" {
		a.Fatalf("project is required")
	}

	counter, err := metric.New(projectID)
	if err != nil {
		a.Fatalf("error creating metric: %s", err)
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

	ctx, err := a.Context()
	if err != nil {
		a.Fatalf("error getting context: %s", err)
	}

	labels := map[string]string{
		"action": ctx.Action,
		"actor":  ctx.Actor,
		"run":    strconv.FormatInt(ctx.RunID, int64Base),
		"repo":   ctx.Repository,
		"ref":    ctx.Ref,
		"sha":    ctx.SHA,
	}

	a.Infof("counting metric %s with value %d...", metricName, metricVal)

	if err := counter.Count(context.Background(), metricName, metricVal, labels); err != nil {
		a.Fatalf("error counting metric: %s", err)
	}

	// set output
	a.SetOutput("metric", metricName)
	a.SetOutput("value", strconv.FormatInt(metricVal, int64Base))
}
