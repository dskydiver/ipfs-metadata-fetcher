package cmd

import (
	"log"
	"sync"

	"github.com/dskydiver/ipfs-metadata-fetcher/cmd/scraper/pkg"
	"github.com/dskydiver/ipfs-metadata-fetcher/pkg/database"
	"github.com/dskydiver/ipfs-metadata-fetcher/pkg/model"
	"github.com/spf13/cobra"
)

var MaxConcurrenyLimit = 5

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrape IPFS metadata",
	Long:  `This command starts the process of scraping metadata from a list of IPFS CIDs provided in a .csv file.`,
	Run:   Scrape,
}

func init() {
	scrapeCmd.Flags().StringP("file", "f", "", "Path to .csv file containing CIDs")
	rootCmd.AddCommand(scrapeCmd)
}

func Scrape(cmd *cobra.Command, args []string) {
	done := make(chan struct{})
	errChan := make(chan error)
	defer close(done)

	file, _ := cmd.Flags().GetString("file")
	jobChan := pkg.ReadCIDs(done, file, errChan)

	go func() {
		for err := range errChan {
			log.Println(err)
		}
	}()

	dbClient := database.NewClient()
	defer dbClient.Close()

	var wg sync.WaitGroup

	modelChan := make(chan model.Model)

	for i := 0; i < MaxConcurrenyLimit; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Println("starting scraper")
			scrape := pkg.NewScraper(done, jobChan, modelChan, errChan)
			scrape.Run()
			log.Println("Finishing scraper")
		}()
	}

	go func() {
		wg.Wait()
		close(modelChan)
	}()

	saver := pkg.NewSaver(modelChan, dbClient, errChan)
	saver.Run()
}
