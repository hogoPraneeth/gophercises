package quiz1

import (
	"fmt"
	"encoding/csv"
	"io"
	"log"
	"os"
	"bufio"
)
type Problems struct {
	question string
	solution int
}

func Run(config string) {
	fmt.Println("hello from quiz")
	fmt.Println(config)
	csvFile, err := os.Open(config)
	if err != nil {
		log.Fatal(err)
	}
    reader := csv.NewReader(bufio.NewReader(csvFile))
//    var probs []Problems
	for {
		line, err := reader.Read()
		if err == io.EOF {
			log.Fatal(err)
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(line[0])
	}
}


