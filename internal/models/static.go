package models

import "encoding/json"

type JavaScript struct {
	SiteData      map[string]*PageDataList
	SiteDataFiles []string
}

func (s *JavaScript) GetSiteDataFilesJSON() []byte {
	byt, err := json.Marshal(s.SiteDataFiles)
	if err != nil {
		panic(err)
	}
	return byt
}
