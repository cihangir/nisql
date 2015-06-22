package nisql

import (
	"database/sql"
	"encoding/json"
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
