package components

import (
	"fmt"

	"github.com/kolukattai/kblog/global"
	"github.com/kolukattai/kblog/models"
	"github.com/kolukattai/kblog/util"
)

func Card(data models.PageData) models.Component {
	cardByte, err := global.TemplateFolder.ReadFile("templates/card.html")
	if err != nil {
		panic(err)
	}
	st := util.HtmlTemplate(string(cardByte), global.TemplateFolder)
	util.IterateStruct(data, func(key string, value any) {
		st.PopulateString(key, fmt.Sprintf("%v", value))
		st.PopulateComponent(key, models.Component(fmt.Sprintf("%v", value)))
	})
	st.PopulateString("Slug", data.Slug)
	st.PopulateString("LandingImage", data.LandingImage)
	st.PopulateString("Title", data.Title)
	st.PopulateComponent("_tags", Tags(data.GetTags()))
	return models.Component(st.Result())
}
