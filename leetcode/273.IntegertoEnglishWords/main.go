package main

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	lessThan20 := []string{
		"",
		"One",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Eleven",
		"Twelve",
		"Thirteen",
		"Fourteen",
		"Fifteen",
		"Sixteen",
		"Seventeen",
		"Eighteen",
		"Nineteen",
	}

	tensWords := []string{
		"",
		"Ten",
		"Twenty",
		"Thirty",
		"Forty",
		"Fifty",
		"Sixty",
		"Seventy",
		"Eighty",
		"Ninety",
	}

	scaleWords := []string{"Billion", "Million", "Thousand", ""}
	scaleIndex := 0
	scaleDivisor := 1000000000 // billion

	var transfer func(int) string
	transfer = func(i int) string {
		tmp := ""
		if i < 20 {
			tmp += lessThan20[i]
		} else if i < 100 {
			currentGroup := i / 10
			remaining := i % 10
			tmp += tensWords[currentGroup]
			if remaining > 0 {
				tmp += " " + lessThan20[remaining]
			}
		} else {
			currentGroup := i / 100
			remaining := i % 100
			tmp += lessThan20[currentGroup] + " Hundred"
			if remaining > 0 {
				tmp += " " + transfer(remaining)
			}
		}

		return tmp
	}

	res := ""
	for scaleDivisor > 0 {
		currentGroup := num / scaleDivisor
		if currentGroup > 0 {
			if res == "" {
				res += transfer(currentGroup)
			} else {
				res += " " + transfer(currentGroup)

			}
			if scaleWords[scaleIndex] != "" {
				res += " " + scaleWords[scaleIndex]
			}

		}
		num = num % scaleDivisor
		scaleDivisor = scaleDivisor / 1000
		scaleIndex++
	}
	return res

}
