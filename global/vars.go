package global

import (
	"embed"

	"github.com/kolukattai/kblog/models"
)

var (
	TemplateFolder embed.FS
	StaticFiles    embed.FS

	// config
	Config *models.Config

	// runtime data
	PageDataList = &models.PageDataList{}
)
