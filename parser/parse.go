package parser

import (
	"fmt"
	"os"
	"strings"

	"github.com/kolukattai/kblog/components"
	"github.com/kolukattai/kblog/global"
	"github.com/kolukattai/kblog/models"
	"github.com/kolukattai/kblog/util"
	"gopkg.in/yaml.v3"
)

type Options struct {
	LandingImage bool
	Tags         bool
	Author       bool
	Config       *models.Config
}

func Parse(fileName string, options Options) (string, *models.PageData) {
	postByte, err := os.ReadFile(fmt.Sprintf("posts/%s.md", fileName))
	if err != nil {
		panic(err)
	}

	pageMetaData := ""
	pageContent := ""

	pageDataArr := strings.Split(string(postByte), "---")
	if len(pageDataArr) == 3 {
		pageContent = pageDataArr[2]
	}

	st := util.HtmlTemplate(ParseData(pageContent), global.TemplateFolder)

	metaData := models.PageData{}
	if len(pageDataArr) == 3 {
		pageMetaData = pageDataArr[1]
	}
	err = yaml.Unmarshal([]byte(pageMetaData), &metaData)
	if err != nil {
		panic(err)
	}

	util.IterateStruct(metaData, func(key string, value any) {
		st.PopulateString(strings.ToLower(key), fmt.Sprintf("%v", value))
	})

	landingImage := components.Img("landing image", metaData.LandingImage)
	if len(landingImage) > 0 {
		st.PopulateComponent("_landingImage", landingImage)
	}

	tags := components.Tags(metaData.GetTags())
	if len(tags) > 0 {
		st.PopulateComponent("_tags", tags)
	}

	st.PopulateString("ga_id", options.Config.GoogleAnalytics)

	return st.Result(), &metaData
}

func ParseData(data string) string {

	st := util.HtmlTemplate("<!--{{main}}-->", global.TemplateFolder).
		PopulateHTML("main", "templates/base.html").
		PopulateHTML("_header", "templates/header.html").
		PopulateHTML("_footer", "templates/footer.html").
		PopulateHTML("_ga", "templates/ga.html").
		PopulateMarkdown("_content", data)

	return st.Result()
}

func ParsePageMetaData(page string) models.PageData {
	pageMetaData := ""

	pageDataArr := strings.Split(page, "---")

	metaData := models.PageData{}
	if len(pageDataArr) == 3 {
		pageMetaData = pageDataArr[1]
	}
	err := yaml.Unmarshal([]byte(pageMetaData), &metaData)
	if err != nil {
		panic(err)
	}
	return metaData
}
