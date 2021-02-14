package main

import "fmt"

func main() {
	
	envVar := readEnvVars()
	token := envVar[0]
	city := envVar[1]
	//port := envVar[2]
	var url = "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + token
	response := httpGetWeather(url)
	temperature := readJSON(response)
	fmt.Println(temperature)
}