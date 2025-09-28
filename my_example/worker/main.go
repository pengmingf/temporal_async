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

	w := worker.New(c, string(consts.Worker2), worker.Options{})

	// w.RegisterWorkflow(my_example.SayHelloWorldWorkflow)
	// w.RegisterActivity(my_example.Greet)
	// w.RegisterActivity(my_example.Greet2)
	// w.RegisterActivity(my_example.DomainResolveTask)

	// 注册新的workerflow 和 activity
	w.RegisterWorkflow(my_example.FiveActivityWorkflow) // 注册新的Workflow
	w.RegisterActivity(my_example.ActivityA1)
	w.RegisterActivity(my_example.ActivityA2)
	w.RegisterActivity(my_example.ActivityA3)
	w.RegisterActivity(my_example.ActivityA4)
	w.RegisterActivity(my_example.ActivityA5)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
