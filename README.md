[![Build Status](https://travis-ci.org/cihangir/nisql.svg?branch=master)](https://travis-ci.org/cihangir/nisql)
[![GoDoc](https://godoc.org/github.com/cihangir/nisql?status.svg)](https://godoc.org/github.com/cihangir/nisql)

# nisql
--
    import "github.com/cihangir/nisql"

Package nisql provides nullable types for database operations

## Usage

#### type NullBool

```go
type NullBool struct {
    sql.NullBool
}
```

NullBool is a type that can be null or a bool

#### func  Bool

```go
func Bool(b bool) NullBool
```
Bool creates a new valid NullBool

#### func (*NullBool) MarshalJSON

```go
func (n *NullBool) MarshalJSON() ([]byte, error)
```
MarshalJSON correctly serializes a NullBool to JSON

#### func (*NullBool) UnmarshalJSON

```go
func (n *NullBool) UnmarshalJSON(b []byte) error
```
UnmarshalJSON turns *NullBool into a json.Unmarshaller.

#### type NullFloat64

```go
type NullFloat64 struct {
    sql.NullFloat64
}
```

NullFloat64 is a type that can be null or a float64

#### func  Float64

```go
func Float64(f float64) NullFloat64
```
Float64 creates a new valid NullFloat64

#### func (*NullFloat64) MarshalJSON

```go
func (n *NullFloat64) MarshalJSON() ([]byte, error)
```
MarshalJSON correctly serializes a NullFloat64 to JSON

#### func (*NullFloat64) UnmarshalJSON

```go
func (n *NullFloat64) UnmarshalJSON(b []byte) error
```
UnmarshalJSON turns *NullFloat64 into a json.Unmarshaller.

#### type NullInt64

```go
type NullInt64 struct {
    sql.NullInt64
}
```

NullInt64 is a type that can be null or an int

#### func  Int64

```go
func Int64(i int64) NullInt64
```
Int64 creates a new valid NullInt64

#### func (*NullInt64) MarshalJSON

```go
func (n *NullInt64) MarshalJSON() ([]byte, error)
```
MarshalJSON correctly serializes a NullInt64 to JSON

#### func (*NullInt64) UnmarshalJSON

```go
func (n *NullInt64) UnmarshalJSON(b []byte) error
```
UnmarshalJSON turns *NullInt64 into a json.Unmarshaller.

#### type NullString

```go
type NullString struct {
    sql.NullString
}
```

NullString is a type that can be null or a string

#### func  String

```go
func String(s string) NullString
```
String creates a new valid NullString

#### func (*NullString) MarshalJSON

```go
func (n *NullString) MarshalJSON() ([]byte, error)
```
MarshalJSON correctly serializes a NullString to JSON

#### func (*NullString) UnmarshalJSON

```go
func (n *NullString) UnmarshalJSON(b []byte) error
```
UnmarshalJSON turns *NullString into a json.Unmarshaller.

#### type NullTime

```go
type NullTime struct {
    pq.NullTime
}
```

NullTime is a type that can be null or a time.Time

#### func  Time

```go
func Time(t time.Time) NullTime
```
Time creates a new valid NullTime

#### func (*NullTime) MarshalJSON

```go
func (n *NullTime) MarshalJSON() ([]byte, error)
```
MarshalJSON correctly serializes a NullTime to JSON

#### func (*NullTime) UnmarshalJSON

```go
func (n *NullTime) UnmarshalJSON(b []byte) error
```
UnmarshalJSON turns *NullTime into a json.Unmarshaller.
