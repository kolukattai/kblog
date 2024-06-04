package server

import (
	"net/http"
	"strings"

	"github.com/kolukattai/kblog/global"
	"github.com/kolukattai/kblog/parser"
)

func handlePosts() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file := r.PathValue("file")
		val := parser.Parse(strings.Replace(file, "posts/", "", 1), parser.Options{
			LandingImage: true,
			Tags:         true,
			Author:       true,
			Config:       global.Config,
		})
		w.WriteHeader(200)
		w.Write([]byte(val))
	})
}
