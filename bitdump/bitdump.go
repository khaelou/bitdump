package bitdump

import (
	"context"
	"fmt"
	"log"
	"time"

	"bitdump/factory"
	"bitdump/worker"
)

const REPEAT_CYCLE = false
const WORKER_COUNT = 1
const JOB_COUNT = 21

func InitClient() {
	fmt.Println("\t** CLIENT INIT **")
	fmt.Println()

	ctx, cancel := context.WithCancel(context.TODO())

	factories := []string{"Mega Millions"}
	constructFactories(ctx, cancel, factories)
}

func constructFactories(ctx context.Context, cancel context.CancelFunc, factories []string) {
	for fID, targetFactory := range factories {
		targetFactory := targetFactory

		log.Println("] FACTORY INITIALIZED:", targetFactory)
		newFactory := Factory{ID: fID, Focus: targetFactory}
		SproutedFactories = append(SproutedFactories, newFactory)

		collector := StartDispatcher(ctx, newFactory, WORKER_COUNT)

		for jID, job := range factory.CreateJobs(JOB_COUNT) {
			collector.JobQueue <- worker.Job{ID: jID, Focus: job.Focus, ExecFunc: job.ExecFunc}
		}
	}

	func() {
		defer cancel()
		fmt.Println("CLEANUP")

		SproutedWorkers = nil
		factory.SproutedJobs = nil
	}()

	defer func() {
		if REPEAT_CYCLE {
			defer InitClient()
			time.Sleep(1 * time.Second)
		} else {
			defer fmt.Println()
			fmt.Println("DONE")
		}
	}()
}
