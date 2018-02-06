package quiz1

import (
	"fmt"
	"encoding/csv"
	"strings"
	"io"
	"log"
	"os"
)

func Run() {
	fmt.Println("hello from quiz")
	file := os.Getenv("GOPATH") + "/config/problems.csv"
	r := csv.NewReader(strings.NewReader(file))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}
