package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re, _ := regexp.Compile(`(?m)^(\[ERR\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re, _ := regexp.Compile(`<[\*~=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re, _ := regexp.Compile(`(?mi)\".*password\"`)
	var count int
	for _, row := range lines {
		count += len(re.FindAllString(row, -1))
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re, _ := regexp.Compile(`end-of-line(\d+)`)
	return strings.Join(re.Split(text, -1), "")
}

func TagWithUserName(lines []string) []string {
	re, _ := regexp.Compile(`User\s+[a-zA-Z]+[0-9]*`)
	re_user, _ := regexp.Compile(`\s+`)
	var response []string
	for _, line := range lines {
		if re.MatchString(line) {
			user := re_user.Split(re.FindStringSubmatch(line)[0], -1)[1]
			response = append(response, fmt.Sprintf("[USR] %s %s", user, line))
		} else {
			response = append(response, line)
		}
	}
	return response
}
