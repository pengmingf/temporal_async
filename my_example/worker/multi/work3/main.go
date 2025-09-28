package main

import (
	"log"
	"test1/my_example/consts"

	"test1/my_example"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, string(consts.Worker1), worker.Options{})

	w.RegisterWorkflow(my_example.SayHelloWorldWorkflow)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
