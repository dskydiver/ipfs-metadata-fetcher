package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ipfs-scraper",
	Short: "A CLI tool to scrape metadata from IPFS",
	Long:  "This tool scrapes metadata from IPFS from given CIDs and stores the result in a specified PostgreSQL database.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
