package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/kolukattai/kblog/internal/global"
	"github.com/kolukattai/kblog/internal/models"
	"github.com/kolukattai/kblog/internal/util"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tm := util.HtmlTemplate(global.TemplateFolder, models.PageTypeHome, "_card", "_pagination", "_aside").
		MdData("",
			global.PostPageData.SiteData["0.json"].GetData(),
			global.Config,
		).
		Result()
	w.WriteHeader(200)
	w.Write([]byte(tm))
}

func handleDataFile(w http.ResponseWriter, r *http.Request) {
	fileName := strings.Replace(r.URL.Path, "/data/", "", 1)
	data := global.PostPageData.SiteData[fileName].GetJSON()
	w.WriteHeader(200)
	w.Write([]byte(data))
}

func handleTagDataFile(w http.ResponseWriter, r *http.Request) {
	fileName := strings.Replace(r.URL.Path, "/data/", "", 1)
	data := global.TagPageData.SiteData[fileName].GetJSON()
	w.WriteHeader(200)
	w.Write([]byte(data))
}

func handleDataMap(w http.ResponseWriter, r *http.Request) {
	data := global.PostPageData.GetSiteDataFilesJSON()
	w.WriteHeader(200)
	w.Write(data)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	tm := util.HtmlTemplate(global.TemplateFolder, models.PageTypePost, "_card", "_pagination", "_aside").
		MdData(slug,
			"",
			global.Config,
		).
		Result()
	w.WriteHeader(200)
	w.Write([]byte(tm))
}

func tagsHandler(w http.ResponseWriter, r *http.Request) {
	tag := fmt.Sprintf("%v.json", r.PathValue("tag"))

	dat, ok := global.TagPageData.SiteData[tag]

	if !ok {
		w.WriteHeader(200)
		w.Write([]byte("not found"))
	}

	tm := util.HtmlTemplate(global.TemplateFolder, models.PageTypeHome, "_card", "_pagination", "_aside").
		MdData("",
			dat.GetData(),
			global.Config,
			models.PageData{
				Title: "#" + strings.Replace(strings.ToUpper(tag), ".JSON", "", 1),
			},
		).
		Result()
	w.WriteHeader(200)
	w.Write([]byte(tm))

}

func categoryHandler(w http.ResponseWriter, r *http.Request) {
	tag := fmt.Sprintf("ca-%v.json", strings.ToLower(r.PathValue("category")))

	dat, ok := global.CategoryPageData.SiteData[tag]

	if !ok {
		w.WriteHeader(200)
		w.Write([]byte("not found"))
	}

	tm := util.HtmlTemplate(global.TemplateFolder, models.PageTypeHome, "_card", "_pagination", "_aside").
		MdData("",
			dat.GetData(),
			global.Config,
			models.PageData{
				Title: strings.Replace(strings.Replace(strings.ToUpper(tag), "CA-", "", 1), ".JSON", "", 1),
			},
		).
		Result()
	w.WriteHeader(200)
	w.Write([]byte(tm))
}
