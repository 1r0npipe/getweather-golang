package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var city = "orenburg"
var port = "8080"
var token = "38c3a66510637c67309b55afc428f4dc"
var url = "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + token

type simpleWeather struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
}

func main()  {

	http.HandleFunc("/", simpleWebJSON)

    http.ListenAndServe(":" + port, nil)
}

func simpleWebJSON(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func httpGetWeather (url string, target interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("The error occurs %v", err)
	}
	defer req.Body.Close()	
	return json.NewDecoder(req.Body).Decode(target)
}