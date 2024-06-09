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

	js := http.FileServer(http.Dir("js"))
	css := http.FileServer(http.Dir("css"))
	images := http.FileServer(http.Dir("images"))

	http.Handle("GET /css/", http.StripPrefix("/css/", css))
	http.Handle("GET /js/", http.StripPrefix("/js/", js))
	http.Handle("GET /images/", http.StripPrefix("/images/", images))
	http.Handle("GET /static/", fs)

	http.Handle("GET /", http.HandlerFunc(homeHandler))

	for _, v := range global.PostPageData.SiteDataFiles {
		path := fmt.Sprintf("GET /data/%v", v)
		http.Handle(path, http.HandlerFunc(handleDataFile))
	}

	fmt.Println(global.TagPageData.SiteDataFiles)
	for _, v := range global.TagPageData.SiteDataFiles {
		path := fmt.Sprintf("GET /data/%v", v)
		fmt.Println(path)
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
