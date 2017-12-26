package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Cmd struct {
	Description string   `json:"description"`
	Cmd         string   `json:"cmd"`
	Args        []string `json:"args"`
}

type Settings struct {
	Port int   `json:"port"`
	Cmds []Cmd `json:"cmds"`
}

func readSettings() Settings {
	file, e := ioutil.ReadFile("config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file))
	var jsonobject Settings
	json.Unmarshal(file, &jsonobject)
	fmt.Printf("Results: %v\n", jsonobject)
	return jsonobject
}
