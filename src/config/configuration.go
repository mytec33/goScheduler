package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Action struct {
	Executable string   `json:"executable"`
	Arguments  []string `json:"arguments"`
}

// A configuration [file] is an array of tasks
type Configuration struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Action      Action `json:"action"`
}

func (c *Configuration) PrettyPrint() {
	pp, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		log.Fatalf("Unable to pretty print configuration: %v", err.Error())
	}
	fmt.Printf("Loaded config:\n%s\n", pp)
}

func (c *Configuration) ReadConfiguration(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err.Error())
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err.Error())
	}

	return nil
}
