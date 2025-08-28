package text

import "regexp"

func RegSplit(re *regexp.Regexp, content string) []string {
	parts := re.Split(content, -1)
	matches := re.FindAllString(content, -1)
	var result []string
	for i, part := range parts {
		result = append(result, part)
		if i < len(matches) {
			result = append(result, matches[i])
		}
	}
	return result
}
