package util

import (
	"fmt"
	"regexp"
)

func RegexMatch(input string, regex string, captureGroup int) (string, error) {
	reg, err := regexp.Compile(regex)
	if err != nil {
		return "", err
	}

	matches := reg.FindAllStringSubmatch(input, -1)
	if len(matches) == 0 {
		return "", fmt.Errorf("no match for regex: %s in %s", regex, input)
	}
	if len(matches[0]) <= captureGroup {
		return "", fmt.Errorf("capture group out of bounds: "+
			"%d (bound / regex_matches) <= %d (captureGroup )",
			len(matches[0]), captureGroup)
	}
	return matches[0][captureGroup], nil
}
