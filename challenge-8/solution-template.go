// Package challenge8 contains the solution for Challenge 8: Chat Server with Channels.
package challenge8

import (
	"errors"
	// Add any other necessary imports
)

// Client represents a connected chat client
type Client struct {
	// TODO: Implement this struct
	// Hint: username, message channel, mutex, disconnected flag
}

// Send sends a message to the client
func (c *Client) Send(message string) {
	// TODO: Implement this method
	// Hint: thread-safe, non-blocking send
}

// Receive returns the next message for the client (blocking)
func (c *Client) Receive() string {
	// TODO: Implement this method
	// Hint: read from channel, handle closed channel
	return ""
}

// ChatServer manages client connections and message routing
type ChatServer struct {
	// TODO: Implement this struct
	// Hint: clients map, mutex
}

// NewChatServer creates a new chat server instance
func NewChatServer() *ChatServer {
	// TODO: Implement this function
	return nil
}

// Connect adds a new client to the chat server
func (s *ChatServer) Connect(username string) (*Client, error) {
	// TODO: Implement this method
	// Hint: check username, create client, add to map
	return nil, nil
}

// Disconnect removes a client from the chat server
func (s *ChatServer) Disconnect(client *Client) {
	// TODO: Implement this method
	// Hint: remove from map, close channels
}

// Broadcast sends a message to all connected clients
func (s *ChatServer) Broadcast(sender *Client, message string) {
	// TODO: Implement this method
	// Hint: format message, send to all clients
}

// PrivateMessage sends a message to a specific client
func (s *ChatServer) PrivateMessage(sender *Client, recipient string, message string) error {
	// TODO: Implement this method
	// Hint: find recipient, check errors, send message
	return nil
}

// Common errors that can be returned by the Chat Server
var (
	ErrUsernameAlreadyTaken = errors.New("username already taken")
	ErrRecipientNotFound    = errors.New("recipient not found")
	ErrClientDisconnected   = errors.New("client disconnected")
	// Add more error types as needed
)
