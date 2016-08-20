package client

import (
	"fmt"
	"regexp"
	"time"
)

var (
	valueRegexp     = regexp.MustCompile(`^"(.*)"$`)
	hashFieldRegexp = regexp.MustCompile(`^([a-zA-Z0-9_]+):"(.*)"$`)
)

func parseValue(str string) (string, error) {
	if matches := valueRegexp.FindStringSubmatch(str); len(matches) > 1 {
		return matches[1], nil
	}
	return "", fmt.Errorf("Invalid value format: %s", str)
}

func parseHashField(str string) (string, string, error) {
	if matches := hashFieldRegexp.FindStringSubmatch(str); len(matches) > 2 {
		return matches[1], matches[2], nil
	}
	return "", "", fmt.Errorf("Invalid hash field format: %s", str)
}

func parseTTL(str string) (time.Duration, error) {
	d, err := time.ParseDuration(str)
	if err != nil {
		return time.Duration(0), err
	}
	return d, nil
}
