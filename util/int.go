package util

import "strconv"

func IntToNewBaseFixedLength(number, base, length int) string {
	if length == 0 {
		return "0"
	}

	otherBasis := strconv.FormatInt(int64(number), base)
	for len(otherBasis) < length {
		otherBasis = "0" + otherBasis
	}
	return otherBasis
}
