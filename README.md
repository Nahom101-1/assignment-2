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
## 🚀 Deployment

The service is deployed at:

> **Floating IP:** `10.212.175.52`
> **Deployed URL:** `http://10.212.175.52:80`

### Example Usage (via cURL):
```sh
# Create a new registration
curl -X POST http://10.212.175.52:80/dashboard/v1/registrations/ -H "Content-Type: application/json" -d '{}'

# Retrieve a specific registration
curl http://10.212.175.52:80/dashboard/v1/registrations/{id}

# Retrieve all registrations
curl http://10.212.175.52:80/dashboard/v1/registrations/

# Update a registration
curl -X PUT http://10.212.175.52:80/dashboard/v1/registrations/{id} -H "Content-Type: application/json" -d '{}'

# Partially update a registration (PATCH)
curl -X PATCH http://10.212.175.52:80/dashboard/v1/registrations/{id} -H "Content-Type: application/json" -d '{}'

# Delete a registration
curl -X DELETE http://10.212.175.52:80/dashboard/v1/registrations/{id}

# Retrieve a populated dashboard
curl http://10.212.175.52:80/dashboard/v1/dashboards/{id}

# Manage webhooks (notifications)
curl -X POST http://10.212.175.52:80/dashboard/v1/notifications/ -H "Content-Type: application/json" -d '{}'

# Check service status
curl http://10.212.175.52:80/dashboard/v1/status/
   ```

## 🛠 How to Get Started with the Project

1. **Clone the repository**:
    ```sh
    git clone <repo-url>
    cd assignment-2
    ```

2. **Install dependencies**:
    ```sh
    go mod tidy
    ```

3. **Run the server locally**:
    ```sh
    go run cmd/server/main.go
    ```


## 📚 Features & Endpoints

### 1. Dashboard Registrations (`/dashboard/v1/registrations/`)
- **POST**: Register a dashboard configuration
- **GET**: Retrieve a specific dashboard configuration
- **GET**: Retrieve all dashboard configurations
- **PUT**: Update a dashboard configuration
- **PATCH**: Partially update a dashboard configuration
- **DELETE**: Delete a dashboard configuration

### 2. Dashboard Retrieval (`/dashboard/v1/dashboards/`)
- **GET**: Retrieve a populated dashboard

### 3. Webhooks Management (`/dashboard/v1/notifications/`)
- **POST**: Register a webhook
- **GET**: Retrieve a specific webhook
- **GET**: Retrieve all webhooks
- **PATCH**: Partially update a webhook
- **DELETE**: Delete a webhook

### 4. Service Status (`/dashboard/v1/status/`)
- **GET**: Monitor availability of external APIs and system health

## 🔔 Supported Webhook Events

- `REGISTER`: Triggered when a new dashboard configuration is registered.
- `CHANGE`: Triggered when a dashboard configuration is updated.
- `DELETE`: Triggered when a dashboard configuration is deleted.
- `INVOKE`: Triggered when a dashboard is retrieved.
- `DASHBOARD_VIEW`: Triggered when a dashboard is viewed.
- `STATUS_CHECK`: Triggered when the status endpoint is accessed.


## 👥 Contribution

| Member   | Contributions                                                                        |
|:---------|:-------------------------------------------------------------------------------------|
| **Nahom** | Project setup, `/dashboard`, `/registration`, `/notification`, deployment, debugging |
| **Tim**   | Project setup, `/dashboard`, testing, deployment, debugging, caching                 |
| **Fredrik** | Project setup/structure, `/status`, debugging                                        |
| **Eirik** | Deployment                                                                           |


## ✨ Extra Features Implemented

- **PATCH Support**: Implemented PATCH functionality for both `/registrations/` and `/notifications/` endpoints following best practices.
- **GDP Feature**: Extended dashboard data to include GDP information per country.
- **Advanced Webhook Events**: Added support for additional events: `DASHBOARD_VIEW` and `STATUS_CHECK`.
- **Webhook Management**: Full webhook lifecycle - Register, Retrieve, Update, Delete.
- **Timestamp Handling**: Consistent `lastChange` timestamp updates on all operations.
- **Purging of Cached Information**: Implemented cache purging for outdated requests older than a configured threshold.



## Data Source

This project uses weather data provided by [Open-Meteo](https://open-meteo.com/).  
The data is licensed under [Creative Commons Attribution 4.0 (CC-BY 4.0)](https://creativecommons.org/licenses/by/4.0/).

Repository: [GitHub Repository Link](https://github.com/Nahom101-1/assignment2)
