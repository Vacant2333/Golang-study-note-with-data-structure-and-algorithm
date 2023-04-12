package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	testCases := []struct {
		name        string
		operations  []string
		values      []int
		expected    []int
		expectError bool
	}{
		{
			name:        "Push and Pop",
			operations:  []string{"Push", "Push", "Pop", "Pop"},
			values:      []int{1, 2, 0, 0},
			expected:    []int{0, 0, 2, 1},
			expectError: false,
		},
		{
			name:        "Empty Stack",
			operations:  []string{"Pop"},
			values:      []int{0},
			expected:    []int{0},
			expectError: true,
		},
		{
			name:        "Push and Peek",
			operations:  []string{"Push", "Push", "Pop", "Pop"},
			values:      []int{3, 4, 0, 0},
			expected:    []int{0, 0, 4, 4},
			expectError: false,
		},
		{
			name:        "Mixed Operations",
			operations:  []string{"Push", "Push", "Pop", "Push", "Pop", "Pop"},
			values:      []int{5, 6, 0, 7, 0, 0},
			expected:    []int{0, 0, 6, 0, 7, 5},
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stack := Create()

			for i, op := range tc.operations {
				switch op {
				case "Push":
					stack.Push(tc.values[i])
				case "Pop":
					delete := true
					if tc.name == "Push and Peek" && i == 2 {
						delete = false
					}
					val, err := stack.Pop(delete)

					if err != nil && !tc.expectError {
						t.Errorf("Unexpected error: %v", err)
					}

					if err == nil && tc.expectError {
						t.Error("Expected error, but got none")
					}

					if val != tc.expected[i] {
						t.Errorf("Expected value %d, got %d", tc.expected[i], val)
					}
				}
			}
		})
	}
}
