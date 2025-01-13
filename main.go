package main

import (
	"fmt"
	"net/http"
	"url-shortner/routes"
)

func main() {
	fmt.Println("Starting URL Shortner App!!")
	routes.RegisterRoutes()

	fmt.Println("Server starting on port localhost:3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Server broken")
	}
}
