package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Check if correct number of arguments are provided
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input_file> <output_file>")
		return
	}

	// Read input file
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	inputData, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}
	modifiedText := string(inputData)
	data := strings.Split(modifiedText, " ")
	// Apply modifications
	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U', 'h', 'H'}
	consonants := []byte{'b', 'c', 'd', 'f', 'g', 'j', 'k', 'k', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z', 'B', 'C', 'D', 'F', 'G', 'J', 'K', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'V', 'W', 'X', 'Y', 'Z'}
	for m := 0; m < len(data)-1; m++ {
		for l := 0; l < len(consonants); l++ {
			if data[m] == "an" && []byte(data[m+1])[0] == consonants[l] {
				data[m] = "a"

			} else if data[m] == "An" && []byte(data[m+1])[0] == consonants[l] {
				data[m] = "A"
			}
		}
	}
	for m := 0; m < len(data)-1; m++ {
		for l := 0; l < len(vowels); l++ {
			if data[m] == "a" && []byte(data[m+1])[0] == vowels[l] {
				data[m] = "an"

			} else if data[m] == "A" && []byte(data[m+1])[0] == vowels[l] {
				data[m] = "An"

			}
		}
	}
	modifiedText1 := ""
	modifiedText2 := ""
	modifiedText = replaceHex(strings.Join(data, " "))
	modifiedText1 = replaceBin(modifiedText)
	modifiedText2 = Up_Cap_Low(strings.Split(modifiedText1, " "))
	slc_text := []byte(modifiedText2)
	pnct := []byte{',', ';', ':', '!', '?', '.'}
	bol := true
	for bol {
		bol = false
		for i := 1; i <= len(slc_text); i++ {
			for j := 0; j < len(pnct); j++ {
				if i < len(slc_text) && slc_text[i-1] == ' ' && slc_text[i] == pnct[j] {
					slc_text[i-1] = pnct[j]
					slc_text[i] = ' '
					bol = true
				}
			}
		}
	}
	// Expression régulière pour trouver deux espaces ou plus
	re := regexp.MustCompile(` {2,}`)
	// Remplacer les deux espaces ou plus par un seul espace
	newText := re.ReplaceAllString(string(slc_text), " ")
	// Expression régulière pour rechercher les single quotes suivis d'un espace
	cmpt := 0
	for i := 0; i < len(inputData); i++ {
		if inputData[i] == '\'' {
			cmpt++
		}
	}
	output1 := ""
	if cmpt%2 == 0 {
		re1 := regexp.MustCompile(`'\s*(.*?)\s*'`)
		output1 = re1.ReplaceAllString(newText, "'$1'")
	} else {
		re2 := regexp.MustCompile(`'\s*(.*?)\s*'\s*(.*?)\s*'`)
		output1 = re2.ReplaceAllString(newText, "'$1'$2'")
	}
	re3 := regexp.MustCompile(`"\s*(.*?)\s*"`)
	output := re3.ReplaceAllString(output1, "\"$1\"")
	err = ioutil.WriteFile(outputFile, []byte(output), 0644)
	if err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		return
	}
}

// replaceHex replaces hexadecimal numbers with their decimal equivalents.
func replaceHex(text string) string {
	re := regexp.MustCompile(`(\b[0-9a-fA-F]+\b) \(hex\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		hex := match[:len(match)-6] // Remove "(hex)" and one space
		decimal, _ := strconv.ParseInt(hex, 16, 64)
		return fmt.Sprintf("%d", decimal)
	})
}

// replaceBin replaces binary numbers with their decimal equivalents.
func replaceBin(text string) string {
	re := regexp.MustCompile(`(\b[01]+\b) \(bin\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		bin := match[:len(match)-6] // Remove "(bin)" and one space
		decimal, _ := strconv.ParseInt(bin, 2, 64)
		return fmt.Sprintf("%d", decimal)
	})
}

func Up_Cap_Low(data []string) string {
	for i := 0; i < len(data); i++ {
		switch data[i] {
		case "(cap)":
			if i > 0 {
				data[i-1] = strings.Title(data[i-1])
			}
			data[i] = ""
		case "(up)":
			if i > 0 {
				data[i-1] = strings.ToUpper(data[i-1])
			}
			data[i] = ""
		case "(low)":
			if i > 0 {
				data[i-1] = strings.ToLower(data[i-1])
			}
			data[i] = ""
		case "(cap,":
			if len(data) > 1 && i < len(data) -1 {
				if len(data[i+1]) > 1 {
				s := data[i+1]
				n := strings.TrimSuffix(s, ")")
				nbr, _ := strconv.Atoi(n)
				if nbr > 0 && nbr < 9223372036854775807 {
					if len(data[:i]) >= nbr {
					for l := 1; l <= nbr; l++ {
						data[i-l] = strings.Title(data[i-l])
						data[i+1] = ""
				        data[i] = ""
					}
				} else {
					fmt.Printf("the max words you can modifys is: %v\n", len(data[:i]))
				}
				} else {
					fmt.Println("error: you must enter a positive number!!")
				}
			}
			}
		case "(up,":
			if len(data) > 1 && i < len(data) - 1 {
				if len(data[i+1]) > 1 {
				s := data[i+1]
				n := strings.TrimSuffix(s, ")")
				nbr, _ := strconv.Atoi(n)
				if nbr > 0 && nbr < 9223372036854775807 {
					if len(data[:i]) >= nbr {
					for l := 1; l <= nbr; l++ {
						data[i-l] = strings.ToUpper(data[i-l])
						data[i+1] = ""
						data[i] = ""
					}
				} else {
					fmt.Printf("the max words you can modifys is: %v\n", len(data[:i]))
				}
				} else {
					fmt.Println("error: you must enter a positive number!!")
				}
			}
			}
		case "(low,":
			if len(data) > 1 && i < len(data) - 1 {
				if len(data[i+1]) > 1 {
				s := data[i+1]
				n := strings.TrimSuffix(s, ")")
				nbr, _ := strconv.Atoi(n)
				if nbr > 0 && nbr < 9223372036854775807 {
					if len(data[:i]) >= nbr {
					for l := 1; l <= nbr; l++ {
						data[i-l] = strings.ToLower(data[i-l])
						data[i+1] = ""
						data[i] = ""
					}
				} else {
					fmt.Printf("the max words you can modifys is: %v\n", len(data[:i]))
				}
				} else {
					fmt.Println("error: you must enter a positive number!!")
				}
			}
			}
		case "(hex)":
			data[i] = ""
		case "(bin)":
			data[i] = ""
		}
	}
	return strings.Join(data, " ")
}
