package mongodbcustomtype

import (
	"bytes"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

type StringUUID string

func NewStringUUID() StringUUID {
	id := uuid.NewV4()
	return StringUUID(id.String())
}

type UUID struct{ uuid.UUID }

const uuidSubType byte = 0x04

// NewUUID generates a new MongoDB compatible UUID.
func NewUUID() UUID {
	id := uuid.NewV4()
	return UUID{UUID: id}
}

// UUIDFromStringOrNil returns a UUID parsed from the input string.
func UUIDFromStringOrNil(input string) *UUID {
	id := uuid.FromStringOrNil(input)
	if id == uuid.Nil {
		return nil
	}
	return &UUID{id}
}

// MarshalBSONValue implements the bson.ValueMarshaler interface.
func (id *UUID) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bsontype.Binary, bsoncore.AppendBinary(nil, 4, id.UUID[:]), nil
}

// UnmarshalBSONValue implements the bson.ValueUnmarshaler interface.
func (id *UUID) UnmarshalBSONValue(t bsontype.Type, raw []byte) error {
	if t != bsontype.Binary {
		return fmt.Errorf("invalid format on unmarshal bson value")
	}

	_, data, _, ok := bsoncore.ReadBinary(raw)
	if !ok {
		return fmt.Errorf("not enough bytes to unmarshal bson value")
	}

	copy(id.UUID[:], data)

	return nil
}

// IsZero implements the bson.Zeroer interface.
func (id *UUID) IsZero() bool {
	return bytes.Compare(id.Bytes(), uuid.Nil.Bytes()) == 0
}
