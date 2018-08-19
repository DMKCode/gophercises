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
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	// run in terminal to test flag: go run main.go -file="../problems2.csv"
	filePathPtr := flag.String("file", "problems.csv", "file path as a string")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	csvContents := readFile(*filePathPtr)
	parsedCsv := parseCsv(csvContents)
	startQuiz(parsedCsv, *timeLimit)
}

func readFile(filepath string) string {
	bs, err := ioutil.ReadFile(filepath)

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

func startQuiz(problems [][]string, timeLimit int) {
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	correctAnswerCount := 0
	for _, p := range problems {
		fmt.Println(p[0])
		answerCh := make(chan string)

		go func() {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter answer: ")
			answer, _ := reader.ReadString('\n')
			answer = strings.TrimSpace(answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("You had %d correct answers!", correctAnswerCount)
			return
		case answer := <-answerCh:
			if answer == p[1] {
				fmt.Println("correct!")
				correctAnswerCount++
			} else {
				fmt.Println("wrong!")
			}
			fmt.Println("")
		}
	}

	fmt.Printf("You had %d correct answers!", correctAnswerCount)
}
