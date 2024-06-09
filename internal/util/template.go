package util

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"os"
	"strings"

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

func (st *htmlTemplate) MdData(post string, data any, conf *models.Config) *htmlTemplate {
	val, err := os.ReadFile(fmt.Sprintf("posts/%s.md", post))
	if err != nil {
		val = []byte("")
	}
	arr := strings.Split(string(val), "---")

	metaData := ""
	content := string(val)

	if len(arr) == 3 {
		metaData = arr[1]
		content = arr[2]
	}

	mData := models.PageData{}

	err = yaml.Unmarshal([]byte(metaData), &mData)
	if err != nil {
		Error(err.Error())
	}

	pageContent := MDtoHTML([]byte(content))

	var result bytes.Buffer

	err = st.templ.
		Execute(&result, models.MDPageData{
			Content:         template.HTML(pageContent),
			MetaData:        mData,
			Data:            data,
			DefaultMetaData: conf,
			PageType:        st.fileType,
			Global: struct {
				Tags       []string
				Catagories []string
			}{
				Tags:       global.Tags,
				Catagories: global.Categories,
			},
		})
	if err != nil {
		Error(err.Error())
	}
	st.result = &result
	return st
}

func (st *htmlTemplate) Result() string {
	return st.result.String()
}

func (st *htmlTemplate) MinifyResult() string {
	return string(Minify().HTML([]byte(st.Result())))
}
