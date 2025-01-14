# Scoreboard for challenge-3
| Username   | Passed Tests | Total Tests |
|------------|--------------|-------------|
| RezaSi | 5 | 5 |


[View the Scoreboard](SCOREBOARD.md)

# Challenge 3: Employee Data Management

## Problem Statement

You are tasked with managing a list of employees with the following details: `ID`, `Name`, `Age`, and `Salary`. Implement a `Manager` struct that provides the following functionalities:

1. **AddEmployee**: Add a new employee to the list.
2. **RemoveEmployee**: Remove an employee based on their ID.
3. **GetAverageSalary**: Calculate the average salary of all employees.
4. **FindEmployeeByID**: Retrieve an employee's details by their ID.

## Structures and Function Signatures

Define the following structures:

```go
type Employee struct {
    ID     int
    Name   string
    Age    int
    Salary float64
}

type Manager struct {
    Employees []Employee
}
```

Implement the following methods:

```go
func (m *Manager) AddEmployee(e Employee)
func (m *Manager) RemoveEmployee(id int)
func (m *Manager) GetAverageSalary() float64
func (m *Manager) FindEmployeeByID(id int) *Employee
```

## Instructions

- **Fork** the repository and **clone** your fork.
- **Create** your submission directory inside `challenge-3/submissions/`.
- **Copy** the `solution-template.go` file into your submission directory.
- Implement and **test** your solution using the test file.
- **Submit** your solution by creating a pull request.

## Sample Usage

```go
manager := Manager{}
manager.AddEmployee(Employee{ID: 1, Name: "Alice", Age: 30, Salary: 70000})
manager.AddEmployee(Employee{ID: 2, Name: "Bob", Age: 25, Salary: 65000})
manager.RemoveEmployee(1)
averageSalary := manager.GetAverageSalary()
employee := manager.FindEmployeeByID(2)
```

## Testing Your Solution Locally

Run the following command in the `challenge-3` directory:

```bash
go test -v
```