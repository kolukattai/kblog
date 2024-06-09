package global

import (
	"embed"

	"github.com/kolukattai/kblog/internal/models"
)

var (
	TemplateFolder embed.FS
	StaticFiles    embed.FS

	// config
	Config *models.Config

	// runtime data
	PageDataList = &models.PageDataList{}

	// static file paths
	JavaScriptLocation = &models.JavaScript{}

	PostPageData = &models.PostPageData{}
	
	TagPageData = &models.PostPageData{}
	
	CategoryPageData = &models.PostPageData{}

	Tags       = []string{}
	Categories = []string{}
)
