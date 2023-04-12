package queue

import (
	"errors"
	"testing"
)

func TestQueueOperations(t *testing.T) {
	testCases := []struct {
		name          string
		operations    []string
		values        []int
		expectedData  []int
		expectedError error
	}{
		{
			name:          "Create and push elements",
			operations:    []string{"push", "push", "push"},
			values:        []int{1, 2, 3},
			expectedData:  []int{1, 2, 3},
			expectedError: nil,
		},
		{
			name:          "Pop without deleting",
			operations:    []string{"push", "push", "pop"},
			values:        []int{1, 2, 0},
			expectedData:  []int{1},
			expectedError: nil,
		},
		{
			name:          "Pop with deleting",
			operations:    []string{"push", "push", "pop"},
			values:        []int{1, 2, 1},
			expectedData:  []int{1},
			expectedError: nil,
		},
		{
			name:          "Pop empty queue",
			operations:    []string{"pop"},
			values:        []int{1},
			expectedData:  []int{},
			expectedError: errors.New("queue is empty"),
		},
		{
			name:          "Create queue and check if empty",
			operations:    []string{"isEmpty"},
			values:        []int{0},
			expectedData:  []int{1},
			expectedError: nil,
		},
		{
			name:          "Push element to queue and check if empty",
			operations:    []string{"push", "isEmpty"},
			values:        []int{1, 0},
			expectedData:  []int{0},
			expectedError: nil,
		},
		{
			name:          "Push multiple elements and pop one by one",
			operations:    []string{"push", "push", "push", "pop", "pop", "pop"},
			values:        []int{1, 2, 3, 1, 1, 1},
			expectedData:  []int{1, 2, 3},
			expectedError: nil,
		},
		{
			name:          "Attempt to pop from empty queue",
			operations:    []string{"pop"},
			values:        []int{1},
			expectedData:  []int{},
			expectedError: errors.New("queue is empty"),
		},
		{
			name:          "Pop and push elements consecutively",
			operations:    []string{"push", "push", "pop", "push", "pop", "push"},
			values:        []int{1, 2, 1, 3, 1, 4},
			expectedData:  []int{1, 2},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			q := Create()
			expectedDataIdx := 0

			for i, op := range tc.operations {
				switch op {
				case "push":
					q.Push(tc.values[i])
				case "pop":
					delete := tc.values[i] == 1
					val, err := q.Pop(delete)
					if err != nil && tc.expectedError == nil {
						t.Errorf("Expected no error, got %v", err)
					} else if err == nil && tc.expectedError != nil {
						t.Errorf("Expected error %v, got no error", tc.expectedError)
					} else if err != nil && tc.expectedError != nil && err.Error() != tc.expectedError.Error() {
						t.Errorf("Expected error %v, got %v", tc.expectedError, err)
					}

					if delete && expectedDataIdx < len(tc.expectedData) && val != tc.expectedData[expectedDataIdx] {
						t.Errorf("Expected value %d, got %d", tc.expectedData[expectedDataIdx], val)
					}
					if delete {
						expectedDataIdx++
					}
				case "isEmpty":
					isEmpty := q.isEmpty()
					expectedIsEmpty := tc.expectedData[expectedDataIdx] == 1
					expectedDataIdx++
					if isEmpty != expectedIsEmpty {
						t.Errorf("Expected isEmpty to be %t, got %t", expectedIsEmpty, isEmpty)
					}
				}
			}
		})
	}
}
