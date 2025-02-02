package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// A configuration [file] is an array of tasks
type Configuration struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Name        string `json:"task"`
	Description string `json:"description"`
}

func (c *Configuration) ReadConfiguration(file string) error {
	data, err := os.ReadFile("./configuration.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err.Error())
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err.Error())
	}

	fmt.Printf("Loaded config: %+v\n", c)
	return nil
}
