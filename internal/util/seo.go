package util

import (
	"fmt"
	"os"
)

func GenerateRobtTxt(domain, distFolder string) {
	d := fmt.Sprintf(`# *
User-agent: *
Allow: /	

# Sitemaps
Sitemap: %s/posts.xml`,
		domain,
	)

	_ = os.WriteFile(fmt.Sprintf("%s/robots.txt", distFolder), []byte(d), 0666)
}
