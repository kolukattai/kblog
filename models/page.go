package models

import "strings"

type PageData struct {
	Title        string `yaml:"title"`
	Description  string `yaml:"description"`
	Keywords     string `yaml:"keywords"`
	Tags         string `yaml:"tags"`
	Category     string `yaml:"category"`
	Author       string `yaml:"author"`
	LandingImage string `yaml:"landingImage"`
}

func (st *PageData) GetTags() []string {
	return strings.Split(strings.ReplaceAll(st.Tags, " ", ""), ",")
}