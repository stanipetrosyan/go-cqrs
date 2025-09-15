package main

import (
	"errors"

	goeventbus "github.com/stanipetrosyan/go-eventbus"
)

type SagaStep struct {
	Transaction string
	Compensate  string
}

type Workflow interface {
	EntryPoint(entry SagaStep) Workflow
	Step(step SagaStep) Workflow
	Commit() error
}

type WireTransferWorkflow struct {
	eventbus   goeventbus.EventBus
	eventStore EventStore
	steps      []SagaStep
	entry      SagaStep
}

func NewWorkflow(eventbus goeventbus.EventBus, eventStore EventStore) Workflow {
	return &WireTransferWorkflow{eventbus: eventbus, eventStore: eventStore, steps: []SagaStep{}, entry: SagaStep{}}
}

func (w *WireTransferWorkflow) EntryPoint(entry SagaStep) Workflow {
	w.entry = entry
	w.steps = append(w.steps, entry)

	return w
}

func (w *WireTransferWorkflow) Step(step SagaStep) Workflow {
	w.steps = append(w.steps, step)
	return w
}

func (w *WireTransferWorkflow) Commit() error {
	if (w.entry == SagaStep{}) {
		return errors.New("Workflow must have an entry point")
	}

	if len(w.steps) == 0 {
		return errors.New("Workflow must be at least one step")
	}

	var completedSteps map[string][]Event = make(map[string][]Event)

	for _, event := range w.steps {
		w.eventbus.Channel(event.Transaction).Subscriber().Listen(func(context goeventbus.Context) {
			eventReceived := context.Result().Extract().(TransactionEvent)
			transactionId := eventReceived.Transaction()

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

	return nil
}
