package factory

import (
	"bitdump/macro"
	"bitdump/worker"
)

var SproutedJobs []worker.Job

func CreateJobs(amount int) []worker.Job {

	for i := 0; i < amount; i++ {
		newJob := worker.Job{ID: i, Focus: "Winning Numbers", ExecFunc: macro.TicketPool}

		SproutedJobs = append(SproutedJobs, newJob)
	}

	return SproutedJobs
}
