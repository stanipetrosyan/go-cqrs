package main

import (
	goeventbus "github.com/stanipetrosyan/go-eventbus"
)

type SagaStep struct {
	Transaction string
	Compansate  Event
}

type Workflow interface {
	Step(step SagaStep) Workflow
	Listen()
}

type WireTransferWorkflow struct {
	eventbus goeventbus.EventBus
	Steps    []SagaStep
}

func NewWorkflow(eventbus goeventbus.EventBus) Workflow {
	return &WireTransferWorkflow{eventbus: eventbus, Steps: []SagaStep{}}
}

func (w *WireTransferWorkflow) Step(step SagaStep) Workflow {
	w.Steps = append(w.Steps, step)
	return w
}

func (w *WireTransferWorkflow) Listen() {
	var completedSteps map[string][]string = make(map[string][]string)

	for _, event := range w.Steps {
		w.eventbus.Channel(event.Transaction).Subscriber().Listen(func(context goeventbus.Context) {
			aggregate := context.Result().Data.(Event).Aggregate()

			completedSteps[aggregate] = append(completedSteps[aggregate], event.Transaction)

			if len(completedSteps[aggregate]) == len(w.Steps) {
				println("stage completed")
				println(aggregate)
			}

		})

		w.eventbus.Channel(event.Compansate.eventName()).Subscriber().Listen(func(context goeventbus.Context) {
			aggregate := context.Result().Data.(Event).Aggregate()
			completedSteps[aggregate] = []string{}
		})

	}
}
