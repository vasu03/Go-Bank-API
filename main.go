// Declaring the main package
package main

// Importing required modules


// Defining the main() method
func main()  {

	// Starting the API server 
	server := NewAPIServer(":3000")
	server.StartServer()

}