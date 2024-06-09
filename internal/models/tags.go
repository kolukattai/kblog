package models

import "encoding/json"

type TagPageData struct {
	SiteData      map[string]*PageDataList
	SiteDataFiles []string
}

func (st *TagPageData) GetOneSiteData(key string) *PageDataList {
	data, ok := st.SiteData[key]
	if !ok {
		return nil
	}
	return data
}

func (s *TagPageData) GetSiteDataFilesJSON() []byte {
	byt, err := json.Marshal(s.SiteDataFiles)
	if err != nil {
		panic(err)
	}
	return byt
}
