package main

import (
	"github.com/mitchellh/mapstructure"
)

type Result = map[string]interface{}
type HelloMsg struct {
	Type    string `json:"type"`
	Version string `json:"version"`
	Agent   string `json:"agent,omitempty"`
}

// Casts request data to the appropriate type and routes the requst
func Route(req map[string]interface{}) (error, Result) {
	reqType, ok := req["type"]
	if !ok {
		return InvalidMessage, nil
	}

	switch reqType {
	case "hello":
		var hello HelloMsg
		mapstructure.Decode(req, &hello)
		return HelloController(hello)
	}

	return InvalidMessage, nil
}
