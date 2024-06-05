package components

import (
	"fmt"

	"github.com/kolukattai/kblog/models"
)

func Tag(tag string) models.Component {
	return models.Component(fmt.Sprintf("<a class='tag' href='/tag/%s'>#%s</a>", tag, tag))
}

func Tags(tags []string) models.Component {
	result := "<div class='tags'>"
	for _, tag := range tags {
		result += string(Tag(string(tag)))
	}
	return models.Component(result + "</div>")
}
