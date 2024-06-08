package boot

import (
	"fmt"
	"sync"
	"time"

	"github.com/kolukattai/kblog/internal/global"
	"github.com/kolukattai/kblog/internal/models"
)

func splitPage(prefix, template string, index int, posts *models.PageDataList, c chan *models.KPageData, perPage int) {
	file := fmt.Sprintf(template, prefix, index)

	data := []*models.PageData{}

	count := index * perPage

	for i := 0; i < perPage; i++ {
		idx := i + count
		da := posts.Get(idx)
		if da == nil {
			break
		}
		data = append(data, da)
	}

	c <- &models.KPageData{Key: file, Value: data}

}

func InitJavascriptMaps(posts *models.PageDataList, perPage int) {
	prefix := fmt.Sprintf("%v", time.Now().Nanosecond())
	siteData := "%v-site-data-%v.json"

	s := &models.JavaScript{
		SiteData: map[string]*models.PageDataList{},
	}

	// files := []string{}
	partition := posts.Length() / perPage
	if posts.Length()%perPage != 0 {
		partition++
	}

	var wg sync.WaitGroup
	wg.Add(1)
	c := make(chan *models.KPageData, partition)

	go func() {
		for i := 0; i < partition; i++ {
			splitPage(prefix, siteData, i, posts, c, perPage)
		}
		wg.Done()
	}()

	wg.Wait()
	close(c)

	for {
		file, ok := <-c
		if !ok {
			break
		}
		d := &models.PageDataList{}
		d.ReplaceData(file.Value)
		s.SiteData[file.Key] = d
		s.SiteDataFiles = append(s.SiteDataFiles, file.Key)
	}

	global.JavaScriptLocation = s
}
