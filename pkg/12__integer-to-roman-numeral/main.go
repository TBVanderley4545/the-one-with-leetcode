package integertoromannumeral

type RomanNumeral struct {
	Value   int
	Numeral string
}

var RomanNumerals = []RomanNumeral{
	{Value: 1000, Numeral: "M"},
	{Value: 900, Numeral: "CM"},
	{Value: 500, Numeral: "D"},
	{Value: 400, Numeral: "CD"},
	{Value: 100, Numeral: "C"},
	{Value: 90, Numeral: "XC"},
	{Value: 50, Numeral: "L"},
	{Value: 40, Numeral: "XL"},
	{Value: 10, Numeral: "X"},
	{Value: 9, Numeral: "IX"},
	{Value: 5, Numeral: "V"},
	{Value: 4, Numeral: "IV"},
	{Value: 1, Numeral: "I"},
}

func IntToRoman(num int) string {
	remainder := num
	res := ""

	for remainder > 0 {
		for _, RomNum := range RomanNumerals {
			if remainder >= RomNum.Value {
				remainder -= RomNum.Value
				res += RomNum.Numeral

				break
			}
		}
	}

	return res
}
