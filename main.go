package main

import (
    "bufio"
    "fmt"
    "os"
	"regexp"
	"sort"
)

func unique(s []string) []string {
    inResult := make(map[string]bool)
    var result []string
    for _, str := range s {
        if _, ok := inResult[str]; !ok {
            inResult[str] = true
            result = append(result, str)
        }
    }
    return result
}

func main() {
    var lines []string

    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        line := scanner.Text()
        lines = append(lines, line)
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "Error reading input:", err)
        os.Exit(1)
    }
	regex := regexp.MustCompile("(<=&|=)([^&=\n]+)|(<=&|=).*\n")

	for i, j := range(lines){
		lines[i] = regex.ReplaceAllString(j,"=")
	}
	sort.Strings(lines)
	lines = unique(lines)
	for _,i := range(lines){
		fmt.Println(i)
	}
}