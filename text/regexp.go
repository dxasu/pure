package text

import (
	"fmt"
	"regexp"
)

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

func RegSubSplit(regexStr, content string) [][]string {
	re := regexp.MustCompile(regexStr)
	parts := re.Split(content, -1)
	matches := re.FindAllString(content, -1)
	var result [][]string
	for i, part := range parts {
		result = append(result, []string{part})
		if i < len(matches) {
			subs := re.FindStringSubmatch(matches[i])
			result = append(result, subs[1:])
		}
	}
	return result
}

type ColorRegex string

const (
	ColorRegexTruth ColorRegex = `\033\[38;2;(\d{1,3});(\d{1,3});(\d{1,3})m([^\033]+)\033\[0m` // \033[38;2;R;G;Bm...\033[0m
	ColorRegexSharp ColorRegex = `#([0-9a-f]{2})([0-9a-f]{2})([0-9a-f]{2})([^#]*)#`            // #RRGGBBtext#
	ColorRegexAnsi  ColorRegex = `\033\[3([0-7])m([^\033]+)\033\[0m`                           // \033[3Xm...\033[0m (30-37 foreground colors
)

func (cr ColorRegex) RegSubSplit(content string) [][]string {
	return RegSubSplit(string(cr), content)
}

// ColorFormat 使用 ANSI 转义序列为终端文本添加颜色
// \033[38;2;ff;00;00mhello world\033[0m
func ColorFormat(f uint32, text string) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", uint8(f>>16), uint8(f>>8), uint8(f), text)
}
