package models

import "encoding/json"

type PostPageData struct {
	SiteData      map[string]*PageDataList
	SiteDataFiles []string
}

func (st *PostPageData) GetOneSiteData(key string) *PageDataList {
	data, ok := st.SiteData[key]
	if !ok {
		return nil
	}
	return data
}

func (s *PostPageData) GetSiteDataFilesJSON() []byte {
	byt, err := json.Marshal(s.SiteDataFiles)
	if err != nil {
		panic(err)
	}
	return byt
}
