package main

import (
	"fmt"
	"net/http"
)

func httpSuccess(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "{ \"code\": 200 }")
}

func httpAPIError(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "{\"code\":401, \"message\": \"Invalid API key.\"}")
}

func httpWrongCity(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "{\"code\":\"404\",\"message\":\"city not found\"}")
}

func startTestServer() {
	http.HandleFunc("/success", httpSuccess)
	http.HandleFunc("/apierror", httpAPIError)
	http.HandleFunc("/wrongcity", httpWrongCity)

	http.ListenAndServe(":8080", nil)
}
