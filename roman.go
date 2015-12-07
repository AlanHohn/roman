package roman

import "errors"

var ErrInvalidFormat = errors.New("Invalid Format")

type numeralData struct {
	value    int
	max      int
	leftSide bool
}

var letterToNumeral = map[byte]numeralData{
	'M': {1000, 0, false},
	'D': {500, 1, false},
	'C': {100, 3, true},
	'L': {50, 1, false},
	'X': {10, 3, true},
	'V': {5, 1, false},
	'I': {1, 3, true},
}

func RomanToInt(roman string) (int, error) {
	if len(roman) == 0 {
		return 0, nil
	}
	val := 0
	maxSeen := 1
	lastSeen := 0
	repeat := 1
	leftSideFlag := false
	for i := len(roman) - 1; i >= 0; i-- {
		current := letterToNumeral[roman[i]]
		// Check for excessive repeating
		if lastSeen == current.value {
			repeat++
			if current.max > 0 && repeat > current.max {
				return 0, ErrInvalidFormat
			}
		} else {
			lastSeen = current.value
			repeat = 1
		}
		// Subtract or add?
		if current.value < maxSeen {
			if !current.leftSide || leftSideFlag {
				return 0, ErrInvalidFormat
			}
			leftSideFlag = true
			val -= current.value
		} else {
			leftSideFlag = false
			maxSeen = current.value
			val += current.value
		}
	}
	return val, nil
}
