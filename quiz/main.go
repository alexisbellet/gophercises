package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var counter int8
	// prep the input reader
	var answer string
	filename := "problems.csv"
	input_r := bufio.NewReader(os.Stdin)

	// what is the desired file name? default problems.csv
	fmt.Println("What is the name of the file you wish to import your quiz from (e.g. myfile.csv)? Press Enter to use default (problems.csv).")
	answer, _ = input_r.ReadString('\n')
	if strings.TrimSpace(answer) != "" {
		filename = strings.TrimSpace(answer)
	}

	// open the file
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	r := csv.NewReader(file)

	for {
		// read line from the csv file
		record, err := r.Read()
		if err == io.EOF {
			fmt.Printf("\nQuiz is over, you've answered %d question(s) correctly\n", counter)
			break
		}
		if err != nil {
			panic(err)
		}

		// prompt the user with the question and wait for the answer
		fmt.Printf("Question: %s? \n", record[0])
		answer, _ = input_r.ReadString('\n')

		// check for correctness
		if strings.TrimSpace(answer) == record[1] {
			counter = counter + 1
		}
	}
}
