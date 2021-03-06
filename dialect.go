package dbr

import "time"

// Dialect abstracts database differences
type Dialect interface {
	QuoteIdent(id string) string

	EncodeString(s string) string
	EncodeBool(b bool) string
	EncodeTime(t time.Time) string
	EncodeBytes(b []byte) string
	Placeholder(n int) string
	OnConflict(constraint string) string
	Proposed(column string) string
	Limit(offset, limit int64) string
}
