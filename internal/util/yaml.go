package util

import (
	"strings"
)

func GetSplitMDData(da string) (frontMatter, postContent string) {
	if len(da) < 10 {
		return "", da
	}

	st := strings.Index(da, "---") + 3
	str := da[st : len(da)-1]
	end := strings.Index(str, "---") + 3
	metaData := da[st:end]

	content := da[end+4:]

	return metaData, content
}

func GetStringInBetweenTwoString(da string, startS string, endS string) (result string, found bool) {
	// s := strings.Index(str, startS)
	// if s == -1 {
	//     return result,false
	// }
	// newS := str[s+len(startS):]
	// e := strings.Index(newS, endS)
	// if e == -1 {
	//     return result,false
	// }
	// result = newS[:e]
	return result, true
}
