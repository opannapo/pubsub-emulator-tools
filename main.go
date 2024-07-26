package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"embed"
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

const version = "v1.0.4"

//go:embed  templates/*
var tmplFS embed.FS
var opening = `
             _               _                                _       _                  _              _     
 _ __  _   _| |__  ___ _   _| |__         ___ _ __ ___  _   _| | __ _| |_ ___  _ __     | |_ ___   ___ | |___ 
| '_ \| | | | '_ \/ __| | | | '_ \ _____ / _ | '_ ' _ \| | | | |/ _' | __/ _ \| '_______| __/ _ \ / _ \| / __|
| |_) | |_| | |_) \__ | |_| | |_) |_____|  __| | | | | | |_| | | (_| | || (_) | | |_____| || (_) | (_) | \__ \
| .__/ \__,_|_.__/|___/\__,_|_.__/       \___|_| |_| |_|\__,_|_|\__,_|\__\___/|_|        \__\___/ \___/|_|___/
|_| v:%s  PID:%d

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
		option.WithoutAuthentication(),
	}
	psClient, err := pubsub.NewClient(ctx, projectID, opt...)
	if err != nil {
		log.Panicf("Failed to create client: %v", err)
	}
	defer psClient.Close()

	pubSimulator, _ := simulator.NewPubSimulator(ctx, psClient)
	subSimulator, _ := simulator.NewSubSimulator(ctx, psClient)
	act = map[int]func(){
		1: func() { compose.Setup(ctx, tmplFS) },
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
func generateLine(length int) string {
	return strings.Repeat("-", length)
}
func cliDisplayOpening() {
	pid := os.Getpid()
	fmt.Printf(fmt.Sprintf(opening, version, pid))

	col1Width := 30
	col2Width := 30
	col3Width := 30

	fmt.Printf("%s\n", generateLine(col1Width+col2Width+col3Width))
	fmt.Printf("%-*s %-*s %-*s\n", col1Width, "Emulator", col2Width, "Application", col3Width, "Simulator")
	fmt.Printf("%s\n", generateLine(col1Width+col2Width+col3Width))

	fmt.Printf("%-*s %-*s %-*s\n", col1Width, "[1] setup-emulator-compose", col2Width, "[2] topic-create", col3Width, "[8] start-pub-http")
	fmt.Printf("%-*s %-*s %-*s\n", col1Width, "", col2Width, "[3] topic-list", col3Width, "[9] start-sub-cli")
	fmt.Printf("%-*s %-*s %-*s\n", col1Width, "", col2Width, "[4] topic-delete", col3Width, "")
	fmt.Printf("%-*s %-*s %-*s\n", col1Width, "", col2Width, "[5] subscription-create", col3Width, "")
	fmt.Printf("%-*s %-*s %-*s\n", col1Width, "", col2Width, "[6] subscription-list", col3Width, "")
	fmt.Printf("%-*s %-*s %-*s\n", col1Width, "", col2Width, "[7] subscription-delete", col3Width, "")

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
