package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"strings"
	"time"
	"os"
)

// resp, err := client.Get("https://official-joke-api.appspot.com/random_joke")
// resp, err := client.Get("https://catfact.ninja/fact")
// resp, err := client.Get("https://ipinfo.io")
// resp, err := client.Get("https://api.adviceslip.com/advice")

//curl https://api.coindesk.com/v1/bpi/currentprice.json
// curl http://numbersapi.com/random/year
//curl http://api.open-notify.org/iss-now.json
// curl https://api.nationalize.io?name=michael

// if (err != nil) {
// 	fmt.Printf("Error occured %v\n", err)
// }
// resBody, err := io.ReadAll(resp.Body)
// fmt.Println(string(resBody))

var countryMap = readCountryFile()

func readCountryFile() map[string]string {
	file, err := os.Open("country.json")
	if err != nil {
		fmt.Println("Error occured when loading country data")
		return nil
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error occured when reading country data")
		return nil
	}
	
	var countries map[string]string 
	err = json.Unmarshal(content, &countries)

	if err != nil {
		fmt.Println("Error occured when parsing country data")
		return nil
	}

	return countries
}

func makeHttpRequest(url string) (map[string]interface{}, error) {
	response, err := client.Get(url)

	if (err != nil) {
		fmt.Println("Error occurred, please try again!")
		return nil, err
	}

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error occurred, please try again!")
		return nil, err
	}
	var result map[string]interface{}
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		fmt.Println("Error occurred, please try again!")
		return nil, err
	}
	defer response.Body.Close()

	return result, nil
}
func printResponse(result map[string]interface{}, fields []string) {

	printProgressBar()


	for _, field := range fields {
		value := extractValueFromKey(result, field)
		fmt.Println(value)
	}
}



func printProgressBar() {
	for i := 0; i < 10; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println("")

}


func extractValueFromKey(json map[string]interface{}, field string) string {
	var value interface{}

	if strings.Contains(field, ".") {
		currentObj := json
		keys := strings.Split(field, ".")
		// if it's a nested key, extract nested key
		var key string
		for i := 0; i < len(keys) - 1; i++ {
			key = keys[i]
			value = currentObj[key]
			if value == nil {
				return ""
			} else {
				jsonInner, ok := value.(map[string]interface{})
				if !ok {
					return ""
				}
				currentObj = jsonInner
			}
		}

		// return last nested key
		return fmt.Sprintf("%v", currentObj[keys[len(keys)-1]])

	} else {
		return fmt.Sprintf("%v",  json[field])
	}
}

func tellRandomJoke() {
	result, err := makeHttpRequest("https://official-joke-api.appspot.com/random_joke")

	if err != nil {
		return
	}

	if result["type"] == "error" {
		fmt.Println(result["message"])
	} else {
		fields := []string{"setup", "punchline"}
		printResponse(result, fields)
	}
}

func tellRandomCatFact() {	
	result, err := makeHttpRequest("https://catfact.ninja/fact")
	if err != nil {
		return
	}
	fields := []string{"fact"}
	printResponse(result, fields)
}

func tellRandomYearFact() {
	response, err := client.Get("http://numbersapi.com/random/year")
	printProgressBar()
	if (err != nil) {
		fmt.Println("Error occurred, please try again!")
	}

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error occurred, please try again!")
	}

	fmt.Println(string(resBody))
}

func printCurrentBitCoinPrice() {
	result, err := makeHttpRequest("https://api.coindesk.com/v1/bpi/currentprice.json")

	if err != nil {
		return
	}

	printResponse(result, []string{"bpi.USD.rate"})
}

func printCurrentISSPosition() {
	result, err := makeHttpRequest("http://api.open-notify.org/iss-now.json")

	if err != nil {
		return
	}

	printResponse(result, []string{"iss_position.latitude","iss_position.longitude"})
}

func printIPAddressInformation() {
	result, err := makeHttpRequest("https://ipinfo.io")
	if err != nil {
		return
	}

	printResponse(result, []string{"ip", "city", "country", "org"})
}

func guessMyCountry() {
	var name string
	fmt.Print("Please enter your name: ")
	scanner.Scan()
	name = scanner.Text()

	printProgressBar()
	name = strings.ReplaceAll(name, " ", "%20") // replace space and ACSII character
	url := fmt.Sprintf("https://api.nationalize.io?name=%v", name)

	result, err := makeHttpRequest(url)

	if err != nil {
		return
	}
	topCountries := result["country"].([]interface{})

	if topCountries == nil {
		fmt.Println("No result")
	} else {
		bestResult := topCountries[0].(map[string]interface{})
		countryId := bestResult["country_id"].(string)
		probability := bestResult["probability"].(float64)

		resultMessage := constructResultMessage(countryId, probability)

		fmt.Println(resultMessage)
	}

}

func constructResultMessage(countryId string, probability float64) string {
	percentage := int(math.Ceil(probability * 100))
	countryName := countryMap[countryId]

	return fmt.Sprintf("You have %v%% of being from %v", percentage, countryName)
}


func giveAdvice() {
	result, err := makeHttpRequest("https://api.adviceslip.com/advice")
	if (err != nil) {
		return
	}

	printResponse(result, []string{"slip.advice"})
}
