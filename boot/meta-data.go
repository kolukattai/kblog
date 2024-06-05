package boot

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/kolukattai/kblog/global"
	"github.com/kolukattai/kblog/models"
	"github.com/kolukattai/kblog/parser"
)

func parsePage(name string, c chan models.PageData) {
	val, err := os.ReadFile(fmt.Sprintf("posts/%s", name))
	if err != nil {
		panic(err)
	}
	d := parser.ParsePageMetaData(string(val))
	d.Slug = strings.Split(name, ".")[0]
	c <- d
}

func InitMetaData() {
	postsDir, err := os.ReadDir("posts")
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(postsDir))

	pages := make(chan models.PageData, len(postsDir))

	for _, post := range postsDir {
		go func(name string) {
			defer wg.Done()
			parsePage(name, pages)
		}(post.Name())
	}

	wg.Wait()
	close(pages)

	for {
		page, ok := <-pages
		if !ok {
			break
		}
		global.PageDataList.Add(&page)
	}

}
