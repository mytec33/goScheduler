package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"sync"
	"time"

	config "github.com/mytec33/goScheduler/src/config"
)

var wg sync.WaitGroup

func runTask(action config.Action) {
	defer wg.Done()

	log.Printf("Running task: %s %v\n\n", action.Executable, action.Arguments)

	cmd := exec.Command(action.Executable, action.Arguments...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error running task %s: %v\n", action.Executable, err)
	}
	fmt.Printf("Output of %s: %s\n", action.Executable, string(output))
}

func SleepM(pMilliSec int) { // sleep millisecounds
	time.Sleep(time.Duration(pMilliSec) * time.Millisecond)
}

func main() {
	var configArg string
	flag.StringVar(&configArg, "config", "", "which configuration to load")
	flag.Parse()

	if configArg == "" {
		flag.Usage()
	}

	for {
		var configuration config.Configuration
		configuration.ReadConfiguration(configArg)
		if len(configuration.Tasks) < 1 {
			log.Fatal("No tasks found. Quitting.")
		}
		configuration.PrettyPrint()

		for _, task := range configuration.Tasks {
			fmt.Printf("Task: %v\n", task.Action)
			wg.Add(1)
			go runTask(task.Action)
		}

		SleepM(10000)
	}
}
