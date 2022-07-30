package factory

import (
	"bitdump/macro"
	"bitdump/worker"
)

var SproutedJobs []worker.Job

func CreateJobs(amount int) []worker.Job {

	for i := 0; i < amount; i++ {
		newJob := worker.Job{ID: i, Focus: "$$$", ExecFunc: macro.TicketPool}

		SproutedJobs = append(SproutedJobs, newJob)
	}

	return SproutedJobs
}
