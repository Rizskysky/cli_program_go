package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func lowerFirst(s string) string {
	a := []rune(s)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}

func stringConvertion(input string) string {
	lorem := strings.Fields(input)
	result := lorem
	for i := 0; i < len(lorem); i++ {
		word := strings.ToUpper(lorem[i])
		word = lowerFirst(word)
		result[i] = word
	}
	return strings.Join((result), " ")
}

func unique(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		temp := strings.ToLower(entry)
		if _, value := keys[temp]; !value {
			keys[temp] = true
			list = append(list, entry)
		}
	}
	return list
}

func letterCounter(inputKedua string) string {
	var parsed string
	reg, err := regexp.Compile(`[^a-zA-Z0-9]+`)
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(inputKedua, "")
	resUnique := unique(strings.Split(processedString, ""))
	list := []string{}
	for i := 0; i < len(resUnique); i++ {
		counter := strings.Count(strings.ToLower(processedString), strings.ToLower(resUnique[i]))
		parsed = strconv.Itoa(counter)
		if parsed == "1" {
			parsed = ""
		}
		list = append(list, fmt.Sprintf("%s%s", resUnique[i], parsed))
	}
	return strings.Join(list, "")
}

func wordstatistic(input string) {
	counter := 0
	showedOnce := 0
	max := 0
	min := 1
	var maxWord, minWord string

	reg, err := regexp.Compile(`[^a-zA-Z0-9]+`)
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(input, " ")
	arrayStringProceed := strings.Fields(strings.ToUpper(processedString))
	counts := make(map[string]int)
	for _, word := range arrayStringProceed {
		counter++
		_, ok := counts[word]
		if ok {
			counts[word]++
		} else {
			counts[word] = 1
		}
	}

	for keys, value := range counts {
		if value == 1 {
			showedOnce++
		}
		if max <= value {
			max = value
			maxWord = keys
		}

		if min >= value {
			min = value
			minWord = keys
		}

		keyModified := strings.Title(strings.ToLower(keys))
		fmt.Printf("word : %s = %d \n", keyModified, value)
	}
	fmt.Printf("+++++++++++++++++++++++++++++\nTotal Words : %d\nTotal Words just showed once : %d\n", counter, showedOnce)
	fmt.Printf("MVP highest count [%d] word : %s \n", max, maxWord)
	fmt.Printf("MVP lowest count  [%d] word : %s ", min, minWord)

}

func command() string {
	args := os.Args[1:]

	if len(args) > 0 {
		return args[0]
	}
	return ""
}

func main() {
	cmdstring := command()

	input := "Lorem Ipsum Dolor Sit Amet babang Biji"
	inputKedua := "Team Engineering PT. Raksasa Laju Lintang "
	inputDewa := ` Go , also known as Golang , [14] is a statically typed , compiled programming language designed at
	Google [15] by Robert Griesemer, Rob Pike , and Ken Thompson . [12] Go is syntactically similar to C ,
	but with memory safety , garbage collection , structural typing , [6] and CSP -style concurrency . [16]
	Go was designed at Google in 2007 to improve programming productivity in an era of multicore ,
	networked machines and large codebases . [23] The designers wanted to address criticism of other
	languages in use at Google , but keep their useful characteristics: [24]
	● Static typing and run-time efficiency (like C++ )
	● Readability and usability (like Python or JavaScript ) [25]
	● High-performance networking and multiprocessing
	The designers were primarily motivated by their shared dislike of C++ . [26] [27] [28]
	Go was publicly announced in November 2009, [29] and version 1.0 was released in March
	2012. [30] [31] Go is widely used in production at Google [32] and in many other organizations and
	open-source projects.
	In November 2016, the Go and Go Mono fonts which are sans-serif and monospaced respectively
	were released by type designers Charles Bigelow and Kris Holmes . Both fonts adhere to WGL4 and
	were designed to be legible with a large x-height and distinct letterforms by conforming to the DIN
	1450 standard. [33] [34]
	In April 2018, the original logo was replaced with a stylized GO slanting right with trailing
	streamlines. However, the Gopher mascot remained the same. [35]
	In August 2018, the Go principal contributors published two ″draft designs″ for new language
	features, Generics and error handling , and asked Go users to submit feedback on them. [36] [37] Lack
	of support for generic programming and the verbosity of error handling in Go 1.x had drawn
	considerable criticism . `

	switch cmdstring {
	case "soalPertama":
		result := stringConvertion(input)
		fmt.Println(result)

	case "soalKedua":
		result := letterCounter(inputKedua)
		fmt.Println(result)

	case "soalKetiga":
		wordstatistic(inputDewa)
	}

}
