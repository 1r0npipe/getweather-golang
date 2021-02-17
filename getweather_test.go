package main

import (
	"reflect"
	"testing"
)

func Test_httpGetWeather(t *testing.T) {
	tests := []struct {
		url string
		codeRequest string
		codeWant string
	}{
		// it is planning to be a code got and want here like {404,404}
	}
	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			if got := httpGetWeather(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("httpGetWeather() = %v, want %v", got, tt.want)
			}
		})
	}
}
