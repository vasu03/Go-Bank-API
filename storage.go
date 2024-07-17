// // An abstract interface to implement the DB storage features // //

// Declaring the main package
package main

// Importing required modules
import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Defining a Storage interface
type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
}

// Defining a structure for DB
type PostgresStore struct {
	db *sql.DB
}

// function to create the PostGres store
func NewPostgreStore() (*PostgresStore, error) {
	// Load the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Grab the connection string
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set.")
	}

	// Openinng a connection with database
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	// Ping the database to check connection status
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// return the db connection
	return &PostgresStore{
		db: db,
	}, nil
}

// function to initialise the connection with postgres
func (s *PostgresStore) Init() error {
	return s.creatAccountTable()
}

// function to create the table for "Account" in the database
func (s *PostgresStore) creatAccountTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS Account (
			id serial primary key,
			firstName varchar(50),
			lastName varchar(50),
			number serial unique,
			balance decimal(10,3),
			createdAt timestamp
	);`

	_, err := s.db.Exec(query)

	return err
}



// // Implementing the db operational functions // //

// function to insert a new account in the database
func (s *PostgresStore) CreateAccount(acc *Account) error {
	// Query to insert a new account record to db
	query := `
		INSERT INTO Account (firstName, lastName, number, balance, createdAt)
		values ($1, $2, $3, $4, $5);
	`
	// implement the query with input parameters for insertion
	res, err := s.db.Query(query, acc.FirstName, acc.LastName, acc.Number, acc.Balance, acc.CreatedAt)

	if err != nil {
		return err
	}
	return nil
}


// function to update an account in the database
func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}


// function to delete an account based on "id" from the database
func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}


// function to gett all accounts from the database
func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	// Query get the all accounts data from db
	query := `select * from Account;`
	rows, err := s.db.Query(query)

	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account := new(Account)
		err := rows.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt)

		if err != nil {
			return nil, err
		}

		// push the obtained account data into our variable
		accounts = append(accounts, account)
	}

	// return the obtained data
	return accounts, nil
}


// function to get a account based on its "id" from database
func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}
