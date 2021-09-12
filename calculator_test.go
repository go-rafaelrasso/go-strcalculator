package calculator

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		value    string
		expected int
	}{
		{value: "", expected: 0},
		{value: "   ", expected: 0},
		{value: "2", expected: 2},
		{value: "1,2,5", expected: 8},
		{value: "1,2,3,4,5,6,7", expected: 28},
		{value: "1,1,1,1,1,1,1", expected: 7},
		{value: "1,0,10,100", expected: 111},
		{value: "1\n,2,3", expected: 6},
		{value: "1,\n2,4", expected: 7},
		{value: "1,2,9\n", expected: 12},
		{value: "\n", expected: 0},
		{value: "//;\n1;3;4", expected: 8},
		{value: "//$\n1$2$3", expected: 6},
		{value: "//@\n2@3@8", expected: 13},
		{value: "//;-2;3;4;-5;6;-7", expected: 0},
		{value: "//@\n2@3@8@1000@1001", expected: 1013},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("ShouldAddReturn_%d_WhenStringIs_%s", test.expected, test.value), func(t *testing.T) {
			result, _ := Add(test.value)
			if result != test.expected {
				t.Errorf("Should Add return %d when string is %s, but returned %d", test.expected, test.value, result)
			}
		})
	}
}

func TestAddException(t *testing.T) {
	_, err := Add("//;-2;3;4;-5;6;-7")
	if err.Error() != "Negatives not allowed! [-2 -5 -7]" {
		t.Errorf("Should not allow negatives but allowed")
	}
}
