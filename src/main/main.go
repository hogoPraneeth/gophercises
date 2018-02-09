package main

import (
	"os"
	"fmt"
	"quiz1"
	"go/build"
)


func main() {
	// default
	appBase := os.Args[0] + "/../../";
	config :=  appBase + "/config/problems.csv"
	fmt.Println(appBase)
	if len(os.Args) > 3 && os.Args[2] == "-c" {
		config = os.Args[3]
	}
	if len(os.Args) < 2 {
		fmt.Println("USAGE: ARGUMENTS::\n\t -h : HELP\n\t 1: Quiz1")
		return
	}
	switch os.Args[1] {
	case "-h":
		fmt.Println("ARGUMENTS::\n\t -h : HELP\n\t 1: Quiz1")
		break
	case "-1":
		quiz1.Run(config)
		break
	default:
		fmt.Println("USAGE: ARGUMENTS::\n\t -h : HELP\n\t 1: Quiz1")
	}
	fmt.Println(config)
}