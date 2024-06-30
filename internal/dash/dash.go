package dash

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kolukattai/kblog/internal/boot"
	"github.com/kolukattai/kblog/internal/global"
)

func Run(port string) {
	boot.InitSiteData()

	p := fmt.Sprintf(":%s", port)

	var staticFS = http.FS(global.StaticFiles)
	fs := http.FileServer(staticFS)

	images := http.FileServer(http.Dir("public"))

	http.Handle("GET /public/", http.StripPrefix("/public/", images))
	http.Handle("GET /static/", fs)

	http.Handle("GET /api/posts", Handler().Home())
	http.Handle("GET /api/posts/{slug}", Handler().GetOnePost())

	log.Fatal(http.ListenAndServe(p, nil))
}
