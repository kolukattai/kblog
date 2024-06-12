/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kolukattai/kblog/internal/global"
	"github.com/spf13/cobra"
)

// previewCmd represents the preview command
var previewCmd = &cobra.Command{
	Use:   "preview",
	Short: "preview the generate html page in given port or default port 3333",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")

		dir := http.Dir(global.Config.OutputFolder)

		fs := http.FileServer(dir)

		mux := http.NewServeMux()

		mux.Handle("/", fs)

		p := fmt.Sprintf(":%s", port)

		fmt.Printf("preview website stated at http://localhost%v\n", p)

		log.Fatal(http.ListenAndServe(p, mux))
	},
}

func init() {
	rootCmd.AddCommand(previewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// previewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// previewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	previewCmd.PersistentFlags().String("port", "3333", "port for the application to preview")
}
