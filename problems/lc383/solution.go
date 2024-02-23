package lc383

// 383. Ransom Note

// Given two strings ransomNote and magazine, return true
// if ransomNote can be constructed by using the letters from magazine
// and false otherwise.
//
// Each letter in magazine can only be used once in ransomNote.

func canConstruct(ransomNote string, magazine string) bool {
	noteLetters := []rune(ransomNote)
	magazineLetters := []rune(magazine)

	var letters = make(map[rune]int)
	for _, r := range noteLetters {
		if count, ok := letters[r]; ok {
			letters[r] = count + 1
		} else {
			letters[r] = 1
		}
	}
	for _, r := range magazineLetters {
		if count, ok := letters[r]; ok {
			count--
			if count > 0 {
				letters[r] = count
			} else {
				delete(letters, r)
			}
		}
	}

	return len(letters) == 0
}
