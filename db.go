package money

import (
	"database/sql/driver"
)

var (
	// DBMoneyValueSeparator is used to join together the Amount and Currency components of money.Money instances
	// allowing them to be stored as strings (via the driver.Valuer interface) and unmarshalled as strings (via
	// the sql.Scanner interface); set this value to use a different separator.
	DBMoneyValueSeparator = DefaultDBMoneyValueSeparator
)

const (
	// DefaultDBMoneyValueSeparator is the default value for DBMoneyValueSeparator; can be used to reset the
	// active separator value
	DefaultDBMoneyValueSeparator = "|"
)

// Value implements driver.Valuer to serialise a Money instance into a delimited string using the DBMoneyValueSeparator
// for example: "amount|currency_code"
func (m *Money) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// Scan implements sql.Scanner to deserialize a Money instance from a DBMoneyValueSeparator-separated string
// for example: "amount|currency_code"
func (m *Money) Scan(src interface{}) error { _ = "STUB: not implemented"; return nil }

// let's support string and int64

// allocate new Money with the scanned amount and currency

// Value implements driver.Valuer to serialize a Currency code into a string for saving to a database
func (c Currency) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *

	// Scan implements sql.Scanner to deserialize a Currency from a string value read from a database
	new(driver.Value), nil
}

func (c *Currency) Scan(src interface{}) error {
	_ = "STUB: not implemented"

	// let's support string only
	return nil
}

// copy the value
