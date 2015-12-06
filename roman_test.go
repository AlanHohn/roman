package roman

import "testing"

var validTests = []struct {
	input    string
	expected int
}{
	{"", 0},
	{"I", 1},
	{"II", 2},
	{"III", 3},
	{"IV", 4},
	{"V", 5},
	{"VI", 6},
	{"IX", 9},
	{"XI", 11},
	{"XX", 20},
	{"XL", 40},
	{"LIV", 54},
	{"LVI", 56},
	{"MMXV", 2015},
}

func TestValid(t *testing.T) {
	for _, tt := range validTests {
		res, err := RomanToInt(tt.input)
		if err != nil {
			t.Errorf("Unexpected error for input %v: %v", tt.input, err)
		}
		if res != tt.expected {
			t.Errorf("Unexpected value for input %v: %v", tt.input, res)
		}
	}
}
