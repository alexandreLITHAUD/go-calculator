// Package calculator provides basic mathematical operations
package calculator

import (
	"errors"
	"math"
)

// Calculator represents a basic calculator
type Calculator struct {
	history []Operation
}

// Operation represents a mathematical operation
type Operation struct {
	A      float64 `json:"a"`
	B      float64 `json:"b"`
	Op     string  `json:"operation"`
	Result float64 `json:"result"`
}

// New creates a new Calculator instance
func New() *Calculator {
	return &Calculator{
		history: make([]Operation, 0),
	}
}

// Add performs addition
func (c *Calculator) Add(a, b float64) float64 {
	result := a + b
	c.addToHistory(a, b, "add", result)
	return result
}

// Subtract performs subtraction
func (c *Calculator) Subtract(a, b float64) float64 {
	result := a - b
	c.addToHistory(a, b, "subtract", result)
	return result
}

// Multiply performs multiplication
func (c *Calculator) Multiply(a, b float64) float64 {
	result := a * b
	c.addToHistory(a, b, "multiply", result)
	return result
}

// Divide performs division
func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}

	result := a / b
	c.addToHistory(a, b, "divide", result)
	return result, nil
}

// Power calculates a^b
func (c *Calculator) Power(a, b float64) float64 {
	result := math.Pow(a, b)
	c.addToHistory(a, b, "power", result)
	return result
}

// Sqrt calculates square root
func (c *Calculator) Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("square root of negative number is not allowed")
	}

	result := math.Sqrt(a)
	c.addToHistory(a, 0, "sqrt", result)
	return result, nil
}

// Percentage calculates percentage
func (c *Calculator) Percentage(value, percentage float64) float64 {
	result := value * (percentage / 100)
	c.addToHistory(value, percentage, "percentage", result)
	return result
}

// GetHistory returns the operation history
func (c *Calculator) GetHistory() []Operation {
	return c.history
}

// ClearHistory clears the operation history
func (c *Calculator) ClearHistory() {
	c.history = make([]Operation, 0)
}

// GetLastOperation returns the last operation performed
func (c *Calculator) GetLastOperation() *Operation {
	if len(c.history) == 0 {
		return nil
	}
	return &c.history[len(c.history)-1]
}

// addToHistory adds an operation to the history
func (c *Calculator) addToHistory(a, b float64, operation string, result float64) {
	op := Operation{
		A:      a,
		B:      b,
		Op:     operation,
		Result: result,
	}
	c.history = append(c.history, op)
}
