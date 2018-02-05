package main

import (
	"os"
	"fmt"
	"quiz1"
)


func main() {
	// defualt
	config := os.Getenv("GOPATH") + "config/problems.csv"
	if len(os.Args) > 3 && os.Args[2] == "-c" {
		config = os.Args[3]
	}
	switch os.Args[1] {
	case "-h":
		fmt.Println("ARGUMENTS::\n\t -h : HELP\n\t 1: Quiz1")
		break
	case "-1":
		quiz1.Run()
		break
	default:
		fmt.Println("USAGE: ARGUMENTS::\n\t -h : HELP\n\t 1: Quiz1")
	}
	fmt.Println(config)
}