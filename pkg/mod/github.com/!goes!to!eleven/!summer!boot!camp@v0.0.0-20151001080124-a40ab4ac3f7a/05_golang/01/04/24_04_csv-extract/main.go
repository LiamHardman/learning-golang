package main

import (
	"fmt"
	"os"
	"log"
	"encoding/csv"
)

func main() {
	// 0. get csv file - http://statetable.com/
	// 1. open and read csv file
	// 2. parse csv
	// 3. show results

	f, err := os.Open("../resources/state_table.csv")
	if err != nil {
		log.Fatalln("couldn't open csv file", err.Error())
	}

	myReader := csv.NewReader(f)

	records, err := myReader.ReadAll()
	if err != nil {
		log.Fatalln("couldn't read it!", err.Error())
	}

	myStates := make(map[string]string)

	for _, value := range records {
		myStates[value[2]] = value[1]
	}

	lookup := os.Args[1]

	fmt.Println(myStates[lookup])

	// run go install
	// run program from command line (cli)
	// add an argument <STATE ABBREVIATION>

}
