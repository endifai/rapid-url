# RapidURL - URL Shortener Service

This is a simple URL shortener service built with Golang and Echo framework, utilizing Redis for data storage. The main functionality of this service is to shorten long URLs into more manageable and shareable short URLs.

## Features

- **Shorten URLs:** Convert long URLs into short, unique, and easy-to-share links.
- **Redirect:** Users can access the short URL, and the service will redirect them to the original long URL.

## Tech Stack

- **Golang:** The backend of the service is written in Golang, providing a fast and efficient execution environment.
- **Echo Framework:** The web framework used for building the RESTful API endpoints.
- **Redis:** Redis is used as the data store to persistently store the mapping between short and long URLs.
- **Docker:** The application is containerized with Docker, simplifying deployment and ensuring consistency across different environments. The Golang app containerized with distroless image.  

## Getting Started

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/) installed on your machine.

### Running the Application

1. **Clone the repository:**

    ```bash
    git clone https://github.com/endifai/rapid-url.git
    ```

2. **Navigate to the project directory:**

    ```bash
    cd rapid-url
    ```

3. **Build and run the Docker containers:**

    ```bash
    docker-compose up -d
    ```

4. **The service will be accessible at [http://localhost:3000](http://localhost:3000).**

## API Endpoints

### Shorten a URL

- **Endpoint:** `/api/shorten-url`
- **Method:** POST
- **Request Body:**

    ```json
    {
      "link": "https://www.example.com/long/url/to/be/shortened"
    }
    ```

### Redirect to Original URL

- **Endpoint:** `/s/:hash`
- **Method:** GET
- **Response:**
  - Redirects to the original long URL.

## Configuration

The application can be configured through environment variables. The following variables are available:

- **`REDIS_ADDRESS`**: Address of the Redis server.
- **`REDIS_PASSWORD`**: Password of the Redis server.
- **`PORT`**: Port on which the service will be available.
