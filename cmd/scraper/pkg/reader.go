package pkg

import (
	"bufio"
	"fmt"
	"os"
)

func ReadCIDs(done <-chan struct{}, filepath string, errChan chan<- error) <-chan Job {
	jobChan := make(chan Job)

	go func() {
		defer close(jobChan)
		defer close(errChan)
		file, err := os.Open(filepath)
		if err != nil {
			errChan <- err
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			cid := scanner.Text()
			jobChan <- Job{CID: cid}
		}

		if err := scanner.Err(); err != nil {
			errChan <- fmt.Errorf("error reading file: %s", err)
			return
		}
	}()
	return jobChan
}
