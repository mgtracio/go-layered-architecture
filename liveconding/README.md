# Live Coding

This project demonstrates my layered architecture approach for the DEUNA live coding session.

## Setup Instructions
   
1. Architecture:

```
/layered_architecture
├── cmd/
│   └── challenge/
│       └── main.go                                             # Main entry point of the application
├── pkg/
│   └── config/                                                 # Configuration-related utilities
├── internal/
│   ├── api/
│   │   └── handler/                                            # HTTP handlers for API routes
│   ├── domain/                                                 # Core domain models (business logic)
│   │   └── character.go, episode.go, info.go, location.go      # Domain models
│   ├── repository/                                             # Repository layer (database operations)
│   │   └── tv_show.go                                          # Tv show repository handling DB operations
│   ├── service/                                                # Business logic layer (service layer)
│   │   └── tv_show.go                                          # Tv show service with business logic
│   ├── router/                                                 # Routes and HTTP route setup
│   │   └── router.go                                           # Router setup
│   └── db/                                                     # DB setup and initialization
│       └── db.go                                               # Initializes DB connection
└── go.mod                                                      # Go modules configuration
```


2. Setting up the Project:

   - Prepare dependencies:
   ```bash
      go mod tidy
   ```

   - Run the app:
   ```bash
      cd cmd/challenge
      go run main.go
   ```

   - Run the app with race detector enabled to verify that there are no data races in the code:
   ```bash
      go run -race main.go
   ```

   - Run tests with coverage report:

   ```bash
      go test -coverprofile=coverage.out ./... && go tool cover -func=coverage.out
   ```

3. Using the endpoint:

   ```bash
   curl -X GET http://localhost:8080/fetch-tv-shows
   ```

   Expected Response
   
   Status Code: 200 OK

   Response Body:
   ```
   {
      "message": "Feeds fetched successfully"
   }
   
   ```
---

**Go Version / Database and ORM**:

- This project uses LevelDB for data storage and GORM as the ORM for database operations.
- This project uses Go 1.23.2. Make sure you have the correct version installed for compatibility.

---
**Author**:

   Marco Guillen (mguillen.developer@gmail.com)
