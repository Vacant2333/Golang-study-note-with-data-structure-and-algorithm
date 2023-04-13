package set_map

import (
	"testing"
)

func TestSetOperations(t *testing.T) {
	testCases := []struct {
		name     string
		elements []int
	}{
		{
			name:     "Test case 1",
			elements: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Test case 2",
			elements: []int{10, 20, 30, 40, 50},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mySet := New()

			// Test Add
			for _, el := range tc.elements {
				mySet.Add(el)
			}

			// Test Contains
			for _, el := range tc.elements {
				if !mySet.Contains(el) {
					t.Errorf("Element %d should be in the set", el)
				}
			}

			// Test Remove
			mySet.Remove(tc.elements[0])
			if mySet.Contains(tc.elements[0]) {
				t.Errorf("Element %d should not be in the set", tc.elements[0])
			}

			// Test Union
			otherSet := New()
			for i := 100; i <= 105; i++ {
				otherSet.Add(i)
			}
			mySet.Union(otherSet)

			for i := 100; i <= 105; i++ {
				if !mySet.Contains(i) {
					t.Errorf("Element %d should be in the set after union", i)
				}
			}
		})
	}
}
