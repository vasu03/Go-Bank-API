// Declaring the main package
package main

import (
	"log"
)

// Importing required modules

// Defining the main() method
func main()  {
	// Create a store interface to DB
	store, err := NewPostgreStore()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to DB
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	// Starting the API server 
	server := NewAPIServer(":3000", store)
	server.StartServer()

}