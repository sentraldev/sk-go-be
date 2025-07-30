package utils

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JSONBArray []string

// Value implements the driver.Valuer interface
func (a JSONBArray) Value() (driver.Value, error) {
	if len(a) == 0 {
		return "[]", nil
	}
	return json.Marshal(a)
}

// Scan implements the sql.Scanner interface
func (a *JSONBArray) Scan(value interface{}) error {
	if value == nil {
		*a = []string{}
		return nil
	}

	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		// Handle string input by converting to []byte
		bytes = []byte(v)
	default:
		return fmt.Errorf("failed to scan JSONBArray: expected []byte or string, got %T", value)
	}

	if len(bytes) == 0 {
		*a = []string{}
		return nil
	}
	return json.Unmarshal(bytes, a)
}
