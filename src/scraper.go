package main

import (
	"regexp"
	"strings"
)

// ExtractTags takes html of the steam store page and extracts all of the game tags
func ExtractTags(html string) []string {
	re := regexp.MustCompile("class=\"app_tag\"[^>]*>([^<]*)")

	matches := re.FindAllStringSubmatch(html, -1)

	tags := make([]string, len(matches))

	for i, match := range matches {
		tags[i] = strings.TrimSpace(match[1])
	}

	return tags
}
