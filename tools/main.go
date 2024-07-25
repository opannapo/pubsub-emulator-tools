package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/opannapo/pubsub-emulator-tools/tools/compose"
	"github.com/opannapo/pubsub-emulator-tools/tools/subscription"
	"github.com/opannapo/pubsub-emulator-tools/tools/topic"
	"google.golang.org/api/option"
	"log"
	"os"
	"strings"
)

const help = `
Action List
[*] Emulator
	[1] setup-emulator-compose
[*] Application
	[2] topic-create
	[3] topic-list
	[4] topic-delete
	[5] subscription-create
	[6] subscription-list
`

var projectID string
var act = map[int]func(){}

func main() {
	ctx := context.Background()
	projectID = os.Getenv("PUBSUB_PROJECT_ID")
	if projectID == "" {
		panic("Empty PUBSUB_PROJECT_ID")
	}

	opt := []option.ClientOption{
		option.WithoutAuthentication(), //option.WithEndpoint("0.0.0.0:8085"),
	}
	psClient, err := pubsub.NewClient(ctx, projectID, opt...)
	if err != nil {
		log.Panicf("Failed to create client: %v", err)
	}
	defer psClient.Close()

	act = map[int]func(){
		1: func() { compose.Setup(ctx) },
		2: func() { topic.Create(ctx, *psClient) },
		3: func() { topic.List(ctx, *psClient) },
		4: func() { topic.Delete(ctx, *psClient) },
		5: func() { subscription.Create(ctx, *psClient) },
		6: func() { subscription.List(ctx, *psClient) },
	}

	cliDisplayOpening()
}

func cliDisplayOpening() {
	fmt.Println(help)
	in := readerActionNumber()
	handlingAction(in)
}
func cliDisplayAsk() {
	fmt.Print("Continue [Y/N] ? ")
	var in string
	_, err := fmt.Scanf("%s", &in)
	if err != nil {
		panic(err)
	}

	if in == strings.ToLower("y") {
		fmt.Printf("\x1bc")
		cliDisplayOpening()
	} else {
		fmt.Printf("\x1bc")
		os.Exit(1)
	}
}
func readerActionNumber() int {
	fmt.Print("Action number : ")
	var in int
	_, err := fmt.Scanf("%d", &in)
	if err != nil {
		return 0
	}

	return in
}
func handlingAction(in int) {
	defer func() {
		cliDisplayAsk()
	}()

	val, ok := act[in]
	if ok {
		val()
	} else {
		fmt.Printf("\x1bc")
		cliDisplayOpening()
	}
}
