package server

import (
	"net/http"

	"github.com/kolukattai/kblog/internal/global"
	"github.com/kolukattai/kblog/internal/util"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tm := util.HtmlTemplate(global.TemplateFolder, "home", "_card").
		MdData("",
			global.PageDataList.GetData(),
		).
		Result()
	w.WriteHeader(200)
	w.Write([]byte(tm))
}

func postHandler(w http.ResponseWriter, r *http.Request) {

}

func tagsHandler(w http.ResponseWriter, r *http.Request) {

}

func categoriesHandler(w http.ResponseWriter, r *http.Request) {

}
