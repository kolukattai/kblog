package models

import (
	"encoding/json"
)

type PageType string

const (
	PageTypePost       PageType = "post"
	PageTypeTag        PageType = "_tag"
	PageTypeTags       PageType = "_tags"
	PageTypeCategory   PageType = "_category"
	PageTypeCategories PageType = "_categories"
	PageTypeHome       PageType = "home"
)

type SlimPageData struct {
	LandingImage string `json:"landingImage"`
	Title        string `json:"title"`
	Category     string `yaml:"category" json:"category"`
	Tags         string `json:"tags"`
	Slug         string `json:"slug"`
}

type PageDataList struct {
	data []*PageData
}

func (s *PageDataList) Length() int {
	return len(s.data)
}

func (s *PageDataList) ForEach(yield func(index int, data *PageData)) {
	for i, v := range s.data {
		yield(i, v)
	}
}

func (s *PageDataList) ReplaceData(data []*PageData) {
	s.data = data
}

func (s *PageDataList) Add(data *PageData) {
	s.data = append(s.data, data)
}

func (s *PageDataList) GetData() []*PageData {
	return s.data
}

func (s *PageDataList) Get(index int) *PageData {
	if (len(s.data) - 1) < index {
		return nil
	}
	return s.data[index]
}

func (s *PageDataList) GetOneBySlug(slug string) *PageData {
	for _, v := range s.data {
		if v.Slug == slug {
			return v
		}
	}
	return nil
}

func (s *PageDataList) GetJSON() string {
	bData, err := json.Marshal(s.data)
	if err != nil {
		panic(err)
	}
	return string(bData)
}

func (s *PageDataList) Filter(yield func(item *PageData) bool) []*PageData {
	items := []*PageData{}
	for _, v := range s.data {
		if yield(v) {
			items = append(items, v)
		}
	}
	return items
}

type PageData struct {
	Title        string   `yaml:"title" json:"title"`
	Description  string   `yaml:"description" json:"description"`
	Keywords     string   `yaml:"keywords" json:"keywords"`
	Tags         []string `yaml:"tags" json:"tags"`
	Category     string   `yaml:"category" json:"category"`
	Author       string   `yaml:"author" json:"author"`
	LandingImage string   `yaml:"landingImage" json:"landingImage"`
	Date         string   `yaml:"date" json:"date"`
	Slug         string   `yaml:"-" json:"slug,omitempty"`
}

func (st *PageData) JSON() string {
	byt, _ := json.Marshal(st)
	return string(byt)
}
