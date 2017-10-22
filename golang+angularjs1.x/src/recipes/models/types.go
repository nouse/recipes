package models

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// From https://gist.github.com/mraaroncruz/87808f979bdfcae734ff337ff2066d2c
type jsonb json.RawMessage

func (j jsonb) Value() (driver.Value, error) {
	if j.IsNull() {
		return nil, nil
	}
	return string(j), nil
}

func (j *jsonb) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("scan source was not string")
	}
	// I think I need to make a copy of the bytes.
	// It seems the byte slice passed in is re-used
	*j = append((*j)[0:0], s...)

	return nil
}

func (j jsonb) IsNull() bool {
	return len(j) == 0 || bytes.Equal(j, []byte("null"))
}

// Same as json.RawMessage
func (m jsonb) MarshalJSON() ([]byte, error) {
	return json.RawMessage(m).MarshalJSON()
}

// Copied from json.RawMessage
func (m *jsonb) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("jsonb: UnmarshalJSON on nil pointer")
	}

	*m = append((*m)[0:0], data...)
	return nil
}
