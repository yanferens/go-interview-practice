[View the Scoreboard](SCOREBOARD.md)

# Challenge 14: Microservices with gRPC

In this challenge, you will implement a simple microservice architecture using gRPC for service-to-service communication. You'll create a User service and a Product service that communicate with each other to provide an order management system.

## Requirements

1. Implement two gRPC services:
   - `UserService`: manages user information  
   - `ProductService`: manages product catalog  
   
2. Define Protocol Buffers for both services with the following methods:
   - UserService:
     - `GetUser(id)`: returns user details
     - `ValidateUser(id)`: validates if user exists and can place orders
   - ProductService:
     - `GetProduct(id)`: returns product details
     - `CheckInventory(id, quantity)`: checks if product is available in requested quantity
     
3. Create a client for these services that demonstrates:
   - Authentication with the UserService
   - Product validation with the ProductService
   - Creating an order that calls both services
   
4. Implement proper error handling with gRPC status codes
5. Use interceptors for logging and basic authentication
6. The included test file has scenarios covering both successful and failed scenarios 