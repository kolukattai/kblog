package models

import (
	"encoding/json"
	"strings"
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

func (s *PageDataList) Add(data *PageData) {
	s.data = append(s.data, data)
}

func (s *PageDataList) Get(index int) *PageData {
	if (len(s.data) - 1) < index {
		return nil
	}
	return s.data[index]
}

func (s *PageDataList) GetJSON() string {
	bData, err := json.Marshal(s.data)
	if err != nil {
		panic(err)
	}
	return string(bData)
}

type PageData struct {
	Title        string `yaml:"title" json:"title"`
	Description  string `yaml:"description" json:"description"`
	Keywords     string `yaml:"keywords" json:"keywords"`
	Tags         string `yaml:"tags" json:"tags"`
	Category     string `yaml:"category" json:"category"`
	Author       string `yaml:"author" json:"author"`
	LandingImage string `yaml:"landingImage" json:"landingImage"`
	Slug         string `json:"slug,omitempty"`
}

func (st *PageData) GetTags() []string {
	return strings.Split(strings.ReplaceAll(st.Tags, " ", ""), ",")
}
