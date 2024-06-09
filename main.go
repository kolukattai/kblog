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

	confFile, err := os.ReadFile("config.yaml")
	if err != nil {
		byt, _ := yaml.Marshal(&models.Config{PerPage: 10})
		_ = os.WriteFile("config.yaml", byt, 0666)
		confFile = byt
	}
	var conf models.Config
	if err := yaml.Unmarshal(confFile, &conf); err != nil {
		panic(err)
	}
	global.Config = &conf
}

func main() {
	cmd.Execute()
}
