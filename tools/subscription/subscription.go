package subscription

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/opannapo/pubsub-emulator-tools/shared"
	"google.golang.org/api/iterator"
	"log"
	"time"
)

func Create(ctx context.Context, psClient pubsub.Client) {
	shared.IOClear()
	subName := shared.IOStdinRead("[subscription-create] subscription name  : ")
	topicName := shared.IOStdinRead("[subscription-create]	topic name  : ")

	topic, err := psClient.CreateTopic(ctx, topicName)
	if err != nil {
		log.Fatalf("Failed to create topic: %v", err)
	}
	topic.PublishSettings = pubsub.PublishSettings{
		DelayThreshold:            0,                            //default
		CountThreshold:            0,                            //default
		ByteThreshold:             0,                            //default
		NumGoroutines:             0,                            //default
		Timeout:                   0,                            //default
		BufferedByteLimit:         0,                            //default
		FlowControlSettings:       pubsub.FlowControlSettings{}, //default
		EnableCompression:         false,                        //default
		CompressionBytesThreshold: 0,                            //default
	}
	topic.EnableMessageOrdering = false //default

	sub, err := psClient.CreateSubscription(ctx, subName, pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: 20 * time.Second,
	})
	if err != nil {
		log.Fatalf("CreateSubscription: %+v", err)
	}
	fmt.Printf("Created subscription: %v\n", sub)
}

func List(ctx context.Context, psClient pubsub.Client) {
	fmt.Printf("\x1bc")

	var subs []*pubsub.Subscription
	it := psClient.Subscriptions(ctx)
	for {
		s, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("next: %+v \n", err)
		}

		subs = append(subs, s)
	}

	for _, t := range subs {
		log.Println(t)
	}

	return
}
