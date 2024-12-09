package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONB map[string]string

// Scan implements the sql.Scanner interface.
func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = make(JSONB)
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, j)
}

// Value implements the driver.Valuer interface.
func (j JSONB) Value() (driver.Value, error) {
	return json.Marshal(j)
}
