package urlutils

import (
	"errors"
	"strings"
)

func ExtractKey(url string) (string, error) {
	segments := strings.Split(url, "/")
	if len(segments) > 3 {
		return "", errors.New("invalid url schema to get key")
	}

	key := strings.TrimSpace(segments[2])
	if key == "" {
		return key, errors.New("no key has been provided")
	}
	return key, nil
}
