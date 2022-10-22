package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type container struct {
	Name  string              `yaml:"name"`
	Image string              `yaml:"image"`
	Ports []map[string]string `yaml:"ports"`
}

type pod struct {
	ApiVersion string            `yaml:"apiVersion"`
	Kind       string            `yaml:"kind"`
	Metadata   map[string]string `yaml:"metadata"`
	Spec       struct {
		Containers []container
	}
}

func main() {
	// open the yaml file
	data, err := os.ReadFile("pods.yaml")
	printErr("read file", err)

	m := &pod{}

	err = yaml.Unmarshal(data, &m)
	printErr("unmarshal pods.yaml", err)

	fmt.Println(m)

	data1, err := json.MarshalIndent(m, "", "   ")
	printErr("print struct", err)
	fmt.Println(string(data1))
}
