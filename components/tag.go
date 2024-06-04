package components

import "fmt"



func Tag(tag string) string {
	return fmt.Sprintf("<a class='tag' href='/tag/%s'>#%s</a>", tag, tag)
}

func Tags(tags []string) string {
	result := "<div class='tags'>"
	for _, tag := range tags {
		result += Tag(tag)
	}
	return result + "</div>"
}




