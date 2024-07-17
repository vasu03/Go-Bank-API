// // File containing the API http Server and http Handlers // //

// Declaring the main package
package main

// Importing required modules
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Creating a structure for server
type APIServer struct {
	listenAddress string
	store         Storage
}

// Creating a structure for the Errors within API
type ApiError struct {
	Error string
}

// Defining a function signature for the handlers in use
type apiFunc func(http.ResponseWriter, *http.Request) error

// Creating a integrator for putting apiFunc() into our handlers
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

// Function that enables writting the JSON data (write/encode anything as JSON)
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	// modify the headers
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

// Creating a Server for API
func NewAPIServer(listenAddress string, store Storage) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		store:         store,
	}
}

// Starting the Server
func (s *APIServer) StartServer() {
	router := mux.NewRouter()

	// make a route and a http handler func
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccountByID))

	//a debug print statement
	log.Println("\nAPI server running on port ", s.listenAddress)

	// start listening & serving at given addr/port
	http.ListenAndServe(s.listenAddress, router)
}



// Handler function for managing the Accounts
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	// according to type of request, dispatching the handler

	// for request of type GET
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}

	// for request of type POST
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}

	// for request of type  DELETE
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("method of type %s not allowed", r.Method)
}

// Handler function for Getting all the Accounts
func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	// fetch the accounts details from db
	accounts, err := s.store.GetAccounts()
	
	if err != nil {
		return err
	}

	// return the fetched data
	return WriteJSON(w, http.StatusOK, accounts)
}


// Handler function for Getting the Accounts based on ID
func (s *APIServer) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	// db.get(id) stuff
	// account := NewAccount("Vasu", "Makadia")
	fmt.Println("id = ", id)

	return WriteJSON(w, http.StatusOK, &Account{})
}


// Handler function for Creating the Accounts
func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	// initialise a new request to create the account
	createAccountReq := new(CreateAccountRequest)

	// Decode the data from the body of request
	if err := json.NewDecoder(r.Body).Decode(createAccountReq); err != nil {
		return err
	}
	
	// Create a new account by passing in the required data 
	account := NewAccount(createAccountReq.FirstName, createAccountReq.LastName)
	// store the created account into DB
	if err := s.store.CreateAccount(account); err != nil {
		return err;
	}

	return WriteJSON(w, http.StatusOK, account)
}


// Handler function for Deleting the Accounts
func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}


// Handler function for transfering money between accounts
func (s *APIServer) handleTransferMoney(w http.ResponseWriter, r *http.Request) error {
	return nil
}
