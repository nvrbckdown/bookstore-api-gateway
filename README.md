## API Endpoints

The API Gateway exposes the following endpoints:

### Health Check
- `GET /health` - Check if the API Gateway is running

### Books
- `GET /api/v1/books` - Get all books
- `GET /api/v1/books/:id` - Get a specific book
- `POST /api/v1/books` - Create a new book
- `PUT /api/v1/books/:id` - Update a book
- `DELETE /api/v1/books/:id` - Delete a book

### Authors
- `GET /api/v1/authors` - Get all authors
- `GET /api/v1/authors/:id` - Get a specific author
- `POST /api/v1/authors` - Create a new author
- `PUT /api/v1/authors/:id` - Update an author
- `DELETE /api/v1/authors/:id` - Delete an author
- `GET /api/v1/authors/:id/books` - Get all books by a specific author

### Orders
- `GET /api/v1/orders` - Get all orders
- `GET /api/v1/orders/:id` - Get a specific order
- `POST /api/v1/orders` - Create a new order
- `PUT /api/v1/orders/:id` - Update an order status
- `DELETE /api/v1/orders/:id` - Delete an order
- `GET /api/v1/orders/customer/:email` - Get all orders by a customer email

## Running with Docker

```bash
# Build the container
docker build -t bookstore-api-gateway .

# Run the container
docker run -p 8080:8080 bookstore-api-gateway

# Or use Docker Compose to run the entire system
docker-compose up
```

## Authentication

Authentication is currently commented out but can be easily enabled by uncommenting the JWT authentication code in the middleware and routes. This provides a foundation for adding proper authentication in the future.

## Error Handling

The API Gateway handles errors from both upstream services and internal errors. It forwards the appropriate HTTP status codes and error messages from the services to the client.