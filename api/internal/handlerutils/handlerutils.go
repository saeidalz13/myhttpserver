package handlerutils

import (
	"strings"
)

func ExtractKey(url string) string {
	segments := strings.Split(url, "/")
	return strings.TrimSpace(segments[len(segments)-1])
}

func IsContentTypeJson(contentType string) bool {
	return contentType == "application/json"
}
