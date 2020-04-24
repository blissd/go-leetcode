package main

func intToRoman(num int) string {
	if num == 0 {
		return ""
	}
	switch {
	case num >= 1000:
		return "M" + intToRoman(num-1000)
	case num >= 900:
		return "CM" + intToRoman(num-900)
	case num >= 500:
		return "D" + intToRoman(num-500)
	case num >= 400:
		return "CD" + intToRoman(num-400)
	case num >= 100:
		return "C" + intToRoman(num-100)
	case num >= 90:
		return "XC" + intToRoman(num-90)
	case num >= 50:
		return "L" + intToRoman(num-50)
	case num >= 40:
		return "XL" + intToRoman(num-40)
	case num >= 10:
		return "X" + intToRoman(num-10)
	case num == 9:
		return "IX"
	case num == 8:
		return "VIII"
	case num == 7:
		return "VII"
	case num == 6:
		return "VI"
	case num == 5:
		return "V"
	case num == 4:
		return "IV"
	case num < 4:
		return "I" + intToRoman(num-1)
	}
	return ""
}
