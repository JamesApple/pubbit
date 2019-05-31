package main

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
)

type pubber struct {
	client *pubsub.Client
	topic  *pubsub.Topic
	ctx    context.Context
}

func (p pubber) pub(message string) {
	p.topic.Publish(p.ctx, &pubsub.Message{Data: []byte(message)})
}

func newPubber(config Config) pubber {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, config.GCPProjectID)
	if err != nil {
		log.Fatal(err)
	}

	topic := client.Topic(config.GCPTopicID)

	return pubber{client: client, topic: topic, ctx: ctx}
}
