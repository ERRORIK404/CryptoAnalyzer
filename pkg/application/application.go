package application

import (
	"sort"
	"strings"
)

// Частоты букв русского алфавита
var russianFreq = map[rune]float64{
	'о': 0.1097, 'е': 0.0845, 'а': 0.0801, 'и': 0.0735, 'н': 0.0670,
	'т': 0.0626, 'с': 0.0547, 'р': 0.0473, 'в': 0.0454, 'л': 0.0440,
	'к': 0.0349, 'м': 0.0321, 'д': 0.0298, 'п': 0.0281, 'у': 0.0262,
	'я': 0.0201, 'ы': 0.0190, 'ь': 0.0174, 'г': 0.0170, 'з': 0.0165,
	'б': 0.0159, 'ч': 0.0144, 'й': 0.0121, 'х': 0.0097, 'ж': 0.0094,
	'ш': 0.0073, 'ю': 0.0064, 'ц': 0.0048, 'щ': 0.0036, 'э': 0.0032,
	'ф': 0.0026, 'ъ': 0.0004,
}

type Replacement struct {
	From rune
	To   rune
}

type CryptoAnalyzer struct {
	CryptoText    string
	DecryptedText string
	Replacements  []Replacement
	History       [][]Replacement
}

func NewCryptoAnalyzer(cryptoText string) *CryptoAnalyzer {
	return &CryptoAnalyzer{
		CryptoText:    cryptoText,
		DecryptedText: cryptoText,
		Replacements:  []Replacement{},
		History:       [][]Replacement{},
	}
}

func (ca *CryptoAnalyzer) AnalyzeFrequency() map[rune]int {
	freq := make(map[rune]int)
	for _, char := range ca.CryptoText {
		if char >= 'а' && char <= 'я' || char >= 'А' && char <= 'Я' {
			freq[char]++
		}
	}
	return freq
}

func (ca *CryptoAnalyzer) SuggestReplacements() []Replacement {
	cryptoFreq := ca.AnalyzeFrequency()
	sortedCrypto := sortByFrequency(cryptoFreq)
	sortedRussian := sortRussianFreq()

	suggestions := []Replacement{}
	for i := 0; i < len(sortedCrypto) && i < len(sortedRussian); i++ {
		suggestions = append(suggestions, Replacement{
			From: sortedCrypto[i].Key,
			To:   sortedRussian[i].Key,
		})
	}
	return suggestions
}

func (ca *CryptoAnalyzer) GroupWordsByLength() map[int][]string {
	words := strings.Fields(ca.CryptoText)
	groups := make(map[int][]string)
	for _, word := range words {
		length := len(word)
		groups[length] = append(groups[length], word)
	}
	return groups
}

func (ca *CryptoAnalyzer) GroupWordsByUnknownLetters() map[int][]string {
	words := strings.Fields(ca.DecryptedText)
	groups := make(map[int][]string)
	for _, word := range words {
		unknown := 0
		for _, char := range word {
			if char == '*' {
				unknown++
			}
		}
		groups[unknown] = append(groups[unknown], word)
	}
	return groups
}

func (ca *CryptoAnalyzer) Replace(oldChar, newChar rune) {
	ca.History = append(ca.History, append([]Replacement(nil), ca.Replacements...))
	ca.Replacements = append(ca.Replacements, Replacement{From: oldChar, To: newChar})
	ca.updateDecryptedText()
}

func (ca *CryptoAnalyzer) Undo() {
	if len(ca.History) > 0 {
		ca.Replacements = ca.History[len(ca.History)-1]
		ca.History = ca.History[:len(ca.History)-1]
		ca.updateDecryptedText()
	}
}

func (ca *CryptoAnalyzer) updateDecryptedText() {
	decrypted := []rune(ca.CryptoText)
	for _, rep := range ca.Replacements {
		for i, char := range decrypted {
			if char == rep.From {
				decrypted[i] = rep.To
			}
		}
	}
	ca.DecryptedText = string(decrypted)
}

func (ca *CryptoAnalyzer) AutoReplace() {
	suggestions := ca.SuggestReplacements()
	for _, suggestion := range suggestions {
		ca.Replace(suggestion.From, suggestion.To)
	}
}

func sortByFrequency(freq map[rune]int) []struct {
	Key   rune
	Value int
} {
	var sorted []struct {
		Key   rune
		Value int
	}
	for k, v := range freq {
		sorted = append(sorted, struct {
			Key   rune
			Value int
		}{k, v})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})
	return sorted
}

func sortRussianFreq() []struct {
	Key   rune
	Value float64
} {
	var sorted []struct {
		Key   rune
		Value float64
	}
	for k, v := range russianFreq {
		sorted = append(sorted, struct {
			Key   rune
			Value float64
		}{k, v})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})
	return sorted
}
