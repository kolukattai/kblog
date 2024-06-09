package boot

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/kolukattai/kblog/internal/global"
	"github.com/kolukattai/kblog/internal/models"
	"github.com/kolukattai/kblog/internal/util"
	"gopkg.in/yaml.v3"
)

func parseData(fileName string, c chan models.PageData) {
	fmt.Println(fileName)
	da, err := os.ReadFile(fmt.Sprintf("posts/%s", fileName))
	if err != nil {
		util.Error(err.Error())
	}

	fMater, _ := util.GetSplitMDData(string(da))

	frontMater := models.PageData{}
	err = yaml.Unmarshal([]byte(fMater), &frontMater)
	if err != nil {
		util.Error(err.Error())
	}
	frontMater.Slug = strings.Replace(fileName, ".md", "", 1)
	c <- frontMater
}

func InitSiteData() {
	files, err := os.ReadDir("posts")
	if err != nil {
		util.Error(err.Error())
	}

	c := make(chan models.PageData, len(files))

	var wg sync.WaitGroup
	wg.Add(len(files))

	go func() {
		for _, file := range files {
			parseData(file.Name(), c)
			wg.Done()
		}
	}()

	wg.Wait()
	close(c)

	pageDataList := []*models.PageData{}
	tags := []string{}
	categories := []string{}

	for {
		res, ok := <-c
		if !ok {
			break
		}
		pageDataList = append(pageDataList, &res)
		tags = append(tags, res.Tags...)
		categories = append(categories, res.Category)
	}

	global.Categories = util.RemoveDuplicate(categories)
	global.Tags = util.RemoveDuplicate(tags)

	sortedPageDataList := util.SortArray(pageDataList, func(left, right *models.PageData) bool {
		l, _ := time.Parse(time.RFC1123, left.Date)
		r, _ := time.Parse(time.RFC1123, right.Date)
		return l.Sub(r).Seconds() < 0
	})

	global.PageDataList.ReplaceData(sortedPageDataList)

}
