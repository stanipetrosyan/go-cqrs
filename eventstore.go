package main

import goeventbus "github.com/stanipetrosyan/go-eventbus"

type EventStore interface {
	Save(aggregateName string, event Event)
	Load(aggregateName string) []Event
}

type InMemoryEventStore struct {
	eventbus goeventbus.EventBus
	events   map[string][]Event
}

func (e InMemoryEventStore) Save(aggregateName string, event Event) {
	_, exists := e.events[aggregateName]

	if !exists {
		e.events[aggregateName] = []Event{event}
	} else {
		e.events[aggregateName] = append(e.events[aggregateName], event)
	}

	message := goeventbus.NewMessageBuilder().SetPayload(event).Build()
	e.eventbus.Channel(event.EventName()).Publisher().Publish(message)
	println("saving event:", event.EventName())

}
func (e InMemoryEventStore) Load(aggregateName string) []Event {
	return e.events[aggregateName]
}

func NewEventStore(eventbus goeventbus.EventBus) EventStore {
	return InMemoryEventStore{eventbus: eventbus, events: map[string][]Event{}}
}
