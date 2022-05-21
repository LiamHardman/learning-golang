package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	srcFile, err := os.Open("hello.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer srcFile.Close()

	scanner := bufio.NewScanner(srcFile)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(">>>", line)
	}
}


// might not be a complete working file
// this file was just pulled part way through lecture
// https://www.youtube.com/playlist?list=PLSak_q1UXfPqJ6pIpJt2oNJTPx9ByBy6K