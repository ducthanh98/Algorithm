package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	res := 0

	if !Include(wordList, endWord) {
		return 0
	}

	wordList = append(wordList, beginWord)

	hash := make(map[string][]string)
	for _, word := range wordList {

		for i, _ := range word {
			pattern := word[:i] + "#" + word[i+1:]
			hash[pattern] = append(hash[pattern], word)
		}

	}
	visited := make(map[string]bool)
	queue := make([]string, 0)
	queue = append(queue, beginWord)

	for len(queue) > 0 {
		res++
		size := len(queue)
		for k := 0; k < size; k++ {
			word := queue[0]
			queue = queue[1:]
			if word == endWord {
				return res
			} else {

				for i, _ := range word {
					pattern := word[:i] + "#" + word[i+1:]

					for _, s := range hash[pattern] {
						if !visited[s] {
							visited[s] = true
							queue = append(queue, s)
						}
					}

				}

			}
		}

	}

	return 0
}

func Include(w []string, s string) bool {
	for _, s2 := range w {
		if s2 == s {
			return true
		}
	}
	return false
}
