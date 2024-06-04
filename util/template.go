package util

import (
	"embed"
	"fmt"
	"strings"
)

type template struct {
	base   string
	folder embed.FS
}

func HtmlTemplate(base string, folder embed.FS) *template {
	return &template{base: base, folder: folder}
}

func (st *template) PopulateString(key, val string) *template {
	st.base = strings.ReplaceAll(st.base, fmt.Sprintf("{{%v}}", key), val)
	return st
}

func (st *template) PopulateHTML(key, file string) *template {
	val, err := st.folder.ReadFile(file)
	if err != nil {
		panic(err)
	}
	st.base = strings.ReplaceAll(st.base, fmt.Sprintf("{{%v}}", key), string(val))
	return st
}

func (st *template) PopulateMarkdown(key, data string) *template {
	st.base = strings.ReplaceAll(st.base, fmt.Sprintf("{{%v}}", key), string(MDtoHTML([]byte(data))))
	return st
}

func (st *template) Result() string {
	return st.base
}
