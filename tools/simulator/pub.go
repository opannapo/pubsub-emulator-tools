package simulator

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/opannapo/pubsub-emulator-tools/shared"
	"io/ioutil"
	"log"
	"net/http"
)

type PubSimulator struct {
	Ctx      context.Context
	PsClient *pubsub.Client
	sub      string
	topic    string
}

func NewPubSimulator(ctx context.Context, psClient *pubsub.Client) (*PubSimulator, error) {
	return &PubSimulator{
		Ctx:      ctx,
		PsClient: psClient,
	}, nil
}

func (p *PubSimulator) StartHttp() {
	shared.IOClear()
	port := shared.IOStdinRead("[start-pub-http] port  : ")
	sub := shared.IOStdinRead("[start-pub-http] subscription ID  : ")
	topic := shared.IOStdinRead("[start-pub-http] topic  : ")

	if port == "" {
		panic("missing port parameter")
	}
	if sub == "" {
		panic("missing subscription ID parameter")
	}
	if topic == "" {
		panic("missing topic parameter")
	}

	p.sub = sub
	p.topic = topic

	http.HandleFunc("/publish", p.publishMessage)
	log.Printf("Starting server on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func (p *PubSimulator) publishMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result := p.PsClient.Topic(p.topic).Publish(p.Ctx, &pubsub.Message{
		Data: body,
	})
	id, err := result.Get(p.Ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to publish message: %v", err), http.StatusInternalServerError)
		return
	}

	ok, _ := fmt.Fprintf(w, "Published a message; msg ID: %v\n", id)
	fmt.Println("success publish message ID:", ok)
}
