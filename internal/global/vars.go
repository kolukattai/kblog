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

	Tags       = []string{}
	Categories = []string{}
)
