
# Assignment

This repository contains the backend implementation. 

## Features

- RESTful API 
- Secure data handling 
- Dockerized setup for easy deployment and scalability.
- Modular code structure for easy maintenance and extensibility.

## Prerequisites

- **Go**: Ensure you have Go installed on your machine (version >= 1.18).
- **Docker**: Install Docker and Docker Compose.


## Project Structure

```
├── database.go        # Database connection and operations
├── docker-compose.yml # Docker Compose configuration
├── Dockerfile         # Dockerfile to build the application
├── go.mod             # Go module dependencies
├── go.sum             # Dependency checksums
├── handlers.go        # API request handlers
├── main.go            # Application entry point
├── models.go          # Data models for the application
├── utils.go           # Utility functions
└── zClient.http       # HTTP client for testing API endpoints
```

## Setup


#### b. Run the Application

Using `docker-compose`:

```bash
docker-compose up
```

#### c. Stop the Application

```bash
docker-compose down
```

### 3. Run Locally (Optional)

1. Ensure you have Go installed and all dependencies resolved.
2. Build and run the application:

```bash
go mod tidy
go run .
```

## API Endpoints

### Base URL

`http://localhost:8080`

### Available Endpoints

- `POST /receipts/process` - Processes a receipt and returns a unique ID.
- `GET /receipts/{id}/points` - Retrieves points for a specific receipt ID.

## Configuration

- Update environment variables in the `.env` file for custom settings.



## Contact

For questions or feedback, feel free to reach out. subhadip Pahari ( Mobile : 657 319 5550 and email: subhadip.apply@gmail.com )

