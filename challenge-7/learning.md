# Learning Materials for Bank Account with Error Handling

## Error Handling in Go

Error handling is a critical aspect of writing robust Go programs. This challenge focuses on implementing a banking system with proper error handling techniques.

### Basic Error Handling

Go uses explicit error handling with return values instead of exceptions:

```go
// Function that may return an error
func divideNumbers(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Error handling with if checks
result, err := divideNumbers(10, 0)
if err != nil {
    fmt.Println("Error:", err)
    return // or handle the error
}
fmt.Println("Result:", result)
```

### Creating Custom Errors

Standard ways to create errors:

```go
// Using errors.New for simple error messages
err := errors.New("insufficient funds")

// Using fmt.Errorf for formatted error messages
amount := 100
err := fmt.Errorf("insufficient funds: need $%d more", amount)
```

### Custom Error Types

Creating custom error types allows for more detailed error handling:

```go
// Define a custom error type
type InsufficientFundsError struct {
    Balance float64
    Amount  float64
}

// Implement the error interface
func (e *InsufficientFundsError) Error() string {
    return fmt.Sprintf("insufficient funds: balance $%.2f, attempted to withdraw $%.2f", 
        e.Balance, e.Amount)
}

// Usage in a function
func (a *Account) Withdraw(amount float64) error {
    if a.Balance < amount {
        return &InsufficientFundsError{
            Balance: a.Balance,
            Amount:  amount,
        }
    }
    a.Balance -= amount
    return nil
}

// Type checking for specific errors
err := account.Withdraw(1000)
if err != nil {
    if insufficientFunds, ok := err.(*InsufficientFundsError); ok {
        fmt.Printf("You need $%.2f more\n", insufficientFunds.Amount - insufficientFunds.Balance)
    } else {
        fmt.Println("Unknown error:", err)
    }
}
```

### Error Wrapping (Go 1.13+)

Go 1.13 introduced error wrapping for better error context:

```go
// Wrapping an error
func processAccount(id string) error {
    account, err := getAccount(id)
    if err != nil {
        return fmt.Errorf("failed to get account %s: %w", id, err)
    }
    
    err = account.Withdraw(100)
    if err != nil {
        return fmt.Errorf("withdrawal failed: %w", err)
    }
    
    return nil
}

// Unwrapping errors
err := processAccount("12345")
if err != nil {
    // Check if it's a specific error type
    var insufficientFunds *InsufficientFundsError
    if errors.As(err, &insufficientFunds) {
        fmt.Printf("You need $%.2f more\n", 
            insufficientFunds.Amount - insufficientFunds.Balance)
    }
    
    // Check for a specific error value
    if errors.Is(err, ErrAccountNotFound) {
        fmt.Println("Account not found, please check the ID")
    }
    
    fmt.Println("Error chain:", err)
}
```

### Sentinel Errors

Predefined errors that can be compared directly:

```go
// Define sentinel errors as package-level variables
var (
    ErrAccountNotFound      = errors.New("account not found")
    ErrInsufficientFunds    = errors.New("insufficient funds")
    ErrInvalidAmount        = errors.New("invalid amount")
)

// Using sentinel errors
func (a *Account) Withdraw(amount float64) error {
    if amount <= 0 {
        return ErrInvalidAmount
    }
    
    if a.Balance < amount {
        return ErrInsufficientFunds
    }
    
    a.Balance -= amount
    return nil
}

// Checking for sentinel errors
err := account.Withdraw(-50)
if err == ErrInvalidAmount {
    fmt.Println("Please enter a positive amount")
} else if err == ErrInsufficientFunds {
    fmt.Println("Not enough money in your account")
}
```

### Error Handling Patterns

#### 1. Return Early Pattern

```go
func processTransaction(tx *Transaction) error {
    // Validate inputs first
    if tx == nil {
        return errors.New("nil transaction")
    }
    
    if tx.Amount <= 0 {
        return errors.New("invalid transaction amount")
    }
    
    // Process after validation passes
    return processValidTransaction(tx)
}
```

#### 2. Error Handler Function

```go
type errHandler func() error

func handleErrors(handlers ...errHandler) error {
    for _, h := range handlers {
        if err := h(); err != nil {
            return err
        }
    }
    return nil
}

// Usage
err := handleErrors(
    func() error { return validateAccount(account) },
    func() error { return checkBalance(account, amount) },
    func() error { return performTransfer(account, amount) },
)
```

### Panic and Recover

While Go prefers explicit error handling, `panic` and `recover` are available for exceptional cases:

```go
func doSomething() (err error) {
    // Set up a recovery
    defer func() {
        if r := recover(); r != nil {
            // Convert panic to error
            err = fmt.Errorf("panic occurred: %v", r)
        }
    }()
    
    // This might panic
    processSomething()
    return nil
}
```

### Thread Safety in Banking Applications

Banking applications need to handle concurrent access:

```go
type Account struct {
    ID      string
    Owner   string
    Balance float64
    mu      sync.Mutex // Protects account operations
}

func (a *Account) Withdraw(amount float64) error {
    a.mu.Lock()
    defer a.mu.Unlock()
    
    if a.Balance < amount {
        return &InsufficientFundsError{
            Balance: a.Balance,
            Amount:  amount,
        }
    }
    
    a.Balance -= amount
    return nil
}

func (a *Account) Deposit(amount float64) error {
    a.mu.Lock()
    defer a.mu.Unlock()
    
    if amount <= 0 {
        return errors.New("deposit amount must be positive")
    }
    
    a.Balance += amount
    return nil
}
```

### Testing Error Scenarios

Testing error handling is critical:

```go
func TestAccountWithdraw(t *testing.T) {
    // Test insufficient funds
    account := &Account{Balance: 100}
    err := account.Withdraw(150)
    
    // Check error type
    insufficientFunds, ok := err.(*InsufficientFundsError)
    if !ok {
        t.Fatalf("expected InsufficientFundsError, got %T", err)
    }
    
    // Check error details
    if insufficientFunds.Balance != 100 || insufficientFunds.Amount != 150 {
        t.Errorf("wrong error details: %v", insufficientFunds)
    }
    
    // Successful withdrawal
    err = account.Withdraw(50)
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }
    
    if account.Balance != 50 {
        t.Errorf("expected balance 50, got %.2f", account.Balance)
    }
}
```

### Error Logging and Reporting

Proper error logging is essential:

```go
func processTransaction(tx *Transaction) error {
    account, err := getAccount(tx.AccountID)
    if err != nil {
        log.Printf("Error retrieving account %s: %v", tx.AccountID, err)
        return fmt.Errorf("account retrieval failed: %w", err)
    }
    
    err = account.Withdraw(tx.Amount)
    if err != nil {
        // Log with context
        log.Printf("Withdrawal of $%.2f failed for account %s: %v", 
            tx.Amount, tx.AccountID, err)
        return err
    }
    
    // Log success
    log.Printf("Successfully processed withdrawal of $%.2f from account %s",
        tx.Amount, tx.AccountID)
    return nil
}
```

## Further Reading

- [Error Handling in Go](https://blog.golang.org/error-handling-and-go)
- [Working with Errors in Go 1.13+](https://blog.golang.org/go1.13-errors)
- [Effective Error Handling in Go](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)
- [Sync Package Documentation](https://pkg.go.dev/sync) 