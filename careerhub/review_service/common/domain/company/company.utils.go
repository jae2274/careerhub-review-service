package company

import "strings"

func RefineNameForSearch(input string) string {
	index := strings.Index(input, "(")
	if index == -1 {
		return input
	}

	return strings.TrimSpace(input[:index])
}
