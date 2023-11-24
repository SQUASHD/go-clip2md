/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/SQUASHD/go-clip2md/internal/clippy"
	"github.com/SQUASHD/go-clip2md/internal/converter"
	"github.com/SQUASHD/go-clip2md/internal/postprocess"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-clip2md",
	Short: "go-clip2md is a tool to convert clipboard content to markdown",
	Long:  `go-clip2md is a tool to convert clipboard content to markdown`,
	Run: func(cmd *cobra.Command, args []string) {

		content, err := clippy.ReadFromClipboard()
		if err != nil {
			fmt.Println("Error reading from clipboard")
			return
		}

		language, _ := cmd.Flags().GetString("lang")
		domain, _ := cmd.Flags().GetString("domain")
		md, err := converter.ConvertHTML2Md(domain, content)
		if err != nil {
			fmt.Println("Error converting to markdown")
			return
		}
		md = postprocess.PostProcessMarkdown(md, language)

		if err = clippy.WriteToClipboard(md); err != nil {
			fmt.Println("Error writing to clipboard")
			return
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("lang", "l", "", "Language to use for syntax highlighting")
	rootCmd.PersistentFlags().StringP("domain", "d", "", "Domain to use for relative links, e.g. https://en.wikipedia.org")
}
