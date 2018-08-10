// read csv file
// give quiz
// keep track of given questions right and wrong
// should go to next question immediately after answer to previous

// default csv - problems.csv, use flag to customise filename
// at the end should show totalquestions correct and wrong

package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	csvContents := readFile("problems.csv")
	parsedCsv := readCsv(csvContents)

	fmt.Println(parsedCsv)
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

func readCsv(csvContents string) [][]string {
	r := csv.NewReader(strings.NewReader(csvContents))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records
}
