# go-gorm-db-transactions
DB Transaction Implementation for multi-layered (Controller, Service, Repository layers) Golang project.

### Uses :-
- GORM  (As ORM library for the DB interactions)
- Go Fiber (As REST Framework)
- MySQL Driver (MySQL as Relational DB choosen for the Demo)

This repository showcases the implementation of the database transactions in a multi-layered project having Controller, Service, Repository, etc. layers.
`DBTransactionMiddleware` is the middleware which takes care of DB transaction functionality.

### How to Run :-
- In Terminal, Run `go run main.go`
- Hit the endpoints through the REST client of your choice (e.g. Postman)
- e.g. To create new users
  ```
  // Create first user
  POST localhost:3000/users
  {
    "email": "test@mail.com",
    "wallet": 1000
  }
  
  // Creates 2nd user
  POST localhost:3000/users
  {
    "email": "testnew@mail.com",
    "wallet": 2000
  }
  ```
- e.g. To carry money transfer between 2 users
  ```
  POST localhost:3000/transfer-money
  {
    "receiver": 1,
    "giver": 2,
    "amount": 500
  }
  ```

Note :-
  - By Default, the transaction should fail in decrementing the money from *Giver's* account
  - To make it working, Go to `user_repository.go` and uncomment the last line by first commenting the error line above it. 
