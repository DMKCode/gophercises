// read csv file
// give quiz
// keep track of given questions right and wrong
// should go to next question immediately after answer to previous

// default csv - problems.csv, use flag to customise filename
// at the end should show totalquestions correct and wrong

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	csvContents := readFile("problems.csv")
	parsedCsv := parseCsv(csvContents)
	startQuiz(parsedCsv)
}

func readFile(filename string) string {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	in := string(bs)

	return in
}

func parseCsv(csvContents string) [][]string {
	r := csv.NewReader(strings.NewReader(csvContents))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records
}

func startQuiz(problems [][]string) {
	correctAnswerCount := 0
	for _, p := range problems {
		fmt.Println(p[0])
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter answer: ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		// checking the type of a variable
		// fmt.Println(reflect.TypeOf(answer))

		if answer == p[1] {
			fmt.Println("correct!")
			correctAnswerCount++
		} else {
			fmt.Println("wrong!")
		}
		fmt.Println("")
	}

	fmt.Printf("You had %d correct answers!", correctAnswerCount)
}
