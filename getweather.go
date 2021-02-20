package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	// ErrorSendRequest - Error if can't perform get
	ErrorSendRequest = errors.New("GET request insuccessfull")
	// ErrorWrongCity Error if wrong city name
	ErrorWrongCity = errors.New("Wrong city name")
	// ErrorWringAPIKey Error if wrong API key
	ErrorWringAPIKey = errors.New("Wrong API key")
	// ErrorJSONParse Error if can't perform parsing
	ErrorJSONParse = errors.New("Can't parse JSON")
	// ErrorReadFile  Error if can't pread a file
	ErrorReadFile = errors.New("Can't read file")
)

// HTTPStatusError Server as stub
type HTTPStatusError struct {
	status int
}

func (e *HTTPStatusError) Error() string {
	return fmt.Sprintf("code: %d", e.status)
}

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

func httpGetWeather(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%w - %s", ErrorSendRequest, err.Error())
	}
	switch resp.StatusCode {
	case 404:
		return nil, fmt.Errorf("%w - %s", ErrorWrongCity, err.Error())
	case 401:
		return nil, fmt.Errorf("%w - %s", ErrorWringAPIKey, err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("The error occurs with reading file: %w - %s", ErrorReadFile, err.Error())
	}
	return body, nil
}

func fileGetWeather(fileDescriptor string) ([]byte, error) {
	jsonData, err := ioutil.ReadFile(fileDescriptor)
	if err != nil {
		return nil, fmt.Errorf("The file %s cannot be opened, info: %w - %s", ErrorReadFile, err.Error())
	}
	return jsonData, nil
}

func readJSON(body []byte) (currentWeather, error) {
	var data httpResponse
	err := json.Unmarshal(body, &data)
	if err != nil {
		return currentWeather{}, fmt.Errorf("The data getting by GET HTTP request cannot be read to JSON: %w - %s", ErrorJSONParse, err.Error())
	}
	return data.Main, nil
}
