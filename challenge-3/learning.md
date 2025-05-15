# Learning Materials for Employee Data Management

## Structs, Methods, and Slices in Go

This challenge focuses on using structs to represent objects, methods to add behavior, and slices to manage collections in Go.

### Structs in Go

Structs in Go are composite data types that group together fields with different data types under a single name. They're similar to classes in object-oriented languages but without inheritance.

```go
// Define a struct
type Employee struct {
    ID        int
    FirstName string
    LastName  string
    Email     string
    Position  string
    Salary    float64
}

// Create an instance of the struct
employee := Employee{
    ID:        1,
    FirstName: "John",
    LastName:  "Doe",
    Email:     "john.doe@example.com",
    Position:  "Developer",
    Salary:    75000.00,
}

// Access fields
fmt.Println(employee.FirstName) // John

// Update fields
employee.Salary = 80000.00
```

### Methods in Go

In Go, methods are functions associated with a particular type. You can add methods to structs to encapsulate behavior.

```go
// Method with a receiver
func (e Employee) FullName() string {
    return e.FirstName + " " + e.LastName
}

// Method that modifies the receiver (note the pointer receiver)
func (e *Employee) GiveRaise(percentage float64) {
    e.Salary = e.Salary * (1 + percentage/100)
}

// Usage
fmt.Println(employee.FullName())  // John Doe
employee.GiveRaise(10)
fmt.Println(employee.Salary)      // 88000
```

The difference between value receivers `(e Employee)` and pointer receivers `(e *Employee)` is important:
- Value receivers get a copy of the struct
- Pointer receivers get a reference and can modify the original

### Slices in Go

Slices are flexible, dynamic arrays that can grow or shrink. They're perfect for managing collections of data.

```go
// Declare a slice of Employees
var employees []Employee

// Create a slice with initial capacity
employees := make([]Employee, 0, 10)

// Add elements
employees = append(employees, employee1)
employees = append(employees, employee2)

// Access elements
fmt.Println(employees[0].FirstName)  // Access first employee

// Slice operations
firstThree := employees[:3]  // First three employees
lastTwo := employees[len(employees)-2:]  // Last two employees
```

### Common Slice Operations

```go
// Length and capacity
fmt.Println(len(employees))  // Number of employees
fmt.Println(cap(employees))  // Capacity of the slice

// Iterate over a slice
for i, emp := range employees {
    fmt.Printf("%d: %s %s\n", i, emp.FirstName, emp.LastName)
}

// Remove an element (at index i)
employees = append(employees[:i], employees[i+1:]...)

// Filter a slice
var highPaidEmployees []Employee
for _, emp := range employees {
    if emp.Salary > 100000 {
        highPaidEmployees = append(highPaidEmployees, emp)
    }
}
```

### Composition over Inheritance

Go favors composition over inheritance. Instead of extending a class, embed one struct into another:

```go
type Address struct {
    Street  string
    City    string
    State   string
    ZipCode string
}

type EmployeeWithAddress struct {
    Employee
    Address Address
}

// Usage
emp := EmployeeWithAddress{
    Employee: Employee{ID: 1, FirstName: "Jane"},
    Address:  Address{City: "San Francisco"},
}
fmt.Println(emp.FirstName)  // Access Employee fields directly
fmt.Println(emp.Address.City)  // Access Address fields
```

### Sorting Slices

Go's `sort` package can be used to sort slices:

```go
import "sort"

// Sort employees by salary
sort.Slice(employees, func(i, j int) bool {
    return employees[i].Salary < employees[j].Salary
})

// Sort employees by last name
sort.Slice(employees, func(i, j int) bool {
    return employees[i].LastName < employees[j].LastName
})
```

## Further Reading

- [Go Tour: Structs](https://tour.golang.org/moretypes/2)
- [Go by Example: Structs](https://gobyexample.com/structs)
- [Go by Example: Methods](https://gobyexample.com/methods)
- [Go Slices: Usage and Internals](https://blog.golang.org/slices-intro) 