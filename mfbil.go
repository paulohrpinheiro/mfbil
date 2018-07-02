package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Language struct {
	source      []rune
	tokens      map[rune]func(*Language)
	memory      [100]int
	loops       []int
	pos         int
	instruction int
}

func (l *Language) Execute() {
	for {
		var token = l.source[l.instruction]

		if function, ok := l.tokens[token]; ok {
			function(l)
		} else {
			log.Fatalf("Invalid token [%c]\n", token)
		}

		if l.pos < 0 {
			l.pos = len(l.memory) - 1
		} else if l.pos >= len(l.memory) {
			l.pos = 0
		}

		l.instruction++
		if l.instruction >= len(l.source) {
			break
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

func search(element rune, array []rune) int {
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
	if l.memory[l.pos] != 0 {
		return
	}

	for depth := 1; depth > 0; {
		l.instruction++

		instruction := l.source[l.instruction]
		if instruction == '[' {
			depth++
		} else if instruction == ']' {
			depth--
		}
	}

}

func close_bracket(l *Language) {
	for depth := 1; depth > 0; {
		l.instruction--
		instruction := l.source[l.instruction]
		if instruction == '[' {
			depth--
		} else if instruction == ']' {
			depth++
		}
	}
	l.instruction--
}

func main() {
	var program Language

	var tokens = map[rune]func(*Language){
		'>': func(l *Language) { l.pos++ },
		'<': func(l *Language) { l.pos-- },
		'+': func(l *Language) { l.memory[l.pos]++ },
		'-': func(l *Language) { l.memory[l.pos]-- },
		'.': func(l *Language) { fmt.Print(string(l.memory[l.pos])) },
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
