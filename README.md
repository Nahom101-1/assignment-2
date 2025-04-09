# Assignment-2 Project Setup

## 👥 Group Members

- **Nahom**: [github.com/Nahom101-1](https://github.com/Nahom101-1)
- **Fredrik**: [github.com/fredrikandreas](https://github.com/fredrikandreas)
- **Tim**:  [github.com/TimHarseth](https://github.com/TimHarseth)
- **Eirik**: [github.com/eirikm02](https://github.com/eirikm02)

---

## 📁 Project Structure
 Below is an overview of the current folder structure:
```
assignment-2/
│
├── cmd/                     # Entry point of the app
│   └── server/
│       └── main.go
│
├── config/                  # Configuration files
│
├── internal/                # Core application logic
│   ├── constants/           # API constants, events, URLs
│   ├── handlers/            # API route handlers
│   │   ├── dashboard/
│   │   ├── notifications/
│   │   └── registrations/
│   ├── models/              # Data/struct models
│   ├── services/            # other logic
│   │   ├── fetch/
│   │   └── notifications/
│   ├── storage/             # Storage logic
│   │   └── firebase.go
│   └── utils/               # Utility helper functions
│
├── static/                  # Static files (html, css, etc)
│   └── index.html
│
├── tests/                   # Extra test files
│
├── .github/                 # GitHub Actions workflows
│   └── workflows/
│       └── gitlab-sync.yml
│
├── Dockerfile               # Docker configuration
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
   go run ./cmd/server
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


## Data Source

This project uses weather data provided by [Open-Meteo](https://open-meteo.com/).  
The data is licensed under [Creative Commons Attribution 4.0 (CC-BY 4.0)](https://creativecommons.org/licenses/by/4.0/).
