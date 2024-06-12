package boot

import (
	"fmt"
	"sync"

	"github.com/kolukattai/kblog/internal/global"
	"github.com/kolukattai/kblog/internal/models"
)

func InitPostData(posts *models.PageDataList, perPage int) {
	prefix := ""
	siteData := "%v%v.json"

	s := &models.PostPageData{
		SiteData: map[string]*models.PageDataList{},
		SiteDataFiles: []string{},
	}

	// files := []string{}
	partition := posts.Length() / perPage
	if posts.Length()%perPage != 0 {
		partition++
	}

	fmt.Println(partition)

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

	global.PostPageData = s
}
