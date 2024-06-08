package models

type MDPageData struct {
	Content         string
	MetaData        PageData
	Data            any
	DefaultMetaData *Config
	PageType        PageType
	Global          struct {
		Tags       []string
		Catagories []string
	}
}
