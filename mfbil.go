package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Language struct {
	source []rune
	tokens map[rune]func(*Language)
	memory [100]int
	pos    int
}

func (l *Language) Execute() {
	for _, token := range l.source {
		if function, ok := l.tokens[token]; ok {
			function(l)
		} else {
			log.Fatalf("Invalid token [%c]\n", token)
		}
	}
}

func (l *Language) AddToken(t rune, f func(*Language)) {
	if l.tokens == nil {
		l.tokens = make(map[rune]func(*Language))
	}

	l.tokens[t] = f
}

func (l *Language) AddToSource(line string) {
	for _, c := range line {
		if _, ok := l.tokens[c]; ok {
			l.source = append(l.source, c)
		}
	}
}

func search(element int, array []int) int {
	var index = -1

	for i, v := range array {
		if element == v {
			index = i
			break
		}
	}

	return index
}

func open_bracket(l *Language) {
	if l.pos == 0 {
		l.pos = search(']', l.memory[:]) + 1
	}
}

func close_bracket(l *Language) {
	if l.pos != 0 {
		l.pos = search('[', l.memory[:]) + 1
	}
}

func main() {
	var program Language

	var tokens = map[rune]func(*Language){
		'>': func(l *Language) { l.pos++ },
		'<': func(l *Language) { l.pos-- },
		'+': func(l *Language) { l.memory[l.pos]++ },
		'-': func(l *Language) { l.memory[l.pos]-- },
		'.': func(l *Language) { fmt.Println(l.memory[l.pos]) },
		'[': func(l *Language) { open_bracket(l) },
		']': func(l *Language) { close_bracket(l) },
	}

	for t, f := range tokens {
		(&program).AddToken(t, f)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		(&program).AddToSource(line)
	}

	(&program).Execute()
}
