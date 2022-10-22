package main

import "log"

func printErr(describeAction string, err interface{}) {
	if err != nil {
		log.Fatalf("Error: unable to%s;\n error message: %v ", describeAction, err)
	}
}
