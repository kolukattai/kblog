package models

import "html/template"

type MDPageData struct {
	Content         template.HTML
	MetaData        PageData
	Data            any
	DefaultMetaData *Config
	PageType        PageType
	Global          struct {
		Tags       []string
		Catagories []string
	}
}
