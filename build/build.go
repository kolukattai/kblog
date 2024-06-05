package build

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/kolukattai/kblog/global"
	"github.com/kolukattai/kblog/models"
	"github.com/kolukattai/kblog/parser"
	"github.com/kolukattai/kblog/util"
)

type pageData struct {
	Name     string           `json:"name"`
	Data     string           `json:"data"`
	Slug     string           `json:"slug"`
	MetaData *models.PageData `json:"metaData"`
}

func Parse(name string, c chan pageData, wg *sync.WaitGroup) {
	defer wg.Done()
	data, metaData := parser.Parse(
		name, parser.Options{
			LandingImage: true,
			Tags:         true,
			Author:       true,
			Config:       global.Config,
		},
	)
	c <- pageData{
		Name:     name,
		Data:     data,
		Slug:     strings.ReplaceAll(name, " ", "-"),
		MetaData: metaData,
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

	tags := []string{}
	categoryes := []string{}
	siteData := []pageData{}

	wg.Add(len(dir))
	for {
		res, ok := <-files
		if !ok {
			break
		}
		tags = append(tags, strings.Split(res.MetaData.Tags, ",")...)
		categoryes = append(categoryes, res.MetaData.Category)
		sd := res
		sd.Data = ""
		siteData = append(siteData, sd)
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

	tags = util.RemoveDuplicate(tags)
	categoryes = util.RemoveDuplicate(categoryes)

	siteDataID := fmt.Sprintf("%v-site.json", time.Now().Nanosecond())

	siteDataByte, _ := json.Marshal(siteData)

	_ = os.MkdirAll("dist/data", os.ModePerm)

	_ = os.WriteFile(fmt.Sprintf("dist/data/%v", siteDataID), siteDataByte, 0666)

	indexPageData := parser.ParseData("")
	if err = os.WriteFile("dist/index.html", []byte(indexPageData), 0666); err != nil {
		panic(err)
	}
}

func searchData(data []pageData) {
	
}

func copyAssets(path string) {
	files, err := fs.ReadDir(global.StaticFiles, path)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, file := range files {
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
	t := time.Now()
	generatePages()
	copyAssets("static")
	d := time.Since(t)
	fmt.Printf("build completed in %vs\n", d.Seconds())
}
