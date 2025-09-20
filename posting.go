package tinygram

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// Posting represents a single document’s entry in a trigram’s posting list.
// Each posting stores:
//   - DocID:     the unique identifier of the document
//   - DocLength: the total number of trigrams extracted from the document field
//   - Frequency: how many times this trigram appears in that document
type Posting struct {
	DocID     string
	DocLength uint16
	Frequency uint8
}

// serializes a posting in binary
func serializePosting(p Posting) ([]byte, error) {
	buf := new(bytes.Buffer)

	// write docID
	if len(p.DocID) > 65535 {
		return nil, fmt.Errorf("docID too long")
	}
	if err := binary.Write(buf, binary.LittleEndian, uint16(len(p.DocID))); err != nil {
		return nil, err
	}
	if _, err := buf.WriteString(p.DocID); err != nil {
		return nil, err
	}

	// write docLength
	if err := binary.Write(buf, binary.LittleEndian, p.DocLength); err != nil {
		return nil, err
	}

	// write frequency
	if err := buf.WriteByte(p.Frequency); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// de-serializes a posting from binary
func deserializePosting(r io.Reader) (Posting, error) {
	var p Posting

	var idLen uint16
	if err := binary.Read(r, binary.LittleEndian, &idLen); err != nil {
		return p, err
	}

	idBytes := make([]byte, idLen)
	if _, err := r.Read(idBytes); err != nil {
		return p, err
	}
	p.DocID = string(idBytes)

	if err := binary.Read(r, binary.LittleEndian, &p.DocLength); err != nil {
		return p, err
	}

	if err := binary.Read(r, binary.LittleEndian, &p.Frequency); err != nil {
		return p, err
	}

	return p, nil
}

func serializeDocFreq(freq uint32) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, freq)
	return buf
}

func deserializeDocFreq(data []byte) uint32 {
	return binary.LittleEndian.Uint32(data)
}
