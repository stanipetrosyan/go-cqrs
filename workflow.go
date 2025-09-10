package main

import (
	goeventbus "github.com/stanipetrosyan/go-eventbus"
)

type SagaStep struct {
	Transaction string
	Compensate  string
}

type Workflow interface {
	Step(step SagaStep) Workflow
	Commit()
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

	for _, event := range w.steps {
		w.eventbus.Channel(event.Transaction).Subscriber().Listen(func(context goeventbus.Context) {
			eventReceived := context.Result().Extract().(TransactionEvent)
			transactionId := context.Result().Extract().(TransactionEvent).Transaction()

			completedSteps[transactionId] = append(completedSteps[transactionId], eventReceived)

			if len(completedSteps[transactionId]) == len(w.steps) {
				commitEvent := WireTransferCompleted{transaction: transactionId}
				w.eventStore.Save(commitEvent.Aggregate(), commitEvent)
			}
		})

		w.eventbus.Channel(event.Compensate).Subscriber().Listen(func(context goeventbus.Context) {
			transactionId := context.Result().Extract().(TransactionEvent).Transaction()

			rollbackEvent := WireTransferRejected{transaction: transactionId}
			w.eventStore.Save(rollbackEvent.Aggregate(), rollbackEvent)
			delete(completedSteps, transactionId)
		})

	}

}
