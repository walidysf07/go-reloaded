package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage : go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	text := string(data)

	text = processText(text)

	err = os.WriteFile(outputFile, []byte(text), 0644)
	if err != nil {
		fmt.Println("Error writing file")
		return
	}
}
func isPunctuation(s string) bool {
	for _, r := range s {
		if !strings.ContainsRune(".,!?;:", r) {
			return false
		}
	}
	return true
}

func processText(text string) string {

	text = strings.ReplaceAll(text, "(cap, ", "(cap,")
	text = strings.ReplaceAll(text, "(up, ", "(up,")
	text = strings.ReplaceAll(text, "(low, ", "(low,")
	text = strings.ReplaceAll(text, ").", ") .")
	text = strings.ReplaceAll(text, "),", ") ,")
	text = strings.ReplaceAll(text, ")!", ") !")
	text = strings.ReplaceAll(text, ")?", ") ?")
	text = strings.ReplaceAll(text, "):", ") :")
	text = strings.ReplaceAll(text, ");", ") ;")
	text = strings.ReplaceAll(text, "'!", "' !")
	text = strings.ReplaceAll(text, "'?", "' ?")
	text = strings.ReplaceAll(text, "'.", "' .")

	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {

		if words[i] == "(hex)" && i > 0 {
			number, err := strconv.ParseInt(words[i-1], 16, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(number, 10)
				words = append(words[:i], words[i+1:]...)
				i--
			}
		} else if words[i] == "(bin)" && i > 0 {
			number, err := strconv.ParseInt(words[i-1], 2, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(number, 10)
				words = append(words[:i], words[i+1:]...)
				i--
			}
		} else if words[i] == "(up)" && i > 0 {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		} else if words[i] == "(low)" && i > 0 {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		} else if words[i] == "(cap)" && i > 0 {
			words[i-1] = strings.ToUpper(words[i-1][:1]) + strings.ToLower(words[i-1][1:])
			words = append(words[:i], words[i+1:]...)
			i--
		} else if strings.HasPrefix(words[i], "(up,") {
			parts := strings.Split(words[i], ",")
			number := strings.TrimSuffix(parts[1], ")")

			n, err := strconv.Atoi(number)
			if err == nil {
				for j := 1; j <= n && i-j >= 0; j++ {
					words[i-j] = strings.ToUpper(words[i-j])
				}
			}

			words = append(words[:i], words[i+1:]...)
			i--
		} else if strings.HasPrefix(words[i], "(low,") {
			parts := strings.Split(words[i], ",")
			number := strings.TrimSuffix(parts[1], ")")

			n, err := strconv.Atoi(number)
			if err == nil {
				for j := 1; j <= n && i-j >= 0; j++ {
					words[i-j] = strings.ToLower(words[i-j])
				}

				words = append(words[:i], words[i+1:]...)
				i--
			}

		} else if strings.HasPrefix(words[i], "(cap,") {
			parts := strings.Split(words[i], ",")
			number := strings.TrimSuffix(parts[1], ")")

			n, err := strconv.Atoi(number)
			if err == nil {
				for j := 1; j <= n && i-j >= 0; j++ {
					words[i-j] = strings.ToUpper(words[i-j][:1]) + strings.ToLower(words[i-j][1:])
				}

				words = append(words[:i], words[i+1:]...)
				i--
			}
		}
	}

	for i := 0; i < len(words); i++ {
		if words[i] == "'" {
			for j := i + 1; j < len(words); j++ {
				if words[j] == "'" {
					words[i+1] = "'" + words[i+1]
					words[j-1] = words[j-1] + "'"

					words = append(words[:j], words[j+1:]...)
					words = append(words[:i], words[i+1:]...)
					i--
					break

				}
			}
		}
	}

	for i := 0; i < len(words); i++ {
		if isPunctuation(words[i]) && i > 0 {
			words[i-1] = words[i-1] + words[i]
			words = append(words[:i], words[i+1:]...)
			i--
			
		} else if len(words[i]) > 1 && words[i][len(words[i])-1] == '\'' {
			punctuation := words[i][:len(words[i])-1]

			if isPunctuation(punctuation) {
				words[i-1] = words[i-1] + punctuation + "'"
				words = append(words[:i], words[i+1:]...)
				i--
			}
		}else if len(words[i]) > 1 && isPunctuation(string(words[i][0])) && i > 0 {
			words[i-1] = words[i-1] + string(words[i][0])
			words[i] = words[i][1:]
	}

	for i := 0; i < len(words); i++ {
		if words[i] == "a" && i+1 < len(words) {

			first := string(words[i+1][0])
			if strings.Contains("aieouh", first) {
				words[i] = "an"
			}
		}
	}

	return strings.Join(words, " ")
}
