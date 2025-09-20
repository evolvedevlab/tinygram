package tinygram

import (
	"errors"
	"fmt"

	"github.com/dgraph-io/badger/v4"
)

var (
	ErrInvalidFieldType = errors.New("invalid field type")
)

type Index struct {
	db *badger.DB
}

func NewIndex(path string) (*Index, error) {
	opts := badger.DefaultOptions(path)
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	return &Index{db: db}, nil
}

func (idx *Index) IndexDocument(doc *Document) error {
	for _, field := range doc.Fields {
		switch field.Type {
		case FieldText:
			s, ok := field.Value.(string)
			if !ok {
				return fmt.Errorf("text field requires string value")
			}

			// generating trigrams from the given string
			trigrams := generateTrigrams(s)

			docLen := uint16(len(trigrams))
			freqs := make(map[string]uint8)
			for _, trigram := range trigrams {
				freqs[trigram]++
			}

			_ = docLen

			// todo
		}
	}

	return ErrInvalidFieldType
}
