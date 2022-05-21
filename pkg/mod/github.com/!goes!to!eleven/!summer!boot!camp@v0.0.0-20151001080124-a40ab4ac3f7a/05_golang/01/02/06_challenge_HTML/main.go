package main

import "fmt"

const miTokm = 1.60934

func main() {
	fmt.Print("Enter a number: ")
	var number float64
	fmt.Scanln(&number)
	fmt.Println("<!DOCTYPE html>")
	fmt.Println("<html>")
	fmt.Println("<head></head>")
	fmt.Println("<body>")
	fmt.Printf("Miles: %f<br>\n", number)
	fmt.Printf("KM: %f<br>\n", number*miTokm)
	fmt.Println("</body>")
	fmt.Println("</html>")
}

/*
at terminal

go install
06_challenge_HTML > /tmp/tmpfile.html

*/
