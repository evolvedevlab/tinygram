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
	store *Store
}

func NewIndex(path string) (*Index, error) {
	opts := badger.DefaultOptions(path)
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	return &Index{
		store: &Store{db: db},
	}, nil
}

func NewIndexWithOptions(opts badger.Options) (*Index, error) {
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	return &Index{
		store: &Store{db: db},
	}, nil
}

func (idx *Index) NewBatch() *Batch {
	return &Batch{
		store: idx.store,
		data:  make(map[string][]string),
	}
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

			return idx.store.Insert(doc.ID, trigrams)
		}
	}

	return ErrInvalidFieldType
}

func (idx *Index) Close() error {
	return idx.store.Close()
}
