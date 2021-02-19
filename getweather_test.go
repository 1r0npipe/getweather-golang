package main

import (
	"testing"
	"net/http"
)

func Test_httpGetWeather(t *testing.T) {
	tests := []struct {
		url string
		codeRequest int
		codeWant int
	}{
		// it is planning to be a code got and want here like {404,404}
		{"localhost:8080/success", 200, 200}, 
	}
	for _, tt := range tests {
		if resp, _ := http.Get(tt.url); resp.StatusCode != tt.codeWant {
				t.Errorf("httpGetWeather() = %v, want %v", tt.codeRequest, tt.codeWant)
		}
	}
}

