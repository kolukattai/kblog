package server

import (
	"net/http"
	"strings"

	"github.com/kolukattai/kblog/components"
	"github.com/kolukattai/kblog/global"
	"github.com/kolukattai/kblog/models"
	"github.com/kolukattai/kblog/parser"
	"github.com/kolukattai/kblog/util"
)

func handlePosts() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file := r.PathValue("file")
		val, _ := parser.Parse(strings.Replace(file, "posts/", "", 1), parser.Options{
			LandingImage: true,
			Tags:         true,
			Author:       true,
			Config:       global.Config,
		})
		w.WriteHeader(200)
		w.Write([]byte(val))
	})
}

func handleHome() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := parser.ParseData("")

		st := util.HtmlTemplate(data, global.TemplateFolder)
		cards := []models.Component{}
		global.PageDataList.ForEach(func(index int, data *models.PageData) {
			cards = append(cards, components.Card(*data))
		})

		result := st.
			PopulateHTML("_content", "templates/home.html").
			PopulateComponents("cards", cards).
			Result()

		w.WriteHeader(200)
		w.Write([]byte(result))
	})
}
