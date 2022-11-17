package converter

type Converter interface {
	IntToRoman(number int) string
}

type converter struct {
	digits  []int
	symbols []string
}

// IntToRoman метод преобразования числа из десятичной системы в римскую
func (c *converter) IntToRoman(number int) string {
	var roman string
	// цикл по словарю
	for idx, digit := range c.digits {
		// пока число больше или равно значению из словаря
		for number >= digit {
			// вычитаем значение из числа
			number -= digit
			// добавляем символ
			roman += c.symbols[idx]
		}
	}
	return roman

}

func NewConverter(values []int, symbols []string) Converter {
	return &converter{values, symbols}
}
