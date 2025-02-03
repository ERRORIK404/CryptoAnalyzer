package main

import (
	"bufio"
	"fmt"
	"os"

	analyzer "github.com/ERRORIK404/CryptoAnalyzer/pkg/application"
)

func main() {
	cryptoText := "ваша криптограмма здесь"
	analyzer := analyzer.NewCryptoAnalyzer(cryptoText)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nТекущий текст:", analyzer.DecryptedText)
		fmt.Println("1. Анализ частот")
		fmt.Println("2. Группировка слов по длине")
		fmt.Println("3. Группировка слов по неизвестным буквам")
		fmt.Println("4. Заменить букву")
		fmt.Println("5. Откат замены")
		fmt.Println("6. Автоматическая замена")
		fmt.Println("7. Выход")
		fmt.Print("Выберите действие: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			freq := analyzer.AnalyzeFrequency()
			fmt.Println("Частоты символов:")
			for char, count := range freq {
				fmt.Printf("%c: %d\n", char, count)
			}
			suggestions := analyzer.SuggestReplacements()
			fmt.Println("Предлагаемые замены:")
			for _, rep := range suggestions {
				fmt.Printf("%c -> %c\n", rep.From, rep.To)
			}
		case "2":
			groups := analyzer.GroupWordsByLength()
			fmt.Println("Слова, сгруппированные по длине:")
			for length, words := range groups {
				fmt.Printf("%d букв: %v\n", length, words)
			}
		case "3":
			groups := analyzer.GroupWordsByUnknownLetters()
			fmt.Println("Слова, сгруппированные по неизвестным буквам:")
			for unknown, words := range groups {
				fmt.Printf("%d неизвестных: %v\n", unknown, words)
			}
		case "4":
			fmt.Print("Введите букву для замены: ")
			scanner.Scan()
			oldChar := rune(scanner.Text()[0])
			fmt.Print("Введите новую букву: ")
			scanner.Scan()
			newChar := rune(scanner.Text()[0])
			analyzer.Replace(oldChar, newChar)
		case "5":
			analyzer.Undo()
			fmt.Println("Последняя замена отменена.")
		case "6":
			analyzer.AutoReplace()
			fmt.Println("Автоматическая замена выполнена.")
		case "7":
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}