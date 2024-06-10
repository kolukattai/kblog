package build

import (
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/kolukattai/kblog/internal/global"
	"github.com/kolukattai/kblog/internal/models"
	"github.com/kolukattai/kblog/internal/util"
)

type copyAssetsType string

const (
	copyAssetsTypeLocal    copyAssetsType = "local"
	copyAssetsTypeEmbedded copyAssetsType = "embedded"
)

func copyAssets(t copyAssetsType, path string, target string) {
	var files []fs.DirEntry
	if t == copyAssetsTypeEmbedded {
		fs, err := fs.ReadDir(global.StaticFiles, path)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		files = fs
	} else {
		fs, err := os.ReadDir(path)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		files = fs
	}

	for _, file := range files {
		dirPath := path + "/" + file.Name()
		if file.IsDir() {
			_ = os.MkdirAll(fmt.Sprintf("%s/%s", target, dirPath), os.ModePerm)
			copyAssets(t, dirPath, target)
		} else {
			var da []byte
			if t == copyAssetsTypeEmbedded {
				da, _ = global.StaticFiles.ReadFile(dirPath)
			} else {
				da, _ = os.ReadFile(dirPath)
			}
			if strings.Contains(dirPath, ".css") {
				da = util.Minify().Css(da)
			}
			_ = os.WriteFile(fmt.Sprintf("%s/%s", target, dirPath), da, 0666)
		}
	}

}

func createPages() {
	posts := global.PageDataList.GetData()

	for _, v := range posts {
		fmt.Println(v.Slug)
		tm := util.
			HtmlTemplate(
				global.TemplateFolder,
				models.PageTypePost,
				"_card", "_pagination", "_aside",
			).MdData(v.Slug, "", global.Config).
			MinifyResult()
		path := fmt.Sprintf("%s/post/%s", global.Config.OutputFolder, v.Slug)
		err := os.MkdirAll(path, 0755)
		if err != nil {
			fmt.Println(err.Error())
		}
		err = os.WriteFile(fmt.Sprintf("%s/index.html", path), util.Minify().
			HTML([]byte(tm)), 0644)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	// home page data
	tm := util.
		HtmlTemplate(
			global.TemplateFolder,
			models.PageTypeHome,
			"_card", "_pagination", "_aside",
		).MdData(
		"",
		global.PostPageData.SiteData["0.json"].GetData(),
		global.Config,
		models.PageData{
			Title:       global.Config.Default.Title,
			Description: global.Config.Default.Description,
			Keywords:    global.Config.Default.Keywords,
		},
	).
		MinifyResult()

	_ = os.WriteFile(global.Config.OutputFolder+"/index.html", []byte(tm), 0755)

	f1 := global.PostPageData.SiteDataFiles

	_ = os.MkdirAll(global.Config.OutputFolder+"/data", 0755)

	_ = os.WriteFile(global.Config.OutputFolder+"/data/data-map.json", global.PostPageData.GetSiteDataFilesJSON(), 0755)

	for _, v := range f1 {
		fn := fmt.Sprintf("%s/data/%s", global.Config.OutputFolder, v)
		da, ok := global.PostPageData.SiteData[v]
		fmt.Println(fn, ok, da)
		if !ok {
			continue
		}
		_ = os.WriteFile(fn, []byte(da.GetJSON()), 0755)
	}

	// tags page data

	f2 := global.TagPageData.SiteDataFiles

	for _, v := range f2 {
		fmt.Println(v)
		fn := fmt.Sprintf("%s/tag/%s/", global.Config.OutputFolder, strings.Replace(v, ".json", "", 1))
		dat, ok := global.TagPageData.SiteData[v]
		if !ok {
			continue
		}

		tm := util.HtmlTemplate(global.TemplateFolder, models.PageTypeHome, "_card", "_pagination", "_aside").
			MdData("",
				dat.GetData(),
				global.Config,
			).
			MinifyResult()
		_ = os.MkdirAll(fn, 0755)
		_ = os.WriteFile(fmt.Sprintf("%s/index.html", fn), []byte(tm), 0755)
	}

	// categories

	f3 := global.CategoryPageData.SiteDataFiles

	fmt.Println(f3)

	for _, v := range f3 {
		fn := fmt.Sprintf("%s/category/%s/", global.Config.OutputFolder, strings.Replace(strings.Replace(v, ".json", "", 1), "ca-", "", 1))
		dat, ok := global.CategoryPageData.SiteData[strings.ToLower(v)]
		if !ok {
			continue
		}

		tm := util.HtmlTemplate(global.TemplateFolder, models.PageTypeHome, "_card", "_pagination", "_aside").
			MdData("",
				dat.GetData(),
				global.Config,
			).
			MinifyResult()
		_ = os.MkdirAll(fn, 0755)
		_ = os.WriteFile(fmt.Sprintf("%s/index.html", fn), []byte(tm), 0755)
	}
}

func Exec() {
	os.RemoveAll(global.Config.OutputFolder)

	copyAssets(copyAssetsTypeEmbedded, "static", global.Config.OutputFolder)
	copyAssets(copyAssetsTypeLocal, "public", global.Config.OutputFolder)
	createPages()
}
