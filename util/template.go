package util

import (
	"embed"
	"fmt"
	"strings"

	"github.com/kolukattai/kblog/models"
)

type template struct {
	base   string
	folder embed.FS
}

func HtmlTemplate(base string, folder embed.FS) *template {
	return &template{base: base, folder: folder}
}

func (st *template) PopulateString(key, val string) *template {
	if len(val) == 0 {
		return st
	}
	st.base = strings.ReplaceAll(st.base, fmt.Sprintf("{{%v}}", key), val)
	return st
}

func (st *template) PopulateHTML(key, file string) *template {
	if len(file) == 0 {
		return st
	}
	val, err := st.folder.ReadFile(file)
	if err != nil {
		panic(err)
	}
	st.base = strings.ReplaceAll(st.base, fmt.Sprintf("<!--{{%v}}-->", key), string(val))
	return st
}

func (st *template) PopulateComponent(key string, data models.Component) *template {
	if len(data) == 0 {
		return st
	}
	st.base = strings.ReplaceAll(st.base, fmt.Sprintf("<!--{{%v}}-->", key), string(data))
	return st
}

func (st *template) PopulateComponents(key string, data []models.Component) *template {
	if len(data) == 0 {
		return st
	}
	da := ""
	for _, v := range data {
		da += string(v)
	}
	st.base = strings.ReplaceAll(st.base, fmt.Sprintf("<!--{{%v}}-->", key), da)
	return st
}

func (st *template) PopulateMarkdown(key, data string) *template {
	if len(data) == 0 {
		return st
	}
	st.base = strings.ReplaceAll(st.base, fmt.Sprintf("<!--{{%v}}-->", key), string(MDtoHTML([]byte(data))))
	return st
}

func (st *template) Result() string {
	return st.base
}
