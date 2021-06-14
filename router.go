package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type HelloMsg struct {
	Type    string `json:"type"`
	Version string `json:"version"`
	Agent   string `json:"agent,omitempty"`
}

// Casts request data to the appropriate type and routes the requst
func Route(req map[string]interface{}) (error, map[string]interface{}) {
	reqType, ok := req["type"]
	if !ok {
		fmt.Println("Type key does not exist in map")
	}

	switch reqType {
	case "hello":
		var hello HelloMsg
		mapstructure.Decode(req, &hello)
		return HelloController(hello)
	}

	return nil, nil
}
