# Assignment-2 Project Setup

## 👥 Group Members

- **Nahom**: [github.com/Nahom101-1](https://github.com/Nahom101-1)
- **Fredrik**: [github.com/fredrikandreas](https://github.com/fredrikandreas)
- **Tim**: _(add your GitHub link here)_
- **Eirik**: [github.com/eirikm02](https://github.com/eirikm02)

---

## 📁 Project Structure
 initial setup of our Go project. Below is an overview of the current folder structure:

```
assignment-2/
│── internal/              # Core logic
│   ├── handlers/          # API route handlers
│   ├── models/            # Data/struct models
│   ├── services/          # logic endpoints, getcities, getpopulation etc..
│── tests/                 # Unit tests
│── config/
├   ├── firebase.json
│── utils/                 # Utility functions
│   ├── check_status.go    # Status check utility
│   ├── get_request.go     # Helper for GET requests
│   ├── handle_ServiceError.go # Error handling utilities
│   ├── post_request.go    # Helper for POST requests
│   ├── read_body.go       # Reads request bodies
│   ├── server_port.go     # Handles server port configuration
│   ├── response.go        # Handles writing to the browser
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

## Features & Endpoints
1. Dashboard Registrations (/dashboard/v1/registrations/)
- Register a dashboard configuration (POST)
- Retrieve a specific dashboard configuration (GET)
- Retrieve all dashboard configurations (GET)
- Update a dashboard configuration (PUT)
- Delete a dashboard configuration (DELETE)

2. Dashboard Retrieval (/dashboard/v1/dashboards/)
- Retrieve a populated dashboard (GET)

3. Webhook Notifications (/dashboard/v1/notifications/)
- Register a webhook (POST)
- Retrieve a specific webhook (GET)
- Retrieve all webhooks (GET)
- Delete a webhook (DELETE)

4. Service Status (/dashboard/v1/status/)
- Monitor availability of external APIs and system health (GET)
