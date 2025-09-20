package tinygram

// Document represents one record to be indexed
type Document struct {
	ID     string
	Fields map[string]Field
}

type FieldType int

const (
	FieldText FieldType = iota
)

type Field struct {
	Name  string
	Type  FieldType
	Value any
}

func NewDocument(id string) *Document {
	return &Document{
		ID:     id,
		Fields: make(map[string]Field),
	}
}

func (doc *Document) AddTextField(name string, value string) {
	doc.Fields[name] = Field{
		Name:  name,
		Type:  FieldText,
		Value: value,
	}
}
