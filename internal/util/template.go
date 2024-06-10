package util

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"os"
	"strings"
	"time"

	"github.com/kolukattai/kblog/internal/global"
	"github.com/kolukattai/kblog/internal/models"
	"gopkg.in/yaml.v3"
)

type htmlTemplate struct {
	templ    *template.Template
	result   *bytes.Buffer
	metaData *models.PageData
	fileType models.PageType
}

func HtmlTemplate(folder embed.FS, fileType models.PageType, files ...string) *htmlTemplate {
	fileNames := []string{
		"templates/layout.html",
		"templates/_head.html",
		"templates/_header.html",
		"templates/_footer.html",
		"templates/_drawer.html",
		fmt.Sprintf("templates/%v.html", string(fileType)),
	}
	if len(files) > 0 {
		for _, v := range files {
			fileNames = append(fileNames, fmt.Sprintf("templates/%v.html", v))
		}
	}
	tmpl, err := template.ParseFS(
		folder,
		fileNames...,
	)
	if err != nil {
		Error(err.Error())
	}
	return &htmlTemplate{templ: tmpl, fileType: fileType}
}

func (st *htmlTemplate) MdData(post string, data any, conf *models.Config, manualMetaData ...models.PageData) *htmlTemplate {
	val, err := os.ReadFile(fmt.Sprintf("posts/%s.md", post))
	if err != nil {
		val = []byte("")
	}

	metaData, content := GetSplitMDData(string(val))

	mData := models.PageData{}

	err = yaml.Unmarshal([]byte(metaData), &mData)
	if err != nil {
		Error(err.Error())
	}

	pageContent := MDtoHTML([]byte(content))

	var result bytes.Buffer

	md := models.MDPageData{
		Content:         template.HTML(pageContent),
		MetaData:        mData,
		Data:            data,
		DefaultMetaData: conf,
		PageType:        st.fileType,
		Year: time.Now().String()[0:4],
		Global: struct {
			Tags       []string
			Catagories []string
		}{
			Tags:       global.Tags,
			Catagories: global.Categories,
		},
	}

	for _, v := range manualMetaData {
		md.MetaData = v
	}

	err = st.templ.
		Execute(&result, md)
	if err != nil {
		Error(err.Error())
	}
	st.result = &result
	return st
}

func (st *htmlTemplate) Result() string {
	v := st.result.String()
	v = strings.ReplaceAll(v, "[x]", "&#x2705;")
	return v
}

func (st *htmlTemplate) MinifyResult() string {
	return string(Minify().HTML([]byte(st.Result())))
}
