package main

import "flag"

var (
	methodType = flag.String("methodType", "online", "Choose online or offline, online by default")
	unitType = flag.String("unitType", "metric", "Choose unit time, Celcius by default")
	cityName = flag.String("city", "", "Choose city")
)