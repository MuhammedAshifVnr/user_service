
# User Service - gRPC API

This project is a User Service that exposes a gRPC API for managing user details, including search functionality. The service is built using Golang and utilizes PostgreSQL for data storage. The service also includes caching with Redis and supports full-text search using PostgreSQL's built-in search capabilities.

## Prerequisites

Before you can build and run the application, ensure you have the following installed:

- Docker
- Go (Golang) 1.22.2 or higher
- Docker Compose (optional, if you need to manage multiple services)

## Setup and Installation

### 1. Clone the Repository

```bash
git clone <repository-url>
cd <repository-directory>
```

### 2. Build the Docker Image

To build the Docker image, run the following command:

```bash
docker build -t user-svc .
```

This will create a Docker image named `user-svc` using the provided `Dockerfile`. The build process includes:

- Downloading Go dependencies.
- Compiling the application.
- Creating a lightweight runtime image with the compiled binary and environment variables.

### 3. Running the Docker Container

Once the image is built, you can run the container:

```bash
docker run -d -p 5001:5001 --name user-svc user-svc
```

This will start the container and expose the application on port `5001`.

### 4. Environment Configuration

The service uses an `.env` file to configure environment variables. This file is copied into the container during the build process. You can modify the `.env` file to set database credentials, Redis configurations, and other service-specific variables.

### 5. Accessing the gRPC Service

Once the container is running, you can interact with the service through the gRPC endpoint exposed on port `5001`.

### 6. Full-Text Search and Caching

This service uses PostgreSQL's full-text search capabilities to perform efficient searches on user data. Additionally, caching is implemented using Redis to reduce database load and improve search performance.

## API Documentation

API documentation can be generated using tools like Swagger or Postman. The service supports CRUD operations for user details and a full-text search endpoint to search users based on criteria such as city, phone, marital status, and more.

## Contributing

Feel free to fork the repository, create a new branch, and submit pull requests for any improvements or bug fixes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
