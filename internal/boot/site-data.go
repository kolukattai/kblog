package boot

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/kolukattai/kblog/internal/global"
	"github.com/kolukattai/kblog/internal/models"
	"github.com/kolukattai/kblog/internal/util"
	"gopkg.in/yaml.v3"
)

func parseData(fileName string, c chan models.PageData) {
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

	for {
		res, ok := <-c
		if !ok {
			break
		}
		pageDataList = append(pageDataList, &res)
	}

	sortedPageDataList := util.SortArray(pageDataList, func(left, right *models.PageData) bool {
		l, _ := time.Parse(time.RFC1123, left.Date)
		r, _ := time.Parse(time.RFC1123, right.Date)
		return l.Sub(r).Seconds() < 0
	})

	global.PageDataList.ReplaceData(sortedPageDataList)

}
