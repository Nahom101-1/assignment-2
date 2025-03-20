# Assignment-2 Project Setup

# group members 
  Nahom : https://github.com/Nahom101-1

  Fredrik: https://github.com/fredrikandreas

  Tim:  legg til egen github
## Project Structure
 initial setup of our Go project. Below is an overview of the current folder structure:

```
assignment-2/
│── internal/              # Core logic
│   ├── handlers/          # API route handlers
│   ├── models/            # Data/struct models
│   ├── services/          # logic endpoints, getcities, getpopulation etc..
│── tests/                 # Unit tests
│── utils/                 # Utility functions
│   ├── check_status.go    # Status check utility
│   ├── get_request.go     # Helper for GET requests
│   ├── handle_ServiceError.go # Error handling utilities
│   ├── post_request.go    # Helper for POST requests
│   ├── read_body.go       # Reads request bodies
│   ├── server_port.go     # Handles server port configuration
│── go.mod                 # Go module dependencies
│── main.go                # Entry point of the application
│── README.md              # Project documentation
```

## How to get started Project
1. Clone the repository:
   ```sh
   git clone <repo-url>
   cd assignment-2
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Run the server:
   ```sh
   go run main.go
   ```
