package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"strings"
)

var config struct {
	IP   string `json:"IP"`
	Port string `json:"Port"`
}

type Response struct {
	Body   interface{} `json:"body"`
	Error  string      `json:"error"`
	Status int         `json:"status"`
}

func main() {
	// Read the configuration from the JSON file
	configFile, err := os.Open("client.json")
	if err != nil {
		log.Fatalf("Error opening configuration file: %v", err)
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&config); err != nil {
		log.Fatalf("Error reading configuration from file: %v", err)
	}
	// Update with your actual server address
	baseURL := "http://" + config.IP + ":" + config.Port

	tlsConfig := &tls.Config{InsecureSkipVerify: true} // Ignore certificate verification
	client := &fasthttp.Client{TLSConfig: tlsConfig}

	reader := bufio.NewReader(os.Stdin)
	flag := false
	for flag == false {
		fmt.Print("Enter a category (country, state, city, coordinates, exit): ")
		category, _ := reader.ReadString('\n')
		category = strings.TrimSpace(category)

		if category == "country" {
			getCountryNames(client, baseURL)
		} else if category == "exit" {
			flag = true
		} else {
			fmt.Print("Enter a value: ")
			value, _ := reader.ReadString('\n')
			value = strings.TrimSpace(value)

			url := fmt.Sprintf("%s/%s/%s", baseURL, category, value)
			getItemsWithValue(client, url, category)
		}
	}
}

func getCountryNames(client *fasthttp.Client, baseURL string) {
	url := baseURL + "/country"
	response, err := callEndpoint(client, url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var resp Response // Assuming the struct is imported and named "structs"
	err = json.Unmarshal(response, &resp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if resp.Status != fasthttp.StatusOK {
		fmt.Printf("Server returned status code: %d\n", resp.Status)
		return
	}

	countries, ok := resp.Body.([]interface{})
	if !ok {
		fmt.Println("Invalid response format from the server.")
		return
	}

	if len(countries) == 0 {
		fmt.Println("No countries found.")
	} else {
		fmt.Println("Country list:")
		for _, country := range countries {
			countryName, ok := country.(string)
			if !ok {
				fmt.Println("Invalid country name format in the response.")
				continue
			}
			fmt.Println(" -", countryName)
		}
	}
}

func getItemsWithValue(client *fasthttp.Client, url, category string) {
	response, err := callEndpoint(client, url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var resp Response // Assuming the struct is imported and named "structs"
	err = json.Unmarshal(response, &resp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if resp.Status != fasthttp.StatusOK {
		fmt.Printf("Server returned status code: %d\n", resp.Status)
		return
	}

	switch category {
	case "coordinates":
		handleCoordinates(resp.Body)
	default:
		handleItemList(resp.Body)
	}
}

func handleCoordinates(data interface{}) {
	coordinates, ok := data.([]interface{})
	if !ok {
		fmt.Println("Invalid response format from the server.")
		return
	}

	if len(coordinates) == 2 {
		latitude, ok := coordinates[0].(string)
		if !ok {
			fmt.Println("Invalid latitude format in the response.")
			return
		}
		longitude, ok := coordinates[1].(string)
		if !ok {
			fmt.Println("Invalid longitude format in the response.")
			return
		}

		coordinatesString := fmt.Sprintf("%s,%s", latitude, longitude)
		fmt.Printf("%s\n", coordinatesString)
	} else {
		fmt.Println("Invalid coordinates format from the server.")
	}
}

func handleItemList(data interface{}) {
	itemList, ok := data.([]interface{})
	if !ok {
		fmt.Println("Invalid response format from the server.")
		return
	}

	if len(itemList) == 0 {
		fmt.Printf("No items found for the specified category and value.\n")
	} else {
		fmt.Println("Items:")
		for _, item := range itemList {
			itemName, ok := item.(string)
			if !ok {
				fmt.Println("Invalid item name format in the response.")
				continue
			}
			fmt.Println(" -", itemName)
		}
	}
}

func callEndpoint(client *fasthttp.Client, url string) ([]byte, error) {
	statusCode, response, err := client.Get(nil, url)
	if err != nil {
		return nil, err
	}

	if statusCode != fasthttp.StatusOK {
		return nil, fmt.Errorf("Server returned status code: %d", statusCode)
	}

	return response, nil
}
