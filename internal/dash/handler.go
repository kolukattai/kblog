package dash

import (
	"net/http"

	"github.com/kolukattai/kblog/internal/global"
)

type handler struct{}

func Handler() *handler {
	return &handler{}
}

func (st handler) Home() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cors(r, w)
		w.WriteHeader(200)
		w.Write([]byte(global.PageDataList.GetJSON()))
	})
}

func (st handler) GetOnePost() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cors(r, w)

		slug := r.PathValue("slug")

		data := global.PageDataList.GetOneBySlug(slug)

		if data == nil {
			w.WriteHeader(404)
			w.Write([]byte("Not Found"))
			return
		}

		w.WriteHeader(200)
		w.Write([]byte(data.JSON()))
	})
}
