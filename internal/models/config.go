package models

type Config struct {
	Version      string `yaml:"version"`
	Name         string `yaml:"name"`
	Logo         string `yaml:"logo"`
	OutputFolder string `yaml:"outputFolder"`
	DomainName   string `yaml:"domainName"`
	LandingImage string `yaml:"landingImage"`
	Default      struct {
		Title       string `yaml:"title"`
		Description string `yaml:"description"`
		Keywords    string `yaml:"keywords"`
		Author      string `yaml:"author"`
	} `yaml:"default"`
	GoogleAnalytics  string   `yaml:"ga"`
	GoogleTagManager string   `yaml:"gtm"`
	Styles           []string `yaml:"styles"`
	Scripts          []string `yaml:"scripts"`
	PerPage          int      `yaml:"perPage"`
	Twitter          string   `yaml:"twitter"`
	Facebook         string   `yaml:"facebook"`
	Instagram        string   `yaml:"instagram"`
}
