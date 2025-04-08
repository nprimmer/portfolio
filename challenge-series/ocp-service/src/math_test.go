package main

import (
	"testing"
)

// TestEvaluatePostfix tests the evaluatePostfix function with various postfix expressions.
func TestEvaluatePostfix(t *testing.T) {
	tests := []struct {
		name     string
		postfix  string
		expected int
	}{
		{
			name:     "Simple addition",
			postfix:  "3 4 +",
			expected: 7,
		},
		{
			name:     "Simple subtraction",
			postfix:  "6 2 -",
			expected: 4,
		},
		{
			name:     "Simple multiplication",
			postfix:  "3 4 *",
			expected: 12,
		},
		{
			name:     "Mixed operations",
			postfix:  "5 1 2 + 4 * + 3 -",
			expected: 14,
		},
		{
			name:     "Negative input",
			postfix:  "-3 5 +",
			expected: 2,
		},
		{
			name:     "Negative result",
			postfix:  "3 -5 +",
			expected: -2,
		},
		{
			name:     "Only negative numbers",
			postfix:  "-3 -5 +",
			expected: -8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := evaluatePostfix(tt.postfix)
			if result != tt.expected {
				t.Errorf("evaluatePostfix(%q) = %d, want %d", tt.postfix, result, tt.expected)
			}
		})
	}
}

func TestToPostfix(t *testing.T) {
	tests := []struct {
		name     string
		infix    string
		expected string
	}{
		{
			name:     "Simple addition",
			infix:    "3 + 4",
			expected: "3 4 +",
		},
		{
			name:     "Simple subtraction",
			infix:    "6 - 2",
			expected: "6 2 -",
		},
		{
			name:     "Negative number",
			infix:    "-3 + 5",
			expected: "0 3 - 5 +",
		},
		{
			name:     "Parentheses",
			infix:    "( 3 + 5 ) * 2",
			expected: "3 5 + 2 *",
		},
		{
			name:     "Complex expression",
			infix:    "3 + 4 * 2 / ( 1 - 5 )",
			expected: "3 4 2 * 1 5 - / +",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := toPostfix(tt.infix)
			if result != tt.expected {
				t.Errorf("toPostfix(%q) = %q, want %q", tt.infix, result, tt.expected)
			}
		})
	}
}
