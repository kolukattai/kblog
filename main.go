/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"embed"
	"os"

	"github.com/kolukattai/kblog/cmd"
	"github.com/kolukattai/kblog/internal/global"
	"github.com/kolukattai/kblog/internal/models"
	"gopkg.in/yaml.v3"
)

//go:embed all:templates/*
var templateFolder embed.FS

//go:embed static/*
var staticFiles embed.FS

func init() {
	global.TemplateFolder = templateFolder
	global.StaticFiles = staticFiles

	os.MkdirAll("posts", os.ModePerm)

	confFile, err := os.ReadFile("config.yaml")
	if err != nil {
		byt, _ := yaml.Marshal(&models.Config{
			PerPage:    10,
			Instagram:  "https://instagram.com/mrboxopener",
			Facebook:   "https://facebook.com/mrboxopener",
			Twitter:    "https://twitter.com/mrboxopener",
			DomainName: "domain.com",
		})
		_ = os.WriteFile("config.yaml", byt, 0666)
		confFile = byt
	}
	var conf models.Config
	if err := yaml.Unmarshal(confFile, &conf); err != nil {
		panic(err)
	}
	if len(conf.DomainName) == 0 {
		conf.DomainName = "domain.com"
	}
	if len(conf.OutputFolder) == 0 {
		conf.OutputFolder = "dist"
	}
	global.Config = &conf
}

func main() {
	cmd.Execute()
}
