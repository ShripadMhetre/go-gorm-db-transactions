# go-gorm-db-transactions
DB Transaction Implementation for multi-layered (Controller, Service, Repository layers) Golang project.

### Uses :-
- GORM  (As ORM library for the DB interactions)
- Go Fiber (As REST Framework)
- MySQL Driver

This repository showcases the implementation of the database transactions in a multi-layered project having Controller, Service, Repository, etc. layers.
`DBTransactionMiddleware` is the middleware which takes care of DB transaction functionality.
