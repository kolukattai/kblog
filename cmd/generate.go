/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/kolukattai/kblog/internal/boot"
	"github.com/kolukattai/kblog/internal/build"
	"github.com/kolukattai/kblog/internal/global"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generating static files")

		boot.InitSiteData()

		boot.InitJavascriptMaps(global.PageDataList, global.Config.PerPage)

		boot.InitPostData(global.PageDataList, global.Config.PerPage)

		boot.InitTagAndCategoryData(global.PageDataList, global.Tags, global.Categories, global.Config.PerPage)

		build.Exec()

		fmt.Printf("build generated in \"%v\"", global.Config.OutputFolder)
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
