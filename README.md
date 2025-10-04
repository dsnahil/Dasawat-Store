# Simple Online Store - Product API

This project is an implementation of a simple Product API for an e-commerce system, as part of a university assignment. The API allows for creating and retrieving products.

## Project Structure

- `main.go`: The main application server code.
- `api.yaml`: The OpenAPI specification for the API.
- `Dockerfile`: Used to containerize the application.
- `go.mod`: Go module file for dependencies.

## Overview

The server is written in Go and provides RESTful endpoints to manage products. Product data is stored in-memory.

### API Endpoints

The implemented API endpoints are based on the `Product` section of the provided `api.yaml`.

- `POST /product`: Creates a new product.
- `GET /product`: Retrieves a list of all products.
- `GET /product/{productID}`: Retrieves a specific product by its ID.

## How to Run

You can run the application either locally using Go or with Docker.

<img width="1037" height="242" alt="image" src="https://github.com/user-attachments/assets/effab141-00bf-4e5c-a8ca-f75cb38bee74" />

<img width="827" height="268" alt="image" src="https://github.com/user-attachments/assets/696dd625-1807-4944-ab29-451669899cec" />

<img width="1034" height="206" alt="image" src="https://github.com/user-attachments/assets/92cbaed5-afc6-4d9e-b2d9-3212ffdba5f5" />




### Running with Go

1.  **Run the server:**
    ```bash     
    go run main.go
    ```
    The server will start on port 8080.

### Running with Docker

1.  **Build the Docker image:**
    ```bash
    docker build -t product-api .
    ```
 
2.  **Run the Docker container:**
    ```bash
    docker run -p 8080:8080 product-api
    ```
    The server will be accessible at `http://localhost:8080`.

## API Usage Examples

Here are some examples of how to interact with the API using `curl`.

### Create a Product

**Request:**

```bash
curl -X POST http://localhost:8080/product -H "Content-Type: application/json" -d '{
  "name": "Laptop",
  "description": "A powerful laptop",
  "price": 1200.50
}'
```

**Successful Response (201 Created):**

The server will respond with the created product, including its new ID.

```json
{
  "id": "some-unique-id",
  "name": "Laptop",
  "description": "A powerful laptop",
  "price": 1200.50
}
```

**Invalid Input Response (400 Bad Request):**

If the request body is invalid.

### Get All Products

**Request:**

```bash
curl -X GET http://localhost:8080/product
```

**Successful Response (200 OK):**

```json
[
  {
    "id": "some-unique-id",
    "name": "Laptop",
    "description": "A powerful laptop",
    "price": 1200.50
  }
]
```

### Get a Product by ID

**Request:**

```bash
curl -X GET http://localhost:8080/product/{productID}
```
Replace `{productID}` with an actual product ID.

**Successful Response (200 OK):**

```json
{
  "id": "some-unique-id",
  "name": "Laptop",
  "description": "A powerful laptop",
  "price": 1200.50
}
```

**Product Not Found Response (404 Not Found):**

If a product with the given ID does not exist.

## Deployment

The infrastructure for deploying this application to AWS ECS can be found in the separate infrastructure repository. The deployment process is automated using Terraform.

For instructions on how to deploy, please refer to the `README.md` in the infrastructure repository.

