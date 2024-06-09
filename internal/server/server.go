package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kolukattai/kblog/internal/global"
)

func Run(port string) {
	var staticFS = http.FS(global.StaticFiles)
	fs := http.FileServer(staticFS)

	images := http.FileServer(http.Dir("public"))

	http.Handle("GET /public/", http.StripPrefix("/public/", images))
	http.Handle("GET /static/", fs)

	http.Handle("GET /", http.HandlerFunc(homeHandler))

	for _, v := range global.PostPageData.SiteDataFiles {
		path := fmt.Sprintf("GET /data/%v", v)
		http.Handle(path, http.HandlerFunc(handleDataFile))
	}

	for _, v := range global.TagPageData.SiteDataFiles {
		path := fmt.Sprintf("GET /data/%v", v)
		http.Handle(path, http.HandlerFunc(handleTagDataFile))
	}
	http.Handle("GET /data/data-map.json", http.HandlerFunc(handleDataMap))

	http.Handle("GET /post/{slug}", http.HandlerFunc(postHandler))
	http.Handle("GET /tags", http.HandlerFunc(tagsHandler))
	http.Handle("GET /tag/{tag}", http.HandlerFunc(tagsHandler))
	http.Handle("GET /category/{category}", http.HandlerFunc(categoryHandler))

	p := fmt.Sprintf(":%s", port)
	fmt.Printf("sight stated at http://localhost%s\n", p)
	log.Fatal(http.ListenAndServe(p, nil))
}
