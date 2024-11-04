package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var url string

func main() {
	RootCommand.Execute()
}

var RootCommand = &cobra.Command{
	Use:   "checker",
	Short: "Website Health Checker",
	Long:  "This tool checks if a website is up or down.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'checker check --url <website>' to check a website's health.")
	},
}

var CheckCommand = &cobra.Command{
	Use:   "check",
	Short: "Check Website Health",
	Long:  "Checks if the specified website is up or down.",
	Run: func(cmd *cobra.Command, args []string) {
		checkWebsite(url)
	},
}

func init() {
	CheckCommand.Flags().StringVarP(&url, "url", "u", "", "URL of the website to check")
	CheckCommand.MarkFlagRequired("url")
	RootCommand.AddCommand(CheckCommand)
}

func checkWebsite(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Website is up:", url)
	} else {
		fmt.Printf("Website is down: %s (Status Code: %d)\n", url, resp.StatusCode)
	}
}
