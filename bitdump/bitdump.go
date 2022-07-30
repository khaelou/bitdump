package bitdump

import (
	"bitdump/factory"
	"bitdump/worker"
	"fmt"
	"log"
)

const REPEAT_CYCLE = false // true if WORKER_COUNT > 1
const WORKER_COUNT = 1
const JOB_COUNT = 21

func InitClient() {
	fmt.Println("\t** CLIENT INIT **")
	fmt.Println()

	factories := []string{"Mega Millions"}
	constructFactories(factories)
}

func constructFactories(factories []string) {
	for fID, targetFactory := range factories {
		targetFactory := targetFactory

		log.Println("] FACTORY INITIALIZED:", targetFactory)
		newFactory := Factory{ID: fID, Focus: targetFactory}
		SproutedFactories = append(SproutedFactories, newFactory)

		collector := StartDispatcher(newFactory, WORKER_COUNT)

		for jID, job := range factory.CreateJobs(JOB_COUNT) {
			collector.JobQueue <- worker.Job{ID: jID, Focus: job.Focus, ExecFunc: job.ExecFunc}
		}

		go func(id int) {
			fmt.Println()
			fmt.Printf("[***] execFunc() DONE @ %s [nF: %d | nW: %d | nWch: %d]\n\n", SproutedFactories[id].Focus, len(SproutedFactories), len(SproutedWorkers), len(WorkerChannel))
		}(fID)
	}

	func() {
		fmt.Println("CLEANUP")

		SproutedWorkers = nil
		factory.SproutedJobs = nil
	}()

	if REPEAT_CYCLE {
		InitClient()
	} else {
		defer fmt.Println("DONE")
	}
}
