package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type httpResponse struct {
	Main currentWeather `json:"main"`
}

type currentWeather struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

func readEnvVars() []string {
	var result []string
	city := os.Getenv("CITY")
	token := os.Getenv("TOKEN")
	unit := os.Getenv("WEATHER_UNIT")
	if token == "" {
		log.Fatal("Token has not been recoginzed, exit...")
	}
	result = append(result, token)
	if city == "" {
		city = "orenburg"
	}
	result = append(result, city)
	if unit == "" {
		unit = "metric"
	}
	result = append(result, unit)

	return result
}


func httpGetWeather(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("The error occurs %v", err)
	}
	switch resp.StatusCode {
	case 404:
		log.Fatal("Wrong city name, the city is not found")
	case 401:
		log.Fatal("Wrong API key (token), please check your API ID")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("The error occurs with reading JSON: %v", err)
	}
	return body
}

func fileGetWeather (fileDescriptor string) []byte {
	jsonData, err := ioutil.ReadFile(fileDescriptor)
	if err != nil {
		log.Fatalf("The file %s cannot be opened, exit...", fileDescriptor)
	}
	return jsonData
}

func readJSON (body []byte) currentWeather {
	var data httpResponse
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("The data getting by GET HTTP request cannot be read to JSON: %v", err)
	}
	return data.Main
}