package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	var tab = make(map[string]string)
	var keys []string

	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	markdownRegex := regexp.MustCompile(`\[([^][]+)]\((https?://[^()]+)\)`)
	results := markdownRegex.FindAllStringSubmatch(string(content), -1)
	for v := range results {
		title := strings.ReplaceAll(
			strings.ReplaceAll(
				results[v][1],
				"\n",
				""),
			"  ",
			" ")
		link := results[v][2]
		tab[title] = link
		keys = append(keys, title)
	}

	sort.Strings(keys)

	if len(os.Args) == 2 {
		for _, title := range keys {
			fmt.Println(title)
		}
	} else if len(os.Args) == 3 {
		var input string = os.Args[2]

		if input == "-" {
			reader := bufio.NewReader(os.Stdin)
			input, _ = reader.ReadString('\n')
			input = strings.TrimRight(input, "\n")
		}

		if v, exists := tab[input]; exists {
			fmt.Println(v)
		}
	}
}
