package main

import "fmt"

func HelloController(hello HelloMsg) (error, map[string]interface{}) {
	fmt.Println("Inside controller", hello)
	test := map[string]interface{}{"type": "hello", "version": "0.1.2", "agent": "Marabu-Core Client 0.7"}

	return nil, test
}
