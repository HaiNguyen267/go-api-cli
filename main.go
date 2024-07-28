package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var tr = &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}

var client = &http.Client{Transport: tr}

var scanner = bufio.NewScanner(os.Stdin)
func main() {

	fmt.Println("Welcome back, let's do something fun together!")
	var option int64
	for true {
		option = askUserOption()

		switch option {
		case 1:
			tellRandomJoke()
		case 2:
			tellRandomCatFact()
		case 3:
			tellRandomYearFact()
		case 4:
			printCurrentBitCoinPrice()
		case 5:
			printCurrentISSPosition()
		case 6:
			printIPAddressInformation()
		case 7:
			guessMyCountry()
		case 8:
			giveAdvice()
		case 0:
			return
		}

	}

	fmt.Println("Goodbye!")

}

func askUserOption() int64 {

	fmt.Println("\n-------------------| Golang |-------------------")
	fmt.Println("1. Tell me a random joke")
	fmt.Println("2. Tell me a random fact about cat")
	fmt.Println("3. Tell me a fact of a random year")
	fmt.Println("4. Current price of Bitcoin")
	fmt.Println("5. Current position of International Space Station (ISS)")
	fmt.Println("6. Information about my IP address")
	fmt.Println("7. Guess my country")
	fmt.Println("8. Give me an advice")
	fmt.Println("0. Exit")

	fmt.Print("Option: ")
	scanner.Scan()
	option, err := strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 64)

	if err != nil || option < 0 || option > 8 {
		fmt.Println("Invalid option, please choose from 0 to 8")
		return askUserOption()
	}
	return option
}
