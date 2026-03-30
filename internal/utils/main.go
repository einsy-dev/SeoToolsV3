package utils

import (
	"regexp"
)

var reg = regexp.MustCompile(`^(?:https?:\/\/)?(?:www\.)?([^\/\s]+\.[^\/\s?]+)`)

func FormatDomain(url string) string {
	matches := reg.FindStringSubmatch(url)
	if len(matches) > 1 {
		return matches[1]
	}
	return url
}
