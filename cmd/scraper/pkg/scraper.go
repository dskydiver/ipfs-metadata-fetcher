package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dskydiver/ipfs-metadata-fetcher/pkg/model"
)

type Job struct {
	CID string
}

type Scraper struct {
	JobChan   <-chan Job
	ModelChan chan<- model.Model
	ErrChan   chan<- error
	Client    *http.Client
	done      <-chan struct{}
}

func NewScraper(done <-chan struct{}, jobChan <-chan Job, modelChan chan<- model.Model, errChan chan<- error) *Scraper {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	return &Scraper{
		JobChan:   jobChan,
		ModelChan: modelChan,
		ErrChan:   errChan,
		Client:    client,
		done:      done,
	}
}

func (s *Scraper) Run() {
	for job := range s.JobChan {
		func() {
			resp, err := s.Client.Get(fmt.Sprintf("https://ipfs.io/ipfs/%s", job.CID))
			if err != nil {
				s.ErrChan <- err
				return
			}
			defer resp.Body.Close()

			var data model.Model

			err = json.NewDecoder(resp.Body).Decode(&data)

			// cmd := exec.Command("curl", fmt.Sprintf("https://ipfs.io/ipfs/%s", job.CID))
			// output, err := cmd.Output()
			// if err != nil {
			// 	s.ErrChan <- err
			// 	return
			// }

			// var data model.Model

			// err = json.Unmarshal(output, &data)

			if err != nil {
				s.ErrChan <- fmt.Errorf("error unmarshalling metadata for %s: %s", job.CID, err)
				return
			}

			data.CID = job.CID

			select {
			case s.ModelChan <- data:
			case <-s.done:
				return
			}
		}()
	}
}
