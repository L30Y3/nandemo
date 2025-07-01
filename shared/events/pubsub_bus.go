package events

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"cloud.google.com/go/pubsub"
	pb "github.com/L30Y3/nandemo/shared/proto/protoevents"
)

type PubSubBus struct {
	client        *pubsub.Client
	topic         *pubsub.Topic
	subscription  *pubsub.Subscription
	ctx           context.Context
	subCancel     context.CancelFunc
	orderHandlers []OrderCreatedEventHandler
	mu            sync.Mutex
}

func NewPubSubBus(ctx context.Context, projectId, topicId, subId string) (*PubSubBus, error) {
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		return nil, fmt.Errorf("Failed to create pubsub client: %w", err)
	}

	topic := client.Topic(topicId)
	sub := client.Subscription(subId)

	ctx, cancel := context.WithCancel(ctx)

	bus := &PubSubBus{
		client:        client,
		topic:         topic,
		subscription:  sub,
		ctx:           ctx,
		subCancel:     cancel,
		orderHandlers: []OrderCreatedEventHandler{},
	}

	go bus.startReceiving()

	return bus, nil
}

func (b *PubSubBus) PublishOrderCreated(event *pb.OrderCreatedEvent) error {
	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("Failed to marshal order event: %w", err)
	}

	result := b.topic.Publish(b.ctx, &pubsub.Message{
		Data: data,
	})
	_, err = result.Get(b.ctx)
	return err
}

func (b *PubSubBus) SubscribeToOrderCreated(handler OrderCreatedEventHandler) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.orderHandlers = append(b.orderHandlers, handler)
}

// startReceiving pulls messages and dispatches to handlers
func (b *PubSubBus) startReceiving() {
	log.Println("[PubSubBus] Starting to receive messages...")

	err := b.subscription.Receive(b.ctx, func(ctx context.Context, msg *pubsub.Message) {
		var evt pb.OrderCreatedEvent
		if err := json.Unmarshal(msg.Data, &evt); err != nil {
			log.Printf("Yabai! Failed to unmarshal OrderCreatedEvent: %v", err)
			msg.Nack()
			return
		}

		log.Printf("Yatta! Received OrderCreatedEvent: %+v", &evt)
		b.mu.Lock()
		handlers := append([]OrderCreatedEventHandler{}, b.orderHandlers...)
		b.mu.Unlock()

		for _, handler := range handlers {
			handler(&evt)
		}

		msg.Ack()
	})

	if err != nil {
		log.Printf("Yabai! Pub/Sub receive error: %v", err)
	}
}

func (b *PubSubBus) Stop() {
	b.subCancel()
	b.topic.Stop()
	_ = b.client.Close()
}
