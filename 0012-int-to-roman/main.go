package main

// represents a lower level and the numeral for that level
type band struct {
	i int    // integer
	n string // numeral
}

var bands = []band{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	roman := make([]uint8, 0, 10)
	for _, b := range bands {
		for num >= b.i {
			roman = append(roman, b.n...)
			num -= b.i
		}
	}
	return string(roman)
}
