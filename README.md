
# User Service Project

This project is a **User Service** built using **Golang**, **PostgreSQL**, and **Redis**. It provides a service for searching users based on various filters, including full-text search for fields like `f_name`, `city`, and `phone`.

## Features

- **Full-Text Search**: Allows searching users by name, city, or any other relevant text fields using PostgreSQL's full-text search with `tsvector` and `GIN` index.
- **Search Filters**: Supports filtering by city, phone, marital status, and a custom query string.
- **Pagination**: Results can be paginated with `LIMIT` and `OFFSET` for optimized large data queries.
- **Caching**: Uses **Redis** for caching query results to improve response time for frequently searched queries.
- **Efficient Data Storage**: Implements a `search_vector` column in PostgreSQL to store precomputed text vectors for full-text search.

## Installation

### Prerequisites

- Go (1.18+)
- PostgreSQL (with full-text search support)
- Redis (for caching)
- Docker (Optional for containerization)

### Setting up the Project

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/MuhammedAshifVnr/user_service.git
   cd user_service
   ```

2. **Install Dependencies**:
   Run the following command to install the Go dependencies:
   ```bash
   go mod tidy
   ```

3. **Set up PostgreSQL Database**:
   Make sure PostgreSQL is installed and running. Create the required database and table. You can use the following SQL commands:
   ```sql
   CREATE DATABASE user_service;

   CREATE TABLE users (
     id SERIAL PRIMARY KEY,
     f_name TEXT,
     city TEXT,
     phone TEXT,
     height FLOAT,
     married BOOLEAN,
     search_vector TSVECTOR
   );

   CREATE INDEX search_vector_idx ON users USING GIN (search_vector);
   ```

4. **Set up Redis**:
   Install and run Redis. You can use a Docker container for Redis:
   ```bash
   docker run --name redis -p 6379:6379 -d redis
   ```

5. **Configure Environment Variables**:
   Set up environment variables for your database and Redis connection.
   - `DB_HOST`: The PostgreSQL host (e.g., `localhost`).
   - `DB_PORT`: The PostgreSQL port (e.g., `5432`).
   - `DB_USER`: The PostgreSQL username (e.g., `user`).
   - `DB_PASSWORD`: The PostgreSQL password (e.g., `password`).
   - `DB_NAME`: The PostgreSQL database name (e.g., `user_service`).
   - `REDIS_HOST`: The Redis host (e.g., `localhost`).
   - `REDIS_PORT`: The Redis port (e.g., `6379`).

6. **Run the Application**:
   After configuring the environment variables, run the application:
   ```bash
   go run main.go
   ```

## API Endpoints

### Search Users

**Endpoint**: `/search`

**Method**: `POST`

**Request Body**:
```json
{
  "city": "New York",
  "phone": "1234567890",
  "query": "Steve",
  "married": true,
  "limit": 10,
  "offset": 0
}
```

**Response**:
```json
{
  "users": [
    {
      "id": 1,
      "f_name": "Steve",
      "city": "New York",
      "phone": "1234567890",
      "height": 5.9,
      "married": true
    }
  ]
}
```

### Caching

- The search results are cached using **Redis** for 5 minutes to improve response times for repeated searches with the same parameters.

## Database and Full-Text Search

- **`search_vector`**: A `tsvector` column is used to store indexed text data, allowing for fast full-text search.
- **`GIN Index`**: A GIN index is created on the `search_vector` column to optimize search queries.

## Testing

1. **Unit Tests**:
   - Write unit tests for repository and service layers to ensure correctness.
   - Example:
   ```bash
   go test ./...
   ```

2. **Integration Tests**:
   - Test the API endpoints using tools like **Postman** or **cURL**.

## Docker (Optional)

You can also run the project in Docker for easy containerization.

### Dockerfile

```dockerfile
FROM golang:1.18

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o user_service .

EXPOSE 8080

CMD ["./user_service"]
```

### Docker Compose (Optional)

You can use Docker Compose to set up PostgreSQL and Redis along with the Go application.

```yaml
version: '3'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      DB_HOST: db
      DB_PORT: 5432
      DB_USER: user
      DB_PASSWORD: password
      DB_NAME: user_service
      REDIS_HOST: redis
      REDIS_PORT: 6379
    depends_on:
      - db
      - redis

  db:
    image: postgres:latest
    environment:
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
      POSTGRES_DB: user_service
    ports:
      - "5432:5432"

  redis:
    image: redis:latest
    ports:
      - "6379:6379"
```

Run the Docker Compose setup:
```bash
docker-compose up
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

