package calculator

import (
	"math"
	"testing"
)

func TestCalculator_Add(t *testing.T) {
	calc := New()

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 2.5, 3.5, 6.0},
		{"negative numbers", -2.5, -3.5, -6.0},
		{"mixed numbers", -2.5, 3.5, 1.0},
		{"with zero", 5.0, 0.0, 5.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%g, %g) = %g, want %g", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculator_Subtract(t *testing.T) {
	calc := New()

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 10.0, 3.0, 7.0},
		{"negative result", 3.0, 10.0, -7.0},
		{"negative numbers", -5.0, -3.0, -2.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Subtract(%g, %g) = %g, want %g", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculator_Multiply(t *testing.T) {
	calc := New()

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 4.0, 3.0, 12.0},
		{"with zero", 5.0, 0.0, 0.0},
		{"negative numbers", -4.0, 3.0, -12.0},
		{"decimal numbers", 2.5, 4.0, 10.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%g, %g) = %g, want %g", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculator_Divide(t *testing.T) {
	calc := New()

	t.Run("normal division", func(t *testing.T) {
		result, err := calc.Divide(10.0, 2.0)
		if err != nil {
			t.Errorf("Divide(10, 2) returned error: %v", err)
		}
		if result != 5.0 {
			t.Errorf("Divide(10, 2) = %g, want 5", result)
		}
	})

	t.Run("division by zero", func(t *testing.T) {
		_, err := calc.Divide(10.0, 0.0)
		if err == nil {
			t.Error("Divide(10, 0) should return an error")
		}
	})
}

func TestCalculator_Power(t *testing.T) {
	calc := New()

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"square", 3.0, 2.0, 9.0},
		{"cube", 2.0, 3.0, 8.0},
		{"power of zero", 5.0, 0.0, 1.0},
		{"power of one", 7.0, 1.0, 7.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Power(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Power(%g, %g) = %g, want %g", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculator_Sqrt(t *testing.T) {
	calc := New()

	t.Run("positive number", func(t *testing.T) {
		result, err := calc.Sqrt(9.0)
		if err != nil {
			t.Errorf("Sqrt(9) returned error: %v", err)
		}
		if result != 3.0 {
			t.Errorf("Sqrt(9) = %g, want 3", result)
		}
	})

	t.Run("zero", func(t *testing.T) {
		result, err := calc.Sqrt(0.0)
		if err != nil {
			t.Errorf("Sqrt(0) returned error: %v", err)
		}
		if result != 0.0 {
			t.Errorf("Sqrt(0) = %g, want 0", result)
		}
	})

	t.Run("negative number", func(t *testing.T) {
		_, err := calc.Sqrt(-1.0)
		if err == nil {
			t.Error("Sqrt(-1) should return an error")
		}
	})
}

func TestCalculator_Percentage(t *testing.T) {
	calc := New()

	tests := []struct {
		name       string
		value      float64
		percentage float64
		expected   float64
	}{
		{"10% of 100", 100.0, 10.0, 10.0},
		{"25% of 200", 200.0, 25.0, 50.0},
		{"50% of 80", 80.0, 50.0, 40.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Percentage(tt.value, tt.percentage)
			if math.Abs(result-tt.expected) > 0.0001 {
				t.Errorf("Percentage(%g, %g) = %g, want %g", tt.value, tt.percentage, result, tt.expected)
			}
		})
	}
}

func TestCalculator_History(t *testing.T) {
	calc := New()

	// Perform some operations
	calc.Add(5, 3)
	calc.Multiply(2, 4)

	history := calc.GetHistory()
	if len(history) != 2 {
		t.Errorf("Expected 2 operations in history, got %d", len(history))
	}

	lastOp := calc.GetLastOperation()
	if lastOp == nil {
		t.Error("Expected last operation, got nil")
	} else if lastOp.Op != "multiply" {
		t.Errorf("Expected last operation to be multiply, got %s", lastOp.Op)
	}

	calc.ClearHistory()
	history = calc.GetHistory()
	if len(history) != 0 {
		t.Errorf("Expected empty history after clear, got %d operations", len(history))
	}
}

func BenchmarkCalculator_Add(b *testing.B) {
	calc := New()
	for i := 0; i < b.N; i++ {
		calc.Add(float64(i), float64(i+1))
	}
}

func BenchmarkCalculator_Divide(b *testing.B) {
	calc := New()
	for i := 0; i < b.N; i++ {
		calc.Divide(float64(i+1), 2.0)
	}
}
