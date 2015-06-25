// Package nisql provides nullable types for database operations
package nisql

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/lib/pq"
)

var nullString = []byte("null")

// String creates a new valid NullString
func String(s string) NullString {
	return NullString{
		sql.NullString{
			String: s,
			Valid:  true,
		},
	}
}

// NullString is a type that can be null or a string
type NullString struct {
	sql.NullString
}

// MarshalJSON correctly serializes a NullString to JSON
func (n *NullString) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return nullString, nil
	}

	return json.Marshal(n.String)
}

// UnmarshalJSON turns *NullString into a json.Unmarshaller.
func (n *NullString) UnmarshalJSON(b []byte) error {
	return unmarshal(n, b)
}

// Float64 creates a new valid NullFloat64
func Float64(f float64) NullFloat64 {
	return NullFloat64{
		sql.NullFloat64{
			Float64: f,
			Valid:   true,
		},
	}
}

// NullFloat64 is a type that can be null or a float64
type NullFloat64 struct {
	sql.NullFloat64
}

// MarshalJSON correctly serializes a NullFloat64 to JSON
func (n *NullFloat64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return nullString, nil
	}

	return json.Marshal(n.Float64)
}

// UnmarshalJSON turns *NullFloat64 into a json.Unmarshaller.
func (n *NullFloat64) UnmarshalJSON(b []byte) error {
	return unmarshal(n, b)
}

// Int64 creates a new valid NullInt64
func Int64(i int64) NullInt64 {
	return NullInt64{
		sql.NullInt64{
			Int64: i,
			Valid: true,
		},
	}
}

// NullInt64 is a type that can be null or an int
type NullInt64 struct {
	sql.NullInt64
}

// MarshalJSON correctly serializes a NullInt64 to JSON
func (n *NullInt64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return nullString, nil
	}

	return json.Marshal(n.Int64)
}

// UnmarshalJSON turns *NullInt64 into a json.Unmarshaller.
func (n *NullInt64) UnmarshalJSON(b []byte) error {
	return unmarshal(n, b)
}

// Bool creates a new valid NullBool
func Bool(b bool) NullBool {
	return NullBool{
		sql.NullBool{
			Bool:  b,
			Valid: true,
		},
	}
}

// NullBool is a type that can be null or a bool
type NullBool struct {
	sql.NullBool
}

// MarshalJSON correctly serializes a NullBool to JSON
func (n *NullBool) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return nullString, nil
	}

	return json.Marshal(n.Bool)
}

// UnmarshalJSON turns *NullBool into a json.Unmarshaller.
func (n *NullBool) UnmarshalJSON(b []byte) error {
	return unmarshal(n, b)
}

// Time creates a new valid NullTime
func Time(t time.Time) NullTime {
	return NullTime{
		pq.NullTime{
			Time:  t,
			Valid: true,
		},
	}
}

// NullTime is a type that can be null or a time.Time
type NullTime struct {
	pq.NullTime
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
	if bytes.Equal(b, nullString) {
		return n.Scan(nil)
	}

	var t time.Time
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	return n.Scan(t)
}

func unmarshal(s sql.Scanner, b []byte) error {
	var d interface{}
	if err := json.Unmarshal(b, &d); err != nil {
		return err
	}

	return s.Scan(d)
}
