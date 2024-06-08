package models

type Config struct {
	Version float64 `yaml:"version"`
	Default struct {
		Title       string `yaml:"title"`
		Description string `yaml:"description"`
		Keywords    string `yaml:"keywords"`
		Author      string `yaml:"author"`
	} `yaml:"default"`
	GoogleAnalytics string   `yaml:"ga"`
	Styles          []string `yaml:"styles"`
	Scripts         []string `yaml:"scripts"`
	PerPage         int      `yaml:"perPage"`
}
