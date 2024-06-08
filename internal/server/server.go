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

	http.Handle("GET /static/", fs)
	http.Handle("GET /css/", css)
	http.Handle("GET /js/", js)

	http.Handle("GET /", http.HandlerFunc(homeHandler))

	// for _, v := range global.JavaScriptLocation.SiteDataFiles {
	// 	http.Handle(fmt.Sprintf("GET /data/{file}"), http.HandlerFunc(homeHandler))
	// }
	http.Handle("GET /post/{slug}", http.HandlerFunc(postHandler))
	http.Handle("GET /tags", http.HandlerFunc(tagsHandler))
	http.Handle("GET /tag/{tag}", http.HandlerFunc(tagsHandler))
	http.Handle("GET /category/{category}", http.HandlerFunc(categoriesHandler))

	p := fmt.Sprintf(":%s", port)
	fmt.Printf("sight stated at http://localhost%s", p)
	log.Fatal(http.ListenAndServe(p, nil))
}
