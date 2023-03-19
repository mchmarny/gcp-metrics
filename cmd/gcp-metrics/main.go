package main

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/mchmarny/gcp-metrics/pkg/metric"
	"github.com/pkg/errors"
	gha "github.com/sethvargo/go-githubactions"
)

var (
	version   = "v0.0.1-default"
	int64Base = 10
)

func main() {
	a := gha.WithFieldsMap(map[string]string{
		"version": version,
		"build":   time.Now().UTC().Format(time.RFC3339),
	})

	a.Infof("starting action")
	defer a.Infof("action completed")

	ctx, err := a.Context()
	if err != nil {
		a.Fatalf("error getting context: %s", err)
	}

	labels := map[string]string{
		"action":   ctx.Action,
		"actor":    ctx.Actor,
		"event":    ctx.EventName,
		"job":      ctx.Job,
		"ref":      parseTag(ctx.Ref),
		"repo":     ctx.Repository,
		"run":      strconv.FormatInt(ctx.RunID, int64Base),
		"sha":      ctx.SHA,
		"workflow": ctx.Workflow,
	}

	m, v, err := count(a.GetInput("project"), a.GetInput("metric"), a.GetInput("value"), labels)
	if err != nil {
		a.Fatalf("error counting metric: %s", err)
	}

	a.SetOutput("metric", m)
	a.SetOutput("value", strconv.FormatInt(v, int64Base))
}

func count(project, name, val string, labels map[string]string) (string, int64, error) {
	if project == "" {
		return "", 0, errors.New("project is required")
	}
	if name == "" {
		return "", 0, errors.New("metric is required")
	}
	if val == "" {
		return "", 0, errors.New("value is required")
	}

	metricVal, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return "", 0, errors.Wrapf(err, "error parsing value: %s", val)
	}

	counter, err := getCounter(project)
	if err != nil {
		return "", 0, err
	}

	metricType := metric.MakeMetricType(name)

	if err := counter.Count(context.Background(), metricType, metricVal, labels); err != nil {
		return "", 0, errors.Wrapf(err, "error counting metric: %s with value: %d", metricType, metricVal)
	}

	return metricType, metricVal, nil
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

func parseTag(v string) string {
	return strings.ReplaceAll(v, "refs/tags/", "")
}
