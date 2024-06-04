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

func Parse(fileName string, options Options) string {
	postByte, err := os.ReadFile(fmt.Sprintf("posts/%s.md", fileName))
	if err != nil {
		panic(err)
	}

	pageDataArr := strings.Split(string(postByte), "---")

	st := util.HtmlTemplate(ParseData(pageDataArr[2]), global.TemplateFolder)

	metaData := models.PageData{}
	err = yaml.Unmarshal([]byte(pageDataArr[1]), &metaData)
	if err != nil {
		panic(err)
	}

	util.IterateStruct(metaData, func(key string, value any) {
		st.PopulateString(strings.ToLower(key), fmt.Sprintf("%v", value))
	})

	if options.LandingImage {
		st.PopulateString("_landingImage", components.Img("landing image", metaData.LandingImage))
	}
	if options.Tags {
		st.PopulateString("_tags", components.Tags(metaData.GetTags()))
	}

	st.PopulateString("ga_id", options.Config.GoogleAnalytics)

	return st.Result()
}

func ParseData(data string) string {

	st := util.HtmlTemplate("{{main}}", global.TemplateFolder).
		PopulateHTML("main", "templates/base.html").
		PopulateHTML("_header", "templates/header.html").
		PopulateHTML("_footer", "templates/footer.html").
		PopulateHTML("_ga", "templates/ga.html").
		PopulateMarkdown("_content", data)

	return st.Result()
}
