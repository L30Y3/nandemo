package events

import events "github.com/L30Y3/nandemo/shared/proto/protoevents"

type EventBus interface {
	PublishOrderCreated(event *events.OrderCreatedEvent) error
}
