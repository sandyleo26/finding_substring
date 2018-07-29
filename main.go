package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

const NotFound = "<No Output>"

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: main text_to_search subtext")
		return
	}

	fmt.Println(Search(os.Args[1], os.Args[2]))
}

//Search finds occurrences of subtext in textToSearch
func Search(textToSearch, subtext string) string {
	if len(textToSearch) == 0 || len(subtext) == 0 || len(textToSearch) < len(subtext) {
		return NotFound
	}

	lowerTextToSearch := ToLowerCase(textToSearch)
	lowerSubtext := ToLowerCase(subtext)

	results := make([]string, 0)
	for i := 0; i <= len(lowerTextToSearch)-len(lowerSubtext); i++ {
		found := true
		for j := 0; j < len(lowerSubtext); j++ {
			if lowerTextToSearch[i+j] != lowerSubtext[j] {
				found = false
				break
			}
		}

		if found {
			results = append(results, strconv.Itoa(i+1))
			i += len(lowerSubtext)
		}
	}

	if len(results) == 0 {
		return NotFound
	}

	if len(results) == 1 {
		return results[0]
	}

	resStr := results[0]
	for i, s := range results {
		if i == 0 {
			continue
		}
		resStr += ", " + s
	}

	return resStr
}

//ToLowerCase converts any occurrences of A-Z to a-z
func ToLowerCase(s string) string {
	var buffer bytes.Buffer
	for _, r := range s {
		if 'A' <= r && r <= 'Z' {

			buffer.WriteRune(r + 32)
			continue
		}
		buffer.WriteRune(r)
	}

	return buffer.String()
}
