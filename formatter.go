package money

// Formatter stores Money formatting information.
type Formatter struct {
	Fraction int
	Decimal  string
	Thousand string
	Grapheme string
	Template string
}

// NewFormatter creates new Formatter instance.
func NewFormatter(fraction int, decimal, thousand, grapheme, template string) *Formatter {
	_ = "STUB: not implemented"
	return nil
}

// Format returns string of formatted integer using given currency template.
func (f *Formatter) Format(amount int64) string {
	_ = "STUB: not implemented"
	// Work with absolute amount value
	return ""
}

// Add minus sign for negative amount.

// ToMajorUnits returns float64 representing the value in sub units using the currency data
func (f *Formatter) ToMajorUnits(amount int64) float64 { _ = "STUB: not implemented"; return 0 }

// abs return absolute value of given integer.
func (f Formatter) abs(amount int64) int64 { _ = "STUB: not implemented"; return 0 }
