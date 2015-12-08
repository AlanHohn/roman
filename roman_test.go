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
	{"X", 10},
	{"XI", 11},
	{"XIV", 14},
	{"XX", 20},
	{"XL", 40},
	{"XLIX", 49},
	{"L", 50},
	{"LIV", 54},
	{"LVI", 56},
	{"LIX", 59},
	{"LXXXVIII", 88},
	{"C", 100},
	{"CCC", 300},
	{"D", 500},
	{"M", 1000},
	{"MDCCLXXVI", 1776},
	{"MMXV", 2015},
	{"MMMM", 4000},
}

var invalidTests = []string{
	"IIX",
	"IIII",
	"VV",
	"VX",
	"XXXX",
	"LL",
	"LC",
	"CCCC",
	"DD",
	"DM",
	"Z",
	"ZI",
	"IZ",
	"123",
	"@",
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

func TestInvalid(t *testing.T) {
	for _, tt := range invalidTests {
		res, err := RomanToInt(tt)
		if err == nil {
			t.Errorf("Expected error for input %v but received %v", tt, res)
		}
		if err != ErrInvalidFormat {
			t.Errorf("Unexpected error for input %v: %v", tt, err)
		}
	}
}
