package components

import "fmt"

func Img(alt, src string) string {
	return fmt.Sprintf("<img src='%s' alt='%s'/>", src, alt)
}
