package roman

var letterToNumeral = map[byte]int {
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
}

func RomanToInt(roman string) (int, error) {
	if len(roman) == 0 {
		return 0, nil
	}
	val := 0
	maxSeen := 1
	for i := len(roman)-1; i >= 0; i-- {
		current := letterToNumeral[roman[i]]
		if current < maxSeen {
			val -= current
		} else {
			maxSeen = current
			val += current
		}
	}
	return val, nil
}
