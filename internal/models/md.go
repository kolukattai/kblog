package models

import "html/template"

type MDPageData struct {
	Content         template.HTML
	MetaData        PageData
	Data            any
	DefaultMetaData *Config
	PageType        PageType
	Year            string
	Month           string
	Date            string
	Global          struct {
		Tags       []string
		Catagories []string
	}
}
