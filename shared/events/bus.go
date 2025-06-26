package events

import pb "github.com/L30Y3/nandemo/shared/proto/protoevents"

type OrderCreatedEventHandler func(event *pb.OrderCreatedEvent)

type EventBus interface {
	PublishOrderCreated(event *pb.OrderCreatedEvent) error
	SubscribeToOrderCreated(handler OrderCreatedEventHandler)
}
