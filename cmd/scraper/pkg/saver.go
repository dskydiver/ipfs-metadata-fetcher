package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/dskydiver/ipfs-metadata-fetcher/ent"
	"github.com/dskydiver/ipfs-metadata-fetcher/pkg/model"
)

type Saver struct {
	ModelChan <-chan model.Model
	Client    *ent.Client
	ErrChan   chan<- error
}

func NewSaver(modelChan <-chan model.Model, client *ent.Client, errChan chan<- error) *Saver {
	return &Saver{
		ModelChan: modelChan,
		Client:    client,
		ErrChan:   errChan,
	}
}

func (s *Saver) Run() {
	for model := range s.ModelChan {
		log.Println(model.CID)
		_, err := s.Client.Metadata.Create().
			SetCid(model.CID).
			SetImage(model.Image).
			SetDescription(model.Description).
			SetName(model.Name).
			Save(context.Background())
		if err != nil {
			s.ErrChan <- fmt.Errorf("error saving metadata for %s: %s", model.CID, err)
		}
	}
}
