# Transjakarta Fleet Management System - Backend

A robust, enterprise-grade backend architecture for managing and tracking the Transjakarta bus fleet in real-time. Built using Go following Clean Architecture principles, and fully orchestrated via Docker Compose for seamless deployment and scalability.

---

## 🏗️ System Architecture & Stack

The application is containerized using Alpine-based microservices, ensuring maximum efficiency, minimal resource footprint, and secure communication:

* **Language:** Go (Golang)
* **Architecture Pattern:** Clean Architecture (`Domain`, `Use Case`, `Interface/Adapters`, `Infrastructure`)
* **Database:** PostgreSQL 15 (Alpine) with automated migration routines
* **Message Broker (Queue):** RabbitMQ 3 (Alpine) for asynchronous background event processing
* **MQTT Broker:** Eclipse Mosquitto 2 with on-the-fly dynamic authentication (`pwfile` generation)
* **Geofencing Engine:** Custom high-performance Equirectangular approximation (`CalculateDistanceFast`) for real-time route boundary checking

---

## 📂 Project Structure

```text
.
├── cmd/
│   └── main.go             # Main application entry point
├── internal/               # Clean Architecture layers (Domain, UseCase, Delivery, Repository)
├── migration/              # Database migration scripts and initializers
├── config/                 # Configuration files
├── Dockerfile              # Multi-stage Docker build for Go binary
├── docker-compose.yml      # Multi-container orchestration setup
├── .env                    # Environment variables (Single Source of Truth)
└── README.md
```

---

## ⚙️ Prerequisites

Make sure you have the following installed on your local machine:

* Docker & Docker Compose
* Go 1.21+ *(optional, if running locally without Docker)*

---

## 🚀 Getting Started & Installation

### 1. Clone the Repository
```bash
git clone https://github.com/your-username/fleet-management-system.git
cd fleet-management-system
```

### 2. Configure Environment Variables
Ensure you have an `.env` file in the root directory configured with your credentials:

```env
NAME=your_service_name
PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USERNAME=
DB_PASSWORD=
DB_NAME=your_db_name
DB_SSL=disable

MQTT_BROKER_HOST=localhost
MQTT_BROKER_PORT=1883
MQTT_BROKER_USER=
MQTT_BROKER_PASSWORD=
MQTT_BROKER_CLIENT_ID=

MESSAGE_QUEUE_HOST=localhost
MESSAGE_QUEUE_PORT=5672
MESSAGE_QUEUE_USER=
MESSAGE_QUEUE_PASSWORD=

RABBITMQ_SUBSCRIBER_QOS=
RABBITMQ_GEOFENCING_EXCHANGE=
RABBITMQ_GEOFENCING_QUEUE=

GEOFENCING_RADIUS_METER=50

SIMULATOR_ENABLE=false
```

### 3. Run the Application via Docker Compose
To build and spin up all services (PostgreSQL, RabbitMQ, Mosquitto, and the Go Fleet App) in detached mode, run:

```bash
docker compose up -d --build
```

### 4. Verify Containers Status
Check if all containers are running smoothly:

```bash
docker compose ps
```

---

## 🛠️ Database Management & Migrations

* **Automated Migration:** When `fleet-app` starts up, it automatically executes database migrations against `fleet_postgres` once the database health check passes successfully.
* **External Access (e.g., DBeaver):**
  * **Host:** `localhost`
  * **Port:** `5432` (or your mapped `DB_PORT`)
  * **Database:** `tj_fleet_managements`
  * **User:** `postgres`

---

## 🛑 Stopping and Resetting Services

* To stop all running containers:
  ```bash
  docker compose stop
  ```
* To completely wipe out containers and clean up persistent volumes *(useful for resetting database state during development)*:
  ```bash
  docker compose down -v
  ```

---

## 📝 License

This project is developed for enterprise fleet tracking and management demonstration purposes.