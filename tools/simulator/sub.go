package simulator

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/opannapo/pubsub-emulator-tools/shared"
	"log"
)

type SubSimulator struct {
	Ctx      context.Context
	PsClient *pubsub.Client
	sub      string
}

func NewSubSimulator(ctx context.Context, psClient *pubsub.Client) (*SubSimulator, error) {
	return &SubSimulator{
		Ctx:      ctx,
		PsClient: psClient,
	}, nil
}

func (s *SubSimulator) StartCli() {
	shared.IOClear()
	sub := shared.IOStdinRead("[start-sub-cli] subscription ID  : ")

	if sub == "" {
		panic("missing subscription ID parameter")
	}
	s.sub = sub

	subscription := s.PsClient.Subscription(s.sub)
	fmt.Println("listening for messages ...")
	err := subscription.Receive(context.Background(), func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("Received message: %s\n", string(msg.Data))
		msg.Ack()
	})
	if err != nil {
		log.Fatalf("Failed to receive messages: %v", err)
	}
}
