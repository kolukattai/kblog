package build

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/kolukattai/kblog/global"
	"github.com/kolukattai/kblog/parser"
)

type pageData struct {
	Name string
	Data string
}

func Parse(name string, c chan pageData, wg *sync.WaitGroup) {
	defer wg.Done()
	c <- pageData{
		Name: name,
		Data: parser.Parse(
			name, parser.Options{
				LandingImage: true,
				Tags:         true,
				Author:       true,
				Config:       global.Config,
			},
		),
	}
}

func generatePages() {
	dir, err := os.ReadDir("posts")
	if err != nil {
		panic(err)
	}

	files := make(chan pageData, len(dir))

	var wg sync.WaitGroup
	wg.Add(len(dir))

	fmt.Println("parsing file...")
	for _, file := range dir {
		go Parse(strings.Replace(file.Name(), ".md", "", 1), files, &wg)
	}
	wg.Wait()
	close(files)
	fmt.Println("parsing completed...")

	wg.Add(len(dir))
	for {
		res, ok := <-files
		fmt.Println("OK", ok)
		if !ok {
			break
		}
		go func(f pageData, wg *sync.WaitGroup) {
			defer wg.Done()
			err := os.MkdirAll(fmt.Sprintf("dist/posts/%s", f.Name), os.ModePerm)
			if err != nil {
				log.Fatal(err.Error())
			}
			err = os.WriteFile(fmt.Sprintf("dist/posts/%s/index.html", f.Name), []byte(f.Data), 0666)
			if err != nil {
				log.Fatal(err.Error())
			}
		}(res, &wg)
	}

	wg.Wait()
}

func copyAssets(path string) {
	files, err := fs.ReadDir(global.StaticFiles, path)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		dirPath := path + "/" + file.Name()
		if file.IsDir() {
			_ = os.MkdirAll("dist/"+dirPath, os.ModePerm)
			copyAssets(dirPath)
		} else {
			val, _ := global.StaticFiles.ReadFile(dirPath)
			_ = os.WriteFile("dist/"+dirPath, val, 0666)
		}
	}

}

func Exec() {
	generatePages()
	copyAssets("static")
	fmt.Println("build completed...")
}
