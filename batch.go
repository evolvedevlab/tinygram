package tinygram

import "fmt"

// Batch represents a temporary container for multiple documents to be inserted together
type Batch struct {
	store *Store
	data  map[string][]string
}

// Append adds a document to the batch.
func (b *Batch) Append(doc *Document) error {
	for _, field := range doc.Fields {
		// TODO: do something for the other field types in future
		if field.Type != FieldText {
			continue
		}

		text, ok := field.Value.(string)
		if !ok {
			return fmt.Errorf("text field requires string valu")
		}

		trigrams := generateTrigrams(text)
		b.data[doc.ID] = trigrams
	}

	return nil
}

// Flush writes all documents in the batch to the store
func (b *Batch) Flush() error {
	if len(b.data) == 0 {
		return nil
	}
	return b.store.InsertMany(b.data)
}
