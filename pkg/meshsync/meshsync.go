package main

import (
	"log"
	"os"
	"time"

	broker "github.com/layer5io/meshery-operator/pkg/broker"
	"github.com/layer5io/meshery-operator/pkg/meshsync/service"
	utils "github.com/layer5io/meshkit/utils/kubernetes"
)

func main() {
	kubeconfig, err := utils.DetectKubeConfig()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	broker, _ := broker.New(broker.NATSKey, "nats://127.0.0.1:4222")

	err = service.Discover(kubeconfig, broker)
	if err != nil {
		log.Printf("Error while discovery: %s", err)
		os.Exit(1)
	}

	service.Informer(kubeconfig, broker)

	err = service.Start(&service.Service{
		Name:      "meshsync",
		Port:      "11000",
		Version:   "v0.0.1-alpha3",
		StartedAt: time.Now(),
	})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
