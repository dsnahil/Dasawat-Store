# Snahil's Online Store - Product API

This project is an implementation of a Product API for an e-commerce system. The API allows for creating and retrieving products.

## Project Structure

- `main.go`: The main application server code.
- `api.yaml`: The OpenAPI specification for the API.
- `Dockerfile`: Used to containerize the application.
- `go.mod`: Go module file for dependencies.

## Overview

The server is written in Go and provides RESTful endpoints to manage products. Product data is stored in memory.

### API Endpoints

The implemented API endpoints are based on the `Product` section of the provided `api.yaml`.

- `POST /product`: Creates a new product.
- `GET /product`: Retrieves a list of all products.
- `GET /product/{productID}`: Retrieves a specific product by its ID.

## How to Run

You can run the application either locally using Go or with Docker.


### Running with Go

1.  **Run the server:**
    ```bash     
    go run main.go
    ```
    The server will start on port 8080.


<img width="1037" height="242" alt="image" src="https://github.com/user-attachments/assets/effab141-00bf-4e5c-a8ca-f75cb38bee74" />

<img width="827" height="268" alt="image" src="https://github.com/user-attachments/assets/696dd625-1807-4944-ab29-451669899cec" />

<img width="1034" height="206" alt="image" src="https://github.com/user-attachments/assets/92cbaed5-afc6-4d9e-b2d9-3212ffdba5f5" />

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

<img width="309" height="36" alt="image" src="https://github.com/user-attachments/assets/c54fe4fa-ac7b-4989-9630-a5fc0c3d6041" />

<img width="1026" height="103" alt="image" src="https://github.com/user-attachments/assets/afaa73e7-a31a-4367-84ec-e3651621f715" />

<img width="738" height="58" alt="image" src="https://github.com/user-attachments/assets/689f2ab2-d9fd-4ce4-b993-e74d551298f3" />



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

The deployment process is automated using Terraform.



<img width="1908" height="278" alt="image" src="https://github.com/user-attachments/assets/386dbcbf-94df-4965-88be-ce4c086c780b" />


<img width="1695" height="181" alt="image" src="https://github.com/user-attachments/assets/841cc392-d0e9-472f-8347-a3b798f54177" />


<img width="1898" height="554" alt="image" src="https://github.com/user-attachments/assets/d9b6f862-787b-437c-a929-1b69c0655dd3" />


<img width="984" height="365" alt="image" src="https://github.com/user-attachments/assets/cdfe3ae3-c139-4fd7-9ef7-33b13e1b0c18" />



The load test was performed with 100 concurrent users against the live application deployed on AWS ECS. 



<img width="1382" height="910" alt="image" src="https://github.com/user-attachments/assets/06dcda11-1e7e-4eb0-bbfa-3829353e72ee" />



---

## Performance and Bottlenecks

The single-container deployment demonstrated a **stable throughput** of approximately **33-34 requests per second**. This indicates that the **server's processing limit** is the primary **bottleneck** in this configuration, rather than client-side limitations.

---

## Latency Analysis

With exceptionally **low response times**, the server proved to be highly efficient. In this scenario, the performance difference between Locust's `HttpUser` and `FastHttpUser` would be negligible, as the bottleneck is the server's capacity, not the client's request-generation speed.

---

## Data Structure Impact

The use of an **in-memory map** is directly responsible for the application's **very fast read operations** (`GET` requests). However, this data structure is limited:

* **Lack of Durability:** The data is **not durable**, meaning all information is permanently lost if the container restarts.
* **Horizontal Scaling Issues:** The system **cannot be scaled horizontally**. Running multiple containers would create separate, inconsistent in-memory databases, making a shared, external database a necessity for growth.


