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

	util.IterationConcurrently(files, func(file fs.DirEntry, _ int) {
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
			if strings.Contains(dirPath, ".js") {
				if len(global.Config.GoogleAnalytics) != 0 && file.Name() == "ga.js" {
					da = []byte(strings.Replace(string(da), "{{ga_id}}", global.Config.GoogleAnalytics, 1))
				}
			}
			_ = os.WriteFile(fmt.Sprintf("%s/%s", target, dirPath), da, 0666)
		}
	})

}

func createPages() {
	posts := global.PageDataList.GetData()

	createPost := func(v *models.PageData) {
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

	util.IterationConcurrently(posts, func(v *models.PageData, index int) {
		createPost(v)
	})

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

	_ = os.MkdirAll(global.Config.OutputFolder+"/data", 0755)

	_ = os.WriteFile(global.Config.OutputFolder+"/data/data-map.json", global.PostPageData.GetSiteDataFilesJSON(), 0755)

	util.IterationConcurrently(global.PostPageData.SiteDataFiles, func(v string, _ int) {
		fn := fmt.Sprintf("%s/data/%s", global.Config.OutputFolder, v)
		da, ok := global.PostPageData.SiteData[v]
		if !ok {
			return
		}
		_ = os.WriteFile(fn, []byte(da.GetJSON()), 0755)
	})

	// tags page data
	util.IterationConcurrently(global.TagPageData.SiteDataFiles, func(v string, _ int) {
		fn := fmt.Sprintf("%s/tag/%s/", global.Config.OutputFolder, strings.Replace(v, ".json", "", 1))
		dat, ok := global.TagPageData.SiteData[v]
		if !ok {
			return
		}

		tm := util.HtmlTemplate(global.TemplateFolder, models.PageTypeHome, "_card", "_pagination", "_aside").
			MdData("",
				dat.GetData(),
				global.Config,
			).
			MinifyResult()
		_ = os.MkdirAll(fn, 0755)
		_ = os.WriteFile(fmt.Sprintf("%s/index.html", fn), []byte(tm), 0755)
	})

	// categories
	util.IterationConcurrently(global.CategoryPageData.SiteDataFiles, func(v string, _ int) {
		fn := fmt.Sprintf("%s/category/%s/", global.Config.OutputFolder, strings.Replace(strings.Replace(v, ".json", "", 1), "ca-", "", 1))
		dat, ok := global.CategoryPageData.SiteData[strings.ToLower(v)]
		if !ok {
			return
		}

		tm := util.HtmlTemplate(global.TemplateFolder, models.PageTypeHome, "_card", "_pagination", "_aside").
			MdData("",
				dat.GetData(),
				global.Config,
			).
			MinifyResult()
		_ = os.MkdirAll(fn, 0755)
		_ = os.WriteFile(fmt.Sprintf("%s/index.html", fn), []byte(tm), 0755)
	})
}

func generateSitemap(domain string, data []*models.PageData) {

	list := ""
	for _, v := range data {
		list += fmt.Sprintf("<url><loc>https://%s/post/%s/</loc><lastmod>%s</lastmod></url>",
			domain, v.Slug, v.Date,
		)
	}
	templ := fmt.Sprintf(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:news="http://www.google.com/schemas/sitemap-news/0.9" xmlns:xhtml="http://www.w3.org/1999/xhtml" xmlns:mobile="http://www.google.com/schemas/sitemap-mobile/1.0" xmlns:image="http://www.google.com/schemas/sitemap-image/1.1" xmlns:video="http://www.google.com/schemas/sitemap-video/1.1">
	%s
	</urlset>`, list)

	_ =os.WriteFile(fmt.Sprintf("%s/posts.xml", global.Config.OutputFolder), []byte(templ), 0755)
}

func Exec() {
	os.RemoveAll(global.Config.OutputFolder)

	copyAssets(copyAssetsTypeEmbedded, "static", global.Config.OutputFolder)
	copyAssets(copyAssetsTypeLocal, "public", global.Config.OutputFolder)
	createPages()
	util.GenerateRobtTxt(global.Config.DomainName, global.Config.OutputFolder)
	generateSitemap(global.Config.DomainName, global.PageDataList.GetData())
}
