package tinygram

import (
	"github.com/dgraph-io/badger/v4"
)

type Store struct {
	db *badger.DB
}

func (s *Store) InsertMany(docs map[string][]string) error {
	wb := s.db.NewWriteBatch()
	defer wb.Cancel()

	for docID, trigrams := range docs {
		docLen := uint16(len(trigrams))
		m := make(map[string]uint8)
		for _, tg := range trigrams {
			m[tg]++
		}

		for trigram, freq := range m {
			key := []byte("trigram:" + trigram + ":" + docID)
			post := Posting{
				DocID:     docID,
				DocLength: docLen,
				Frequency: freq,
			}

			b, err := serializePosting(post)
			if err != nil {
				return err
			}

			if err := wb.Set(key, b); err != nil {
				return err
			}
		}
	}

	return wb.Flush()
}

// Insert is a sequential method to insert trigrams for a document
func (s *Store) Insert(docID string, trigrams []string) error {
	wb := s.db.NewWriteBatch()
	defer wb.Cancel()

	docLen := uint16(len(trigrams))

	// map for storing trigram and its occurance frequency
	m := make(map[string]uint8)
	for _, tg := range trigrams {
		m[tg]++
	}

	for trigram, freq := range m {
		key := []byte("trigram:" + trigram + ":" + docID)
		p := Posting{
			DocID:     docID,
			DocLength: docLen,
			Frequency: freq,
		}

		b, err := serializePosting(p)
		if err != nil {
			return err
		}

		if err := wb.Set(key, b); err != nil {
			return err
		}
	}

	return wb.Flush()
}

func (s *Store) Close() error {
	return s.db.Close()
}
