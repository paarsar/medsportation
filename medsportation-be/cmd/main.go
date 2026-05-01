package main

import (
	"log"
	"os"

	"medsportation-be" // This assumes the package name in function.go is 'backend' and the module is 'medsportation-be'

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func main() {
	// Register the function
	functions.HTTP("RequestQuote", backend.RequestQuote)

	// Use PORT environment variable, or default to 8080
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	if err := functions.StartServer(port); err != nil {
		log.Fatalf("functions.StartServer: %v\n", err)
	}
}
