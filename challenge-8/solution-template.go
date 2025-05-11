// Package challenge8 contains the solution for Challenge 8: Chat Server with Channels.
package challenge8

import (
	"errors"
	"sync"
	// Add any other necessary imports
)

// Message represents a chat message
type Message struct {
	Sender    string
	Content   string
	IsPrivate bool
	Recipient string // Only used for private messages
}

// Client represents a connected chat client
type Client struct {
	// TODO: Implement this struct
	// Hint: You'll need username, channels for incoming/outgoing messages, etc.
}

// Send sends a message to the client
func (c *Client) Send(message string) {
	// TODO: Implement this method
}

// Receive returns the next message for the client (blocking)
func (c *Client) Receive() string {
	// TODO: Implement this method
	return ""
}

// ChatServer manages client connections and message routing
type ChatServer struct {
	// TODO: Implement this struct
	// Hint: You'll need a map of clients, mutex for thread-safety, etc.
}

// NewChatServer creates a new chat server instance
func NewChatServer() *ChatServer {
	// TODO: Implement this function
	return nil
}

// Connect adds a new client to the chat server
func (s *ChatServer) Connect(username string) (*Client, error) {
	// TODO: Implement this method
	// 1. Check if username is already taken
	// 2. Create a new client with appropriate channels
	// 3. Add the client to the server's client map
	// 4. Return the client
	return nil, nil
}

// Disconnect removes a client from the chat server
func (s *ChatServer) Disconnect(client *Client) {
	// TODO: Implement this method
	// 1. Remove the client from the server's client map
	// 2. Close the client's channels
	// 3. Notify other clients about the disconnection
}

// Broadcast sends a message to all connected clients
func (s *ChatServer) Broadcast(sender *Client, message string) {
	// TODO: Implement this method
	// 1. Create a message with sender information
	// 2. Send the message to all connected clients
}

// PrivateMessage sends a message to a specific client
func (s *ChatServer) PrivateMessage(sender *Client, recipient string, message string) error {
	// TODO: Implement this method
	// 1. Find the recipient client
	// 2. If recipient exists, send the message
	// 3. If recipient doesn't exist, return an error
	return nil
}

// Common errors that can be returned by the Chat Server
var (
	ErrUsernameAlreadyTaken = errors.New("username already taken")
	ErrRecipientNotFound    = errors.New("recipient not found")
	ErrClientDisconnected   = errors.New("client disconnected")
	// Add more error types as needed
) 