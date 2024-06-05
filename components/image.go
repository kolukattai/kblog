package components

import (
	"fmt"

	"github.com/kolukattai/kblog/models"
)

func Img(alt, src string) models.Component {
	if len(src) != 0 {
		return models.Component(fmt.Sprintf("<img src='%s' alt='%s'/>", src, alt))
	} else {
		return models.Component("")
	}
}
