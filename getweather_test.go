package main

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_httpGetWeather(t *testing.T) {
	tests := []struct {
		url         string
		codeRequest int
		codeWant    int
	}{
		// it is planning to be a code got and want here like {404,404}
		{"https://httpstat.us/200", 200, 200},
		{"https://httpstat.us/401", 401, 401},
		{"https://httpstat.us/404", 404, 404},
	}
	for _, tt := range tests {
		resp, _ := http.Get(tt.url)
		if resp, _ := http.Get(tt.url); resp.StatusCode != tt.codeWant {
			t.Errorf("httpGetWeather() = %v, want %v", tt.codeRequest, tt.codeWant)
		}
		fmt.Printf("response: %d\n", resp.StatusCode)
	}
}

func Test_readJSON(t *testing.T) {
	tests := []struct {
		fileName string
		want     currentWeather
	}{
		{"./data/weatherOffline.json", currentWeather{264.15, 256.0, 1018, 73}},
	}
	for _, tt := range tests {
		response, _ := fileGetWeather(tt.fileName)
		temperature, _ := readJSON(response)
		if temperature != tt.want {
			t.Errorf("The issue with input data, got %+v, want %+v", temperature, tt.want)
		}

	}
}
