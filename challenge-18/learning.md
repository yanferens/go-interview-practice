# Learning Materials for Temperature Converter

## Working with Floating Point Numbers in Go

### Basic Float Types

Go has two floating-point types:
- `float32` - 32-bit floating-point number (single precision)
- `float64` - 64-bit floating-point number (double precision)

For most applications, `float64` is recommended as it provides more precision.

```go
var temperature float64 = 23.5
```

### Arithmetic Operations

Floating-point numbers support standard arithmetic operations:

```go
var celsius float64 = 25.0
var fahrenheit float64

// Addition
fahrenheit = celsius + 20.0 // 45.0

// Subtraction
celsius = celsius - 5.0 // 20.0

// Multiplication
fahrenheit = celsius * 1.8 // 36.0

// Division
celsius = fahrenheit / 1.8 // 20.0
```

### Precision and Rounding

Floating-point arithmetic is subject to precision limitations due to how numbers are represented in binary. This can lead to small inaccuracies.

For example, the expression `0.1 + 0.2` might not yield exactly `0.3` but something very close like `0.30000000000000004`.

To handle this, you can round numbers to a specific number of decimal places:

```go
import "math"

func Round(value float64, decimals int) float64 {
    precision := math.Pow10(decimals)
    return math.Round(value*precision) / precision
}

// Example usage
x := 0.1 + 0.2                  // 0.30000000000000004
rounded := Round(x, 1)          // 0.3
```

### Temperature Conversion

The standard formulas for converting between Celsius and Fahrenheit are:

1. **Celsius to Fahrenheit**: F = C × 9/5 + 32
2. **Fahrenheit to Celsius**: C = (F - 32) × 5/9

Here's how to implement these conversions in Go:

```go
func CelsiusToFahrenheit(celsius float64) float64 {
    return celsius*9.0/5.0 + 32.0
}

func FahrenheitToCelsius(fahrenheit float64) float64 {
    return (fahrenheit - 32.0) * 5.0 / 9.0
}
```

### Formatting Float Output

When displaying floating-point numbers, it's often desirable to format them with a specific number of decimal places. You can use the `fmt` package for this:

```go
import "fmt"

celsius := 25.0
fahrenheit := CelsiusToFahrenheit(celsius)

// Print with 2 decimal places
fmt.Printf("%.2f°C is equal to %.2f°F\n", celsius, fahrenheit)
// Output: 25.00°C is equal to 77.00°F
```

### Constants and Mathematical Operations

For mathematical operations, Go provides the `math` package with constants and functions:

```go
import "math"

// Constants
pi := math.Pi                // 3.141592653589793
e := math.E                   // 2.718281828459045

// Functions
absValue := math.Abs(-15.5)  // 15.5
sqrt := math.Sqrt(16)        // 4.0
power := math.Pow(2, 3)      // 8.0 (2³)
```

### Error Handling for Invalid Inputs

When working with temperature conversions, you might want to validate inputs or handle edge cases:

```go
// Check for absolute zero violation in Celsius
func ValidateCelsius(celsius float64) error {
    if celsius < -273.15 {
        return fmt.Errorf("temperature below absolute zero: %f°C", celsius)
    }
    return nil
}

// Check for absolute zero violation in Fahrenheit
func ValidateFahrenheit(fahrenheit float64) error {
    if fahrenheit < -459.67 {
        return fmt.Errorf("temperature below absolute zero: %f°F", fahrenheit)
    }
    return nil
}
```

### Other Temperature Scales

While this challenge focuses on Celsius and Fahrenheit, there are other temperature scales:

1. **Kelvin (K)**: K = C + 273.15
2. **Rankine (°R)**: °R = F + 459.67
3. **Réaumur (°Ré)**: °Ré = C × 0.8

For a complete temperature conversion library, you might implement conversions between all these scales.

## Further Reading

- [Go by Example: Floating Point](https://gobyexample.com/floating-point)
- [IEEE 754 standard](https://en.wikipedia.org/wiki/IEEE_754) - the standard for floating-point arithmetic
- [Temperature conversion formulas](https://en.wikipedia.org/wiki/Conversion_of_scales_of_temperature)
- [math package documentation](https://pkg.go.dev/math) 