package util

import "strings"

func GetSplitMDData(val string) (frontMatter, postContent string) {

	arr := strings.Split(string(val), "---")

	metaData := ""
	content := string(val)

	if len(arr) == 3 {
		metaData = arr[1]
		content = arr[2]
	}

	return metaData, content
}
