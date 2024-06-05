/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/charmbracelet/log"

	"os"
	"strings"

	"github.com/kolukattai/kblog/models"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Prealse provide file name")
		}
		fileName := strings.ReplaceAll(strings.Join(args, "-"), " ", "-")
		pageData := &models.PageData{
			Title:        strings.Join(args, " "),
			Description:  "post description",
			Keywords:     "one, two, three",
			Tags:         "one, two, three",
			Category:     "undefined",
			Author:       "<your name>",
			LandingImage: "",
		}
		yamlData, err := yaml.Marshal(pageData)
		if err != nil {
			log.Fatal(err.Error())
		}
		data := fmt.Sprintf("---\n%s---\n\n# %s\n\npage content goes heare", string(yamlData), pageData.Title)
		path := fmt.Sprintf("posts/%s.md", fileName)

		_, err = os.Stat(path)

		if err == nil {
			log.Error("failed to create post", "err", "post with same name already exists")
			os.Exit(1)
		}

		err = os.WriteFile(path, []byte(data), 0666)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s created", path)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
