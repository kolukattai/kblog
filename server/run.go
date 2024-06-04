package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kolukattai/kblog/global"
)

func Run(port string) {
	var staticFS = http.FS(global.StaticFiles)
	fs := http.FileServer(staticFS)

	http.Handle("/static/", fs)

	http.Handle("GET /posts/{file}", handlePosts())

	fmt.Printf("application stated at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
