package main

import (
	"flag"
	"fmt"
)

func main() {
	file := "./data/weatherOffline.json"
	envVar := readEnvVars()
	token := envVar[0]
	city := envVar[1]
	unit := envVar[2]
	flag.Parse()
	if *cityName != "" {
		city = *cityName
	}
	var url = "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + token + "&units=" + unit
	response := httpGetWeather(url)
	temperature := readJSON(response)
	fmt.Printf("Online temp for %s: %.2f, humidity: %d%%\n", city, temperature.Temp, temperature.Humidity)
	if *methodType == "offline" {
		city = "Orenburg"
		fmt.Println("Read data from Json file")
		response = fileGetWeather(file)
		temperature = readJSON(response)
		fmt.Printf("Offline temp for %s: %.2f (Kelvin), humidity: %d%%\n", city, temperature.Temp, temperature.Humidity)
	}
	startTestServer()
}
