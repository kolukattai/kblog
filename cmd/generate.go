/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/kolukattai/kblog/internal/boot"
	"github.com/kolukattai/kblog/internal/build"
	"github.com/kolukattai/kblog/internal/global"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate static html files",
	Long:  `generate static html files`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generating static files")

		t := time.Now()

		boot.InitSiteData()

		boot.InitJavascriptMaps(global.PageDataList, global.Config.PerPage)

		boot.InitPostData(global.PageDataList, global.Config.PerPage)

		boot.InitTagAndCategoryData(global.PageDataList, global.Tags, global.Categories, global.Config.PerPage)

		build.Exec()

		since := time.Since(t)

		fmt.Printf("build generated in %v\n", since)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
