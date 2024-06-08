package server

import (
	"net/http"
	"strings"

	"github.com/kolukattai/kblog/internal/global"
	"github.com/kolukattai/kblog/internal/models"
	"github.com/kolukattai/kblog/internal/util"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	

	tm := util.HtmlTemplate(global.TemplateFolder, models.PageTypeHome, "_card", "_pagination", "_aside").
		MdData("",
			global.PageDataList.GetData(),
			global.Config,
		).
		Result()
	w.WriteHeader(200)
	w.Write([]byte(tm))
}

func handleDataFile(w http.ResponseWriter, r *http.Request) {
	fileName := strings.Replace(r.URL.Path, "/data/", "", 1)
	data := global.JavaScriptLocation.SiteData[fileName].GetJSON()
	w.WriteHeader(200)
	w.Write([]byte(data))
}

func handleDataMap(w http.ResponseWriter, r *http.Request) {
	data := global.JavaScriptLocation.GetSiteDataFilesJSON()
	w.WriteHeader(200)
	w.Write(data)
}

func postHandler(w http.ResponseWriter, r *http.Request) {

}

func tagsHandler(w http.ResponseWriter, r *http.Request) {

}

func categoriesHandler(w http.ResponseWriter, r *http.Request) {

}
