package money

type calculator struct{}

func (c *calculator) add(a, b Amount) Amount { _ = "STUB: not implemented"; return *new(Amount) }

func (c *calculator) subtract(a, b Amount) Amount { _ = "STUB: not implemented"; return *new(Amount) }

func (c *calculator) multiply(a Amount, m int64) Amount {
	_ = "STUB: not implemented"
	return *new(Amount)
}

func (c *calculator) divide(a Amount, d int64) Amount {
	_ = "STUB: not implemented"
	return *new(Amount)
}

func (c *calculator) modulus(a Amount, d int64) Amount {
	_ = "STUB: not implemented"
	return *new(Amount)
}

func (c *calculator) allocate(a Amount, r, s int64) Amount {
	_ = "STUB: not implemented"
	return *new(Amount)
}

func (c *calculator) absolute(a Amount) Amount { _ = "STUB: not implemented"; return *new(Amount) }

func (c *calculator) negative(a Amount) Amount { _ = "STUB: not implemented"; return *new(Amount) }

func (c *calculator) round(a Amount, e int) Amount { _ = "STUB: not implemented"; return *new(Amount) }
