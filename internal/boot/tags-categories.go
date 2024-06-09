package boot

import (
	"fmt"
	"strings"

	"github.com/kolukattai/kblog/internal/global"
	"github.com/kolukattai/kblog/internal/models"
	"github.com/kolukattai/kblog/internal/util"
)

func InitTagAndCategoryData(posts *models.PageDataList, tags, category []string, perPage int) {

	s := &models.PostPageData{
		SiteData:      map[string]*models.PageDataList{},
		SiteDataFiles: []string{},
	}
	
	ca := &models.PostPageData{
		SiteData:      map[string]*models.PageDataList{},
		SiteDataFiles: []string{},
	}

	// files := []string{}
	partition := posts.Length() / perPage
	if posts.Length()%perPage != 0 {
		partition++
	}

	for _, v := range tags {
		fileName := fmt.Sprintf("%v.json", v)
		s.SiteData[fileName] = &models.PageDataList{}
		s.SiteData[fileName].ReplaceData(posts.Filter(func(item *models.PageData) bool {
			return util.Includes(item.Tags, func(itm string, index int) bool {
				return itm == v
			})
		}))
		s.SiteDataFiles = append(s.SiteDataFiles, fileName)
	}

	for _, v := range category {
		fileName := strings.ToLower(fmt.Sprintf("ca-%v.json", v))
		ca.SiteData[fileName] = &models.PageDataList{}
		ca.SiteData[fileName].ReplaceData(posts.Filter(func(item *models.PageData) bool {
			return strings.EqualFold(item.Category, v)
		}))
		ca.SiteDataFiles = append(ca.SiteDataFiles, fileName)
	}

	global.TagPageData = s
	global.CategoryPageData = ca
}
