# OrderServiceApp

`OrderServiceApp` is a scalable and efficient `Golang and gRPC` backend application for ecommerce order management. It provides a robust solution for retrieving product information, placing orders, and tracking order status. This repository contains the source code and resources required to run the application.

## Features

`Product Service:` Retrieve product information such as availability, price, and category.
`Order Service:` Get order details including order value, dispatch date, order status, and product quantity.
`Catalogue Update:` Automatically update the product catalogue when an order is placed.
`Order Discount:` Apply a 10% discount on the order value if it contains 3 different premium products.
`Order Status Update:` Update the order status for a specific order.
`Dispatch Date:` Populate the dispatch date only when the order status is set to 'Dispatched'.


## Technologies Used

1. `Golang:` The application is built using the Go programming language, which provides performance, concurrency, and scalability.
2. `gRPC:` The communication protocol used for efficient and reliable inter-service communication.
3. `NoSQL in-memory map:` For simplicity, an in-memory map is used as the data persistence layer.
4. `JSON:` All request payloads and responses follow the JSON format, ensuring interoperability and ease of integration.
5. `Modular Architecture:` The codebase follows a modular structure, making it easy to understand, maintain, and extend.

## Prerequisites

Before running the OrderServiceApp, ensure you have the following:
Go programming language installed 
gRPC framework installed
Any required dependencies specified in the go.mod file

## Installation
Clone this repository to your local machine.
Navigate to the project directory: cd OrderServiceApp.
Install the dependencies: go mod download.

## Usage.
Start the OrderServiceApp server:
`gRPC`: localhost:50051
`REST`: localhost:5051
Use your preferred gRPC client to interact with the available endpoints.

## MakeFile
`To run the gPRC server`

```
make app 
```

`To run the REST server`

```
make rest
```

## API Endpoints
The OrderServiceApp provides the following gRPC  and REST endpoints 

## Contributing
Contributions to OrderServiceApp are welcome! If you find any bugs or have suggestions for improvements, please open an issue or submit a pull request. Follow the contribution guidelines mentioned in the CONTRIBUTING.md file.

## License
OrderServiceApp is licensed under the Apache License.

## Contact
For any inquiries or feedback, please contact us at sumitthakur769@gmail.com.
