package money

import (
	"errors"
)

// Injection points for backward compatibility.
// If you need to keep your JSON marshal/unmarshal way, overwrite them like below.
//
//	money.UnmarshalJSON = func (m *Money, b []byte) error { ... }
//	money.MarshalJSON = func (m Money) ([]byte, error) { ... }
var (
	// UnmarshalJSON is injection point of json.Unmarshaller for money.Money
	UnmarshalJSON = defaultUnmarshalJSON
	// MarshalJSON is injection point of json.Marshaller for money.Money
	MarshalJSON = defaultMarshalJSON

	// ErrCurrencyMismatch happens when two compared Money don't have the same currency.
	ErrCurrencyMismatch = errors.New("currencies don't match")

	// ErrInvalidJSONUnmarshal happens when the default money.UnmarshalJSON fails to unmarshal Money because of invalid data.
	ErrInvalidJSONUnmarshal = errors.New("invalid json unmarshal")
)

func defaultUnmarshalJSON(m *Money, b []byte) error { _ = "STUB: not implemented"; return nil }

func defaultMarshalJSON(m Money) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Amount is a data structure that stores the amount being used for calculations.
type Amount = int64

// Money represents monetary value information, stores
// currency and amount value.
type Money struct {
	amount   Amount    `db:"amount"`
	currency *Currency `db:"currency"`
}

// New creates and returns new instance of Money.
func New(amount int64, code string) *Money { _ = "STUB: not implemented"; return nil }

// NewFromFloat creates and returns new instance of Money from a float64.
// Always rounding trailing decimals down.
func NewFromFloat(amount float64, code string) *Money { _ = "STUB: not implemented"; return nil }

// Currency returns the currency used by Money.
func (m *Money) Currency() *Currency {
	_ = "STUB: not implemented"

	// Amount returns a copy of the internal monetary value as an int64.
	return nil
}

func (m *Money) Amount() int64 {
	_ = "STUB: not implemented"

	// SameCurrency check if given Money is equals by currency.
	return 0
}

func (m *Money) SameCurrency(om *Money) bool { _ = "STUB: not implemented"; return false }

func (m *Money) assertSameCurrency(om *Money) error { _ = "STUB: not implemented"; return nil }

func (m *Money) compare(om *Money) int { _ = "STUB: not implemented"; return 0 }

// Equals checks equality between two Money types.
func (m *Money) Equals(om *Money) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// GreaterThan checks whether the value of Money is greater than the other.
func (m *Money) GreaterThan(om *Money) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// GreaterThanOrEqual checks whether the value of Money is greater or equal than the other.
func (m *Money) GreaterThanOrEqual(om *Money) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// LessThan checks whether the value of Money is less than the other.
func (m *Money) LessThan(om *Money) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// LessThanOrEqual checks whether the value of Money is less or equal than the other.
func (m *Money) LessThanOrEqual(om *Money) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IsZero returns boolean of whether the value of Money is equals to zero.
func (m *Money) IsZero() bool { _ = "STUB: not implemented"; return false }

// IsPositive returns boolean of whether the value of Money is positive.
func (m *Money) IsPositive() bool { _ = "STUB: not implemented"; return false }

// IsNegative returns boolean of whether the value of Money is negative.
func (m *Money) IsNegative() bool { _ = "STUB: not implemented"; return false }

// Absolute returns new Money struct from given Money using absolute monetary value.
func (m *Money) Absolute() *Money { _ = "STUB: not implemented"; return nil }

// Negative returns new Money struct from given Money using negative monetary value.
func (m *Money) Negative() *Money { _ = "STUB: not implemented"; return nil }

// Add returns new Money struct with value representing sum of Self and Other Money.
func (m *Money) Add(ms ...*Money) (*Money, error) { _ = "STUB: not implemented"; return nil, nil }

// Subtract returns new Money struct with value representing difference of Self and Other Money.
func (m *Money) Subtract(ms ...*Money) (*Money, error) { _ = "STUB: not implemented"; return nil, nil }

// Multiply returns new Money struct with value representing Self multiplied value by multiplier.
func (m *Money) Multiply(muls ...int64) *Money { _ = "STUB: not implemented"; return nil }

// Round returns new Money struct with value rounded to nearest zero.
func (m *Money) Round() *Money { _ = "STUB: not implemented"; return nil }

// Split returns slice of Money structs with split Self value in given number.
// After division leftover pennies will be distributed round-robin amongst the parties.
// This means that parties listed first will likely receive more pennies than ones that are listed later.
func (m *Money) Split(n int) ([]*Money, error) { _ = "STUB: not implemented"; return nil, nil }

// Add leftovers to the first parties.

// Allocate returns slice of Money structs with split Self value in given ratios.
// It lets split money by given ratios without losing pennies and as Split operations distributes
// leftover pennies amongst the parties with round-robin principle.
func (m *Money) Allocate(rs ...int) ([]*Money, error) { _ = "STUB: not implemented"; return nil, nil }

// Calculate sum of ratios.

// if the sum of all ratios is zero, then we just returns zeros and don't do anything
// with the leftover

// Calculate leftover value and divide to first parties.

// Display lets represent Money struct as string in given Currency value.
func (m *Money) Display() string { _ = "STUB: not implemented"; return "" }

// AsMajorUnits lets represent Money struct as subunits (float64) in given Currency value
func (m *Money) AsMajorUnits() float64 { _ = "STUB: not implemented"; return 0 }

// UnmarshalJSON is implementation of json.Unmarshaller
func (m *Money) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON is implementation of json.Marshaller
func (m Money) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Compare function compares two money of the same type
		//
		//	if m.amount > om.amount returns (1, nil)
		//	if m.amount == om.amount returns (0, nil
		//	if m.amount < om.amount returns (-1, nil)
		//
		// If compare moneys from distinct currency, return (m.amount, ErrCurrencyMismatch)
		nil
}

func (m *Money) Compare(om *Money) (int, error) { _ = "STUB: not implemented"; return 0, nil }
