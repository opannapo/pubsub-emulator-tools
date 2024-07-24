package topic

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/opannapo/pubsub-emulator-tools/shared"
	"google.golang.org/api/iterator"
	"log"
)

func Create(ctx context.Context, psClient pubsub.Client) {
	shared.IOClear()
	param := shared.IOStdinRead("[topic-create] topic name  : ")

	topic, err := psClient.CreateTopic(ctx, param)
	if err != nil {
		log.Fatalf("Failed to create topic: %v", err)
	}

	fmt.Printf("Topic %v created.\n", topic)
}

func List(ctx context.Context, psClient pubsub.Client) {
	fmt.Printf("\x1bc")

	var topics []*pubsub.Topic
	it := psClient.Topics(ctx)
	for {
		topic, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("next: %+v \n", err)
		}
		topics = append(topics, topic)
	}

	for _, t := range topics {
		log.Println(t)
	}
}

func Delete(ctx context.Context, psClient pubsub.Client) {
	shared.IOClear()
	List(ctx, psClient)
	topicName := shared.IOStdinRead("[topic-delete] topic name  : ")

	t := psClient.Topic(topicName)
	if err := t.Delete(ctx); err != nil {
		log.Fatalf("Failed to delete topic: %v", err)
	}
	fmt.Printf("Topic %v deleted.\n", topicName)
}
