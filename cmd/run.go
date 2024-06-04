/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/kolukattai/kblog/global"
	"github.com/kolukattai/kblog/parser"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("running blog")

		port, _ := cmd.Flags().GetString("port")

		path := "."

		if len(args) > 0 {
			path = args[0]
		}

		_, err := os.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}

		var staticFS = http.FS(global.StaticFiles)
		fs := http.FileServer(staticFS)

		http.Handle("/static/", fs)

		http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			path := r.URL.Path[1:]
			fmt.Println("PATH", path, len(path))
			if len(path) == 0 {
				w.Write([]byte("home.Home()"))
				return
			}
			if strings.Index(path, "posts") == 0 {
				val := parser.Parse(strings.Replace(path, "posts/", "", 1), parser.Options{
					LandingImage: true,
					Tags:         true,
					Author:       true,
					Config:       global.Config,
				})
				w.Write([]byte(val))
				return
			}
			w.Write([]byte("var"))
		}))
		fmt.Printf("application stated at http://localhost:%s\n", port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
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
	runCmd.PersistentFlags().String("port", "8080", "port for the application to run")
}
