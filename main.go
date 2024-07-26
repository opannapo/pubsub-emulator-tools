package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/opannapo/pubsub-emulator-tools/shared"
	"github.com/opannapo/pubsub-emulator-tools/tools/compose"
	"github.com/opannapo/pubsub-emulator-tools/tools/simulator"
	"github.com/opannapo/pubsub-emulator-tools/tools/subscription"
	"github.com/opannapo/pubsub-emulator-tools/tools/topic"
	"google.golang.org/api/option"
	"log"
	"os"
	"strings"
)

const help = `
             _               _                                _       _                  _              _     
 _ __  _   _| |__  ___ _   _| |__         ___ _ __ ___  _   _| | __ _| |_ ___  _ __     | |_ ___   ___ | |___ 
| '_ \| | | | '_ \/ __| | | | '_ \ _____ / _ | '_ ' _ \| | | | |/ _' | __/ _ \| '_______| __/ _ \ / _ \| / __|
| |_) | |_| | |_) \__ | |_| | |_) |_____|  __| | | | | | |_| | | (_| | || (_) | | |_____| || (_) | (_) | \__ \
| .__/ \__,_|_.__/|___/\__,_|_.__/       \___|_| |_| |_|\__,_|_|\__,_|\__\___/|_|        \__\___/ \___/|_|___/
|_| v1.0.3


Action List :
[*] Emulator
	[1] setup-emulator-compose
[*] Application
	[2] topic-create
	[3] topic-list
	[4] topic-delete
	[5] subscription-create
	[6] subscription-list
	[7] subscription-delete
[*] Simulator
	[8] start-pub-http
	[9] start-sub-cli
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

	pubSimulator, _ := simulator.NewPubSimulator(ctx, psClient)
	subSimulator, _ := simulator.NewSubSimulator(ctx, psClient)
	act = map[int]func(){
		1: func() { compose.Setup(ctx) },
		2: func() { topic.Create(ctx, *psClient) },
		3: func() { topic.List(ctx, *psClient) },
		4: func() { topic.Delete(ctx, *psClient) },
		5: func() { subscription.Create(ctx, *psClient) },
		6: func() { subscription.List(ctx, *psClient) },
		7: func() { subscription.Delete(ctx, *psClient) },
		8: func() { pubSimulator.StartHttp() },
		9: func() { subSimulator.StartCli() },
	}

	cliDisplayOpening()
}
func cliDisplayOpening() {
	fmt.Println(help)
	in, err := shared.IOStdinScan[int]("Action number : ")
	if err != nil {
		panic(err)
	}
	handlingAction(in)
}
func cliDisplayAsk() {
	in, err := shared.IOStdinScan[string]("Continue [Y/N] ? : ")
	if err != nil {
		panic(err)
	}

	if in == strings.ToLower("y") {
		shared.IOClear()
		cliDisplayOpening()
	} else {
		fmt.Printf("\x1bc")
		os.Exit(1)
	}
}
func handlingAction(in int) {
	defer func() {
		cliDisplayAsk()
	}()

	val, ok := act[in]
	if ok {
		val()
	} else {
		shared.IOClear()
		cliDisplayOpening()
	}
}
