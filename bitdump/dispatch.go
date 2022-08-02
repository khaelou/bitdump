package bitdump

import (
	"context"
	"fmt"
	"log"
	"math"

	"bitdump/worker"
)

var SproutedFactories []Factory
var SproutedWorkers []worker.Worker

var FactoryChannel = make(chan chan Factory, math.MaxInt8)
var WorkerChannel = make(chan chan worker.Job, math.MaxInt8)

type Factory struct {
	ID    int
	Focus string
}

type Collector struct {
	JobQueue chan worker.Job
	EndCycle chan bool
}

func StartDispatcher(ctx context.Context, targetFactory Factory, workerCount int) Collector {
	var i int

	input := make(chan worker.Job)
	end := make(chan bool)
	collector := Collector{JobQueue: input, EndCycle: end}

	roles := []string{"Khaelou"}

	for i < workerCount {
		for _, role := range roles {
			i++

			log.Printf("~ Starting Worker #%d (%s @ %s)\n", i, role, targetFactory.Focus)

			worker := worker.Worker{
				ID:            i,
				Factory:       targetFactory.Focus,
				Role:          role,
				WorkerChannel: WorkerChannel,
				JobChannel:    make(chan worker.Job),
				EndShift:      make(chan bool),
			}

			worker.StartWorker(ctx)
			SproutedWorkers = append(SproutedWorkers, worker)
		}
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				if err := ctx.Err(); err != nil {
					_ = fmt.Sprintf("\nCTX DONE: %v", err)
				}

				// Close channels
				close(collector.JobQueue)
				close(collector.EndCycle)

				//fmt.Printf("COMPLETE.\n\n")
				return
			case <-end:
				for _, w := range SproutedWorkers {
					w.StopWorker(ctx)
				}
				return
			case job := <-input:
				worker := <-WorkerChannel
				worker <- job
			}
		}
	}()

	return collector
}
