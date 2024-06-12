package cmdadd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kolukattai/kblog/internal/models"
	"github.com/kolukattai/kblog/internal/util"
	"gopkg.in/yaml.v3"
)

func Create(title string) {
	location := fmt.Sprintf("posts/%v.md", strings.ReplaceAll(strings.ToLower(title), " ", "-"))

	_, err := os.Stat(location)
	if err == nil {
		util.Error("post with same name already exists")
	}

	cTitle := fmt.Sprintf("%s%s", strings.ToUpper(title[0:1]), title[1:])
	pageData := &models.PageData{
		Title:        cTitle,
		Description:  "this is post description",
		Keywords:     "one, two, three",
		Tags:         []string{"one", "two", "three"},
		Category:     "technology",
		Author:       "<your name>",
		LandingImage: "<image location>",
		Date:         time.Now().Format(time.RFC1123),
		Slug:         location,
	}

	frontMatter, err := yaml.Marshal(pageData)
	if err != nil {
		util.Error(err.Error())
	}

	payload := fmt.Sprintf("---\n%s---\n\nthis is page content", string(frontMatter))

	err = os.WriteFile(location, []byte(payload), 0666)

	if err != nil {
		util.Error(err.Error())
	}

	fmt.Printf("Post created in %v", location)
}
