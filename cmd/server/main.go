package main

import (
	"context"
	"flag"
	"fmt"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/vmware/vmware-go-kcl/logger"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "The port number for the service.")
	flag.Parse()

	config := logger.Configuration{
		EnableConsole:     true,
		ConsoleLevel:      logger.Info,
		ConsoleJSONFormat: false,
	}
	log := logger.NewLogrusLoggerWithConfig(config)

	ctx := cloudevents.ContextWithTarget(context.Background(), "http://localhost:8080/")

	p, err := cloudevents.NewHTTP()
	if err != nil {
		log.Fatalf("failed to create protocol: %s", err.Error())
	}

	c, err := cloudevents.NewClient(p, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	log.Infof("server listen on :8080")
	if err = c.StartReceiver(ctx, receive); err != nil {
		log.Fatalf("Error in start server.")
	}
}

func receive(ctx context.Context, event cloudevents.Event) {
	fmt.Printf("%s", event)
}