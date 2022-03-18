package count_vowel_substrings

func isVowel(char string) bool {
	return char == "a" || char == "e" || char == "i" || char == "o" || char == "u"
}

func countVowelSubstrings(word string) int {
	var j, k, vow, count int
	mapVowel := make(map[string]int)
	for i := 0; i < len(word); i++ {
		char := string(word[i])
		if isVowel(char) {

			mapVowel[char]++
			if mapVowel[char] == 1 {
				vow++
			}
			for vow == 5 {
				mapVowel[string(word[k])]--
				if mapVowel[string(word[k])] == 0 {
					vow--
				}
				k++
			}
			count += k - j
		} else {
			for k, _ := range mapVowel {
				mapVowel[k] = 0
			}
			vow = 0
			k = i + 1
			j = k
		}

	}
	return count

}
