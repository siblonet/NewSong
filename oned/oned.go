package oned

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Oned() {
	url := "https://storage.googleapis.com/seeme-7a462.appspot.com/newson/newson.js" // Replace this with the actual URL of the JSON data

	// Make an HTTP GET request to fetch the JSON data
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer response.Body.Close()

	// Check if the response status code is 200 (OK)
	if response.StatusCode != http.StatusOK {
		fmt.Printf("Server returned non-OK status: %d\n", response.StatusCode)
		return
	}

	// Decode the JSON data
	var data []map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Print elements from index 48 to 103 (inclusive)
	for i := 48; i < len(data) && i < 50; i++ {
		item := data[i]
		fmt.Printf("index: %d\n", i)

		// Handle the fields dynamically using type assertions or other techniques
		for key, value := range item {
			fmt.Printf("%s: %v\n", key, value)
		}

		fmt.Println()
	}
}
