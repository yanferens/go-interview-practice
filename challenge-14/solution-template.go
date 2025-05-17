package main

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Note: In a real implementation, these would be in separate .proto files
// and compiled with protoc. For this challenge, we'll define the interfaces directly.

// UserService defines operations for managing users
type UserService interface {
	GetUser(ctx context.Context, userID int64) (*User, error)
	ValidateUser(ctx context.Context, userID int64) (bool, error)
}

// ProductService defines operations for managing products
type ProductService interface {
	GetProduct(ctx context.Context, productID int64) (*Product, error)
	CheckInventory(ctx context.Context, productID int64, quantity int32) (bool, error)
}

// User represents a user in the system
type User struct {
	ID       int64
	Username string
	Email    string
	Active   bool
}

// Product represents a product in the catalog
type Product struct {
	ID        int64
	Name      string
	Price     float64
	Inventory int32
}

// Order represents an order in the system
type Order struct {
	ID        int64
	UserID    int64
	ProductID int64
	Quantity  int32
	Total     float64
}

// UserServiceServer implements the UserService
type UserServiceServer struct {
	// A map to store users (in a real application, this would be a database)
	users map[int64]*User
}

// NewUserServiceServer creates a new UserServiceServer
func NewUserServiceServer() *UserServiceServer {
	// Initialize with some test users
	users := map[int64]*User{
		1: {ID: 1, Username: "alice", Email: "alice@example.com", Active: true},
		2: {ID: 2, Username: "bob", Email: "bob@example.com", Active: true},
		3: {ID: 3, Username: "charlie", Email: "charlie@example.com", Active: false},
	}
	return &UserServiceServer{users: users}
}

// GetUser retrieves a user by ID
func (s *UserServiceServer) GetUser(ctx context.Context, userID int64) (*User, error) {
	// TODO: Implement this method
	// 1. Check if the user exists
	// 2. Return the user or an appropriate gRPC error
	return nil, errors.New("not implemented")
}

// ValidateUser checks if a user exists and is active
func (s *UserServiceServer) ValidateUser(ctx context.Context, userID int64) (bool, error) {
	// TODO: Implement this method
	// 1. Check if the user exists and is active
	// 2. Return true if valid, false otherwise, with appropriate errors
	return false, errors.New("not implemented")
}

// ProductServiceServer implements the ProductService
type ProductServiceServer struct {
	// A map to store products (in a real application, this would be a database)
	products map[int64]*Product
}

// NewProductServiceServer creates a new ProductServiceServer
func NewProductServiceServer() *ProductServiceServer {
	// Initialize with some test products
	products := map[int64]*Product{
		1: {ID: 1, Name: "Laptop", Price: 999.99, Inventory: 10},
		2: {ID: 2, Name: "Phone", Price: 499.99, Inventory: 20},
		3: {ID: 3, Name: "Headphones", Price: 99.99, Inventory: 0},
	}
	return &ProductServiceServer{products: products}
}

// GetProduct retrieves a product by ID
func (s *ProductServiceServer) GetProduct(ctx context.Context, productID int64) (*Product, error) {
	// TODO: Implement this method
	// 1. Check if the product exists
	// 2. Return the product or an appropriate gRPC error
	return nil, errors.New("not implemented")
}

// CheckInventory checks if a product is available in the requested quantity
func (s *ProductServiceServer) CheckInventory(ctx context.Context, productID int64, quantity int32) (bool, error) {
	// TODO: Implement this method
	// 1. Check if the product exists
	// 2. Check if the inventory is sufficient
	// 3. Return true if available, false otherwise, with appropriate errors
	return false, errors.New("not implemented")
}

// OrderService handles order creation
type OrderService struct {
	userClient    UserService
	productClient ProductService
	orders        map[int64]*Order
	nextOrderID   int64
}

// NewOrderService creates a new OrderService
func NewOrderService(userClient UserService, productClient ProductService) *OrderService {
	return &OrderService{
		userClient:    userClient,
		productClient: productClient,
		orders:        make(map[int64]*Order),
		nextOrderID:   1,
	}
}

// CreateOrder creates a new order
func (s *OrderService) CreateOrder(ctx context.Context, userID, productID int64, quantity int32) (*Order, error) {
	// TODO: Implement this method
	// 1. Validate the user with UserService
	// 2. Get the product details and check inventory with ProductService
	// 3. Create the order if everything is valid
	// 4. Return the order or an appropriate error
	return nil, errors.New("not implemented")
}

// GetOrder retrieves an order by ID
func (s *OrderService) GetOrder(orderID int64) (*Order, error) {
	order, exists := s.orders[orderID]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "order not found")
	}
	return order, nil
}

// LoggingInterceptor is a server interceptor for logging
func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// TODO: Implement a logging interceptor that logs the method being called
	// and returns the result from the handler
	return nil, errors.New("not implemented")
}

// AuthInterceptor is a client interceptor for authentication
func AuthInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// TODO: Implement an auth interceptor that adds authentication information
	// to the context and calls the invoker
	return errors.New("not implemented")
}

// StartUserService starts the user service on the given port
func StartUserService(port string) (*grpc.Server, error) {
	// TODO: Implement this function
	// 1. Create a new gRPC server with the logging interceptor
	// 2. Register the UserServiceServer
	// 3. Start listening on the given port
	// 4. Return the server
	return nil, errors.New("not implemented")
}

// StartProductService starts the product service on the given port
func StartProductService(port string) (*grpc.Server, error) {
	// TODO: Implement this function
	// 1. Create a new gRPC server with the logging interceptor
	// 2. Register the ProductServiceServer
	// 3. Start listening on the given port
	// 4. Return the server
	return nil, errors.New("not implemented")
}

// Connect to both services and return an OrderService
func ConnectToServices(userServiceAddr, productServiceAddr string) (*OrderService, error) {
	// TODO: Implement this function
	// 1. Connect to both the user and product services with gRPC
	// 2. Create clients with the auth interceptor
	// 3. Return a new OrderService with these clients
	return nil, errors.New("not implemented")
}

func main() {
	// Example usage:
	// 1. Start both services on different ports
	// 2. Connect to these services
	// 3. Create an order and print the result
}
