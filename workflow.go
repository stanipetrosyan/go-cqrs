package main

import (
	goeventbus "github.com/stanipetrosyan/go-eventbus"
)

type SagaStep struct {
	Transaction string
}

type Workflow interface {
	Step(step SagaStep) Workflow
	Commit()
	//Rollback()
}

type WireTransferWorkflow struct {
	eventbus   goeventbus.EventBus
	eventStore EventStore
	steps      []SagaStep
}

func NewWorkflow(eventbus goeventbus.EventBus, eventStore EventStore) Workflow {
	return &WireTransferWorkflow{eventbus: eventbus, eventStore: eventStore, steps: []SagaStep{}}
}

func (w *WireTransferWorkflow) Step(step SagaStep) Workflow {
	w.steps = append(w.steps, step)
	return w
}

func (w *WireTransferWorkflow) Commit() {
	if len(w.steps) == 0 {
		return
	}

	var completedSteps map[string][]Event = make(map[string][]Event)
	var aggregateCompleted = ""

	go func() {
		for {
			if aggregateCompleted != "" {
				println("stage completed")

				for _, event := range completedSteps[aggregateCompleted] {
					println(event.Aggregate())
					println(event.eventName())
					w.eventStore.Save(event.Aggregate(), event)
				}
				aggregateCompleted = ""
			}
		}

	}()

	for _, event := range w.steps {
		w.eventbus.Channel(event.Transaction).Subscriber().Listen(func(context goeventbus.Context) {
			eventReceived := context.Result().Extract().(Event)
			aggregate := context.Result().Extract().(Event).Aggregate()

			completedSteps[aggregate] = append(completedSteps[aggregate], eventReceived)

			if len(completedSteps[aggregate]) == len(w.steps) {
				aggregateCompleted = aggregate
			}

		})

	}

}
