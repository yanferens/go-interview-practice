[View the Scoreboard](SCOREBOARD.md)

# Challenge 12: File Processing Pipeline with Advanced Error Handling

## Problem Statement

Implement a file processing pipeline that reads, transforms, and writes data with comprehensive error handling that demonstrates Go's idiomatic approach to errors.

## Requirements

1. Implement a modular file processing pipeline that:
   - Reads data from various sources (files, network, in-memory)
   - Validates and transforms data through multiple processing stages
   - Writes results to a destination
   - Implements comprehensive error handling at each stage

2. You must implement the following error handling techniques:
   - Custom error types with embedded standard errors
   - Error wrapping to preserve context across pipeline stages
   - Sentinel errors for specific conditions
   - Type-based error handling with errors.Is and errors.As
   - Proper error propagation in concurrent contexts

3. The pipeline should have these components:
   - **Reader**: Reads data from a source (file, URL, memory)
   - **Validator**: Validates the data according to rules
   - **Transformer**: Transforms valid data into a different format
   - **Writer**: Writes processed data to a destination
   - **Pipeline**: Orchestrates the flow between components

## Function Signatures

```go
// Core interfaces
type Reader interface {
    Read(ctx context.Context) ([]byte, error)
}

type Validator interface {
    Validate(data []byte) error
}

type Transformer interface {
    Transform(data []byte) ([]byte, error)
}

type Writer interface {
    Write(ctx context.Context, data []byte) error
}

// Custom error types
type ValidationError struct {
    Field   string
    Message string
    Err     error
}

type TransformError struct {
    Stage string
    Err   error
}

type PipelineError struct {
    Stage string
    Err   error
}

// Error methods
func (e *ValidationError) Error() string
func (e *ValidationError) Unwrap() error

func (e *TransformError) Error() string
func (e *TransformError) Unwrap() error

func (e *PipelineError) Error() string
func (e *PipelineError) Unwrap() error

// Sentinel errors
var (
    ErrInvalidFormat    = errors.New("invalid data format")
    ErrMissingField     = errors.New("required field missing")
    ErrProcessingFailed = errors.New("processing failed")
    ErrDestinationFull  = errors.New("destination is full")
)

// Pipeline implementation
type Pipeline struct {
    Reader      Reader
    Validators  []Validator
    Transformers []Transformer
    Writer      Writer
}

func NewPipeline(r Reader, v []Validator, t []Transformer, w Writer) *Pipeline

func (p *Pipeline) Process(ctx context.Context) error

// Helper for handling errors in concurrent operations
func (p *Pipeline) handleErrors(ctx context.Context, errs <-chan error) error
```

## Constraints

- All errors must provide meaningful context about where and why they occurred
- Use error wrapping with %w where appropriate to maintain error chains
- Implement proper error comparison using errors.Is and errors.As
- Demonstrate both sentinel error checking and type-based error handling
- Implement graceful error handling in concurrent operations
- When a pipeline fails, it should clean up resources properly

## Sample Usage

```go
// Create pipeline components
fileReader := NewFileReader("input.json")
validators := []Validator{
    NewJSONValidator(),
    NewSchemaValidator(schema),
}
transformers := []Transformer{
    NewFieldTransformer("date", dateFormatter),
    NewDataEnricher(enrichmentService),
}
fileWriter := NewFileWriter("output.json")

// Create and run pipeline
pipeline := NewPipeline(fileReader, validators, transformers, fileWriter)

// Run with context for cancellation
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

err := pipeline.Process(ctx)
if err != nil {
    // Check for specific error types
    var validationErr *ValidationError
    if errors.As(err, &validationErr) {
        fmt.Printf("Validation error in field '%s': %s\n", validationErr.Field, validationErr.Message)
    } else if errors.Is(err, ErrInvalidFormat) {
        fmt.Println("The file format is invalid")
    } else {
        fmt.Printf("Pipeline failed: %v\n", err)
    }
    
    // Print the full error chain
    fmt.Printf("Error chain: %+v\n", err)
}
```

## Instructions

- **Fork** the repository.
- **Clone** your fork to your local machine.
- **Create** a directory named after your GitHub username inside `challenge-12/submissions/`.
- **Copy** the `solution-template.go` file into your submission directory.
- **Implement** the required interfaces and error types.
- **Test** your solution locally by running the test file.
- **Commit** and **push** your code to your fork.
- **Create** a pull request to submit your solution.

## Testing Your Solution Locally

Run the following command in the `challenge-12/` directory:

```bash
go test -v
``` 