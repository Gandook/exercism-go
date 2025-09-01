package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re, err := regexp.Compile(`\[(TRC)|(DBG)|(INF)|(WRN)|(ERR)|(FTL)\]`)
    if err != nil {
        panic("the regexp did not compile")
    }

    return len(text) >= 5 && re.MatchString(text[:5])
}

func SplitLogLine(text string) []string {
	re, err := regexp.Compile(`<[~\*\=\-]*>`)
    if err != nil {
        panic("the regexp did not compile")
    }

    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re, err := regexp.Compile(`".*((?i)password).*"`)
	if err != nil {
        panic("the regexp did not compile")
    }

    answer := 0

    for _, line := range lines {
        if re.MatchString(line) {
            answer++
        }
    }

    return answer
}

func RemoveEndOfLineText(text string) string {
	re, err := regexp.Compile(`end\-of\-line\d+`)
    if err != nil {
        panic("the regexp did not compile")
    }

    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re, err := regexp.Compile(`User\s+((?i)[a-z\d]*)`)
    if err != nil {
        panic("the regexp did not compile")
    }

    results := make([]string, 0)
    var matches []string

    for _, line := range lines {
        if matches = re.FindStringSubmatch(line); matches != nil {
            results = append(results, "[USR] " + matches[1] + " " + line)
        } else {
            results = append(results, line)
        }
    }

    return results
}
