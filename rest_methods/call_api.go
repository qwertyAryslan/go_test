package call_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func GetMethod() {
	url := "https://api.restful-api.dev/objects"

	//request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	//response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	fmt.Println(string(body))
}

func GetMethodById(id int) {
	url := "https://api.restful-api.dev/objects?id=" + strconv.Itoa(id)

	//request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	//response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	fmt.Println(string(body))
}

func PostMethod(name string, year int, price float32, cpu string, hd string) {
	url := "https://api.restful-api.dev/objects"

	// Create the data to be sent in the POST request
	data := map[string]interface{}{
		"name": name,
		"data": map[string]interface{}{
			"year":           year,
			"price":          price,
			"CPU model":      cpu,
			"Hard disk size": hd,
		},
	}

	// Convert data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to convert data to JSON: %v", err)
	}

	// Create a new POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Print the response status
	fmt.Println("Response Status:", resp.Status)

	// Optionally, read and print the response body
	responseBody := new(bytes.Buffer)
	responseBody.ReadFrom(resp.Body)
	fmt.Println("Response Body:", responseBody.String())
}

func PutMethod(name string, year int, price float32, cpu string, hd string, color string) {
	url := "https://api.restful-api.dev/objects/7"

	// Create the data to be sent in the POST request
	data := map[string]interface{}{
		"name": name,
		"data": map[string]interface{}{
			"year":           year,
			"price":          price,
			"CPU model":      cpu,
			"Hard disk size": hd,
			"color":          color,
		},
	}

	// Convert data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to convert data to JSON: %v", err)
	}

	// Create a new POST request
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Print the response status
	fmt.Println("Response Status:", resp.Status)

	// Optionally, read and print the response body
	responseBody := new(bytes.Buffer)
	responseBody.ReadFrom(resp.Body)
	fmt.Println("Response Body:", responseBody.String())
}
