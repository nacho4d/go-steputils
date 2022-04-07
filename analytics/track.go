package analytics

import (
	"fmt"

	"github.com/bitrise-io/go-utils/v2/analytics"
	"github.com/bitrise-io/go-utils/v2/env"
)

// TrackerFactory ...
type TrackerFactory func(...analytics.Properties) analytics.Tracker

const (
	stepExecutionIDEnvKey = "BITRISE_STEP_EXECUTION_ID"
	stepExecutionID       = "step_execution_id"
)

// NewStepTracker ...
func NewStepTracker(repository env.Repository, trackerFactory TrackerFactory) (analytics.Tracker, error) {
	id := repository.Get(stepExecutionIDEnvKey)
	if id == "" {
		return nil, fmt.Errorf("no step execution ID found")
	}
	return trackerFactory(analytics.Properties{stepExecutionID: id}), nil
}

// NewDefaultStepTracker ...
func NewDefaultStepTracker(repository env.Repository) (analytics.Tracker, error) {
	return NewStepTracker(repository, analytics.NewDefaultTracker)
}
