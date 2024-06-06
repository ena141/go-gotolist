# go TODO List

A simple TODO list application built with Go and PostgreSQL. The application is containerized using Docker and can be deployed using Docker Compose. 

## Features

- Add, complete, and delete tasks
- Persistent storage using PostgreSQL
- Containerized using Docker for easy deployment

## Prerequisites

- Go
- PostgreSQL
- Docker

## Getting Started

### 1. Clone the Repository

```sh
git clone https://github.com/ena141/go-gotolist.git
cd go-gotolist
```

### 2. Build and Run with Docker Compose

Ensure you have Docker and Docker Compose installed. Then, you can build and run the application using the following commands:

```
docker-compose up --build
```

This command will:

1. Build the Docker image for the Go application.
2. Pull the PostgreSQL image from Docker Hub.
3. Start both containers and set up the necessary network.

### 3. Access the Application

Once the containers are up and running, you can access the TODO list application at `http://localhost:8080`.

## Configuration

The application uses the following environment variables to configure the database connection:

- `DB_HOST`: Hostname of the PostgreSQL database (default: `db`).
- `DB_PORT`: Port of the PostgreSQL database (default: `5432`).
- `DB_USER`: Username for the PostgreSQL database (default: `postgres`).
- `DB_PASSWORD`: Password for the PostgreSQL database (default: `123456`).
- `DB_NAME`: Name of the PostgreSQL database (default: `postgres`).

These environment variables are set in the `docker-compose.yml` file.
