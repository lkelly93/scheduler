package scheduler

import (
	"github.com/lkelly93/scheduler/internal/container"
)

//StartNewScheduler starts a new scheduler with the given options.
//returns the IP address for the given scheduler.
func StartNewScheduler(schedulerName string) (string, error) {
	return container.StartNewScheduler(schedulerName)
}
