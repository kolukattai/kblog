/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/kolukattai/kblog/internal/boot"
	"github.com/kolukattai/kblog/internal/global"
	"github.com/kolukattai/kblog/internal/server"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run development server",
	Long:  `run development server either in default port 8080 or in provide port using port flag`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")

		port, _ := cmd.Flags().GetString("port")

		boot.InitSiteData()

		boot.InitJavascriptMaps(global.PageDataList, global.Config.PerPage)

		boot.InitPostData(global.PageDataList, global.Config.PerPage)

		boot.InitTagAndCategoryData(global.PageDataList, global.Tags, global.Categories, global.Config.PerPage)

		server.Run(port)

	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	runCmd.Flags().String("port", "8080", "port for the application to run")
}
