# Receipt Processor Service

The Receipt Processor Service is a web service designed to process receipts, calculate points based on certain rules, and retrieve those points via API calls. This service is built using Go and can be run directly or within a Docker container.

## Prerequisites

-   Go (if running without Docker)
-   Docker (for containerized execution)

## Running the Service Without Docker

1.  **Clone the Repository**
    
    bashCopy code
    
    `git clone https://github.com/tanand4/Receipt-Processor-Service`
    
    `cd Receipt-Processor-Service/receipt-processor` 
    
2.  **Build the Application**
    
    Ensure Go is installed on your system, then run:
    
    bashCopy code
    
    `go build -o receipt-processor ./cmd/server` 
    
3.  **Run the Application**
    
    Execute the binary:
    
    bashCopy code
    
    `./receipt-processor` 
    
    The service will start, listening on the default port (e.g., `:8080`).
    

## Running the Service With Docker

1.  **Clone the Repository**
    
    bashCopy code
    
    `git clone https://github.com/tanand4/Receipt-Processor-Service`
    
     
     `cd Receipt-Processor-Service/receipt-processor` 
    
2.  **Build the Docker Image**
    
    bashCopy code
    
    `docker build -t receipt-processor .` 
    
3.  **Run the Docker Container**
    
    bashCopy code
    
    `docker run -p 8080:8080 receipt-processor` 
    
    This command maps port 8080 on the container to port 8080 on your host, allowing you to access the service at `http://localhost:8080`.
    

## API Endpoints

-   **Process Receipts**
    
    -   **Endpoint:** `/receipts/process`
    -   **Method:** POST
    -   **Payload:** JSON representing the receipt.
    -   **Description:** Processes the given receipt, calculates points, and returns a unique ID for the receipt.
-   **Get Points**
    
    -   **Endpoint:** `/receipts/{id}/points`
    -   **Method:** GET
    -   **Description:** Retrieves the number of points awarded for the receipt identified by the given ID.

## Example Requests

### Process Receipt

bashCopy code

`curl -X POST http://localhost:8080/receipts/process \
-H "Content-Type: application/json" \
-d '{ "retailer": "Store", "items": [{ "description": "Item 1", "price": "10.00" }], "total": "10.00" }'` 

### Get Points

bashCopy code

`curl http://localhost:8080/receipts/{id}/points` 

Replace `{id}` with the actual ID returned from the process receipt call.
