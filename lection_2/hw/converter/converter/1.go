package converter

// A Converter containing a number translation method
type Converter interface {
	IntToRoman(number int) string
}

// Contains values for translation from one number system to another
type converter struct {
	digits  []int
	symbols []string
}

// IntToRoman takes a number from the decimal system and returns a string in Roman
func (c *converter) IntToRoman(number int) string {
	var roman string
	// While long as the number is greater than or equal to the value from the slice
	// we subtract this value and add the corresponding Roman symbol
	for idx, digit := range c.digits {
		for number >= digit {
			number -= digit
			roman += c.symbols[idx]
		}
	}
	return roman
}

// NewConverter Interface constructor
func NewConverter(values []int, symbols []string) Converter {
	return &converter{values, symbols}
}
