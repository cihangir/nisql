package nisql

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

var nullString = []byte("null")

// NullString is a type that can be null or a string
type NullString struct {
	sql.NullString
}

// NullFloat64 is a type that can be null or a float64
type NullFloat64 struct {
	sql.NullFloat64
}

// NullInt64 is a type that can be null or an int
type NullInt64 struct {
	sql.NullInt64
}

// NullBool is a type that can be null or a bool
type NullBool struct {
	sql.NullBool
}

// MarshalJSON correctly serializes a NullString to JSON
func (n *NullString) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return nullString, nil
	}

	return json.Marshal(n.String)
}

// MarshalJSON correctly serializes a NullInt64 to JSON
func (n *NullInt64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return nullString, nil
	}

	return json.Marshal(n.Int64)
}

// MarshalJSON correctly serializes a NullFloat64 to JSON
func (n *NullFloat64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return nullString, nil
	}

	return json.Marshal(n.Float64)
}

// MarshalJSON correctly serializes a NullBool to JSON
func (n *NullBool) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return nullString, nil
	}

	return json.Marshal(n.Bool)
}

// UnmarshalJSON turns *NullString into a json.Unmarshaller.
func (n *NullString) UnmarshalJSON(b []byte) error {
	return unmarshal(n, b)
}

// UnmarshalJSON turns *NullInt64 into a json.Unmarshaller.
func (n *NullInt64) UnmarshalJSON(b []byte) error {
	return unmarshal(n, b)
}

// UnmarshalJSON turns *NullFloat64 into a json.Unmarshaller.
func (n *NullFloat64) UnmarshalJSON(b []byte) error {
	return unmarshal(n, b)
}

// UnmarshalJSON turns *NullBool into a json.Unmarshaller.
func (n *NullBool) UnmarshalJSON(b []byte) error {
	return unmarshal(n, b)
}

func unmarshal(s sql.Scanner, b []byte) error {
	var d interface{}
	if err := json.Unmarshal(b, &d); err != nil {
		return err
	}

	return s.Scan(d)
}

type NullTime struct {
	Time  time.Time
	Valid bool
}

func (n *NullTime) Scan(value interface{}) error {
	if value == nil {
		n.Time, n.Valid = time.Time{}, false
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		n.Time, n.Valid = v, true
		return nil
	}

	n.Valid = false
	return fmt.Errorf("Can't convert %T to time.Time", value)
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}

	return nt.Time, nil
}

// MarshalJSON correctly serializes a NullTime to JSON
func (n *NullTime) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return nullString, nil
	}

	return json.Marshal(n.Time)
}

// UnmarshalJSON turns *NullTime into a json.Unmarshaller.
func (n *NullTime) UnmarshalJSON(b []byte) error {
	// scan for JSON timestamp
	var t time.Time
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	return n.Scan(t)
}
