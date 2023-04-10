package bf

import "errors"

var ErrorNoInputData = errors.New("no more input data to fetch")

type VirtualMachine struct {
	source      []rune
	memory      []rune
	pos         int
	instruction int
	Output      []rune
	inputData   []rune
}

// New returns a new VirtualMachine pointer.
//
//		s: source code
//	 	m: memory size
//		i: input data
func New(s string, m int, i string) *VirtualMachine {
	return &VirtualMachine{
		source:    []rune(s),
		memory:    make([]rune, m),
		inputData: []rune(i),
	}
}

func (v *VirtualMachine) Run() error {
	for {
		token := v.source[v.instruction]
		switch token {
		case '>':
			v.next()
		case '<':
			v.previous()
		case '+':
			v.add1()
		case '-':
			v.sub1()
		case '.':
			v.show()
		case '[':
			v.open()
		case ']':
			v.close()
		case ',':
			err := v.input()
			if err != nil {
				return err
			}
		}

		if v.pos < 0 {
			v.pos = len(v.memory) - 1
		} else if v.pos >= len(v.memory) {
			v.pos = 0
		}

		v.instruction++
		if v.instruction >= len(v.source) {
			break
		}
	}

	return nil
}

func (v *VirtualMachine) next() {
	v.pos++
}

func (v *VirtualMachine) previous() {
	v.pos--
}

func (v *VirtualMachine) add1() {
	v.memory[v.pos]++
}

func (v *VirtualMachine) sub1() {
	v.memory[v.pos]--
}

func (v *VirtualMachine) show() {
	v.Output = append(v.Output, rune(v.memory[v.pos]))
}

func (v *VirtualMachine) open() {
	if v.memory[v.pos] != 0 {
		return
	}

	for depth := 1; depth > 0; {
		v.instruction++

		instruction := v.source[v.instruction]
		if instruction == '[' {
			depth++
		} else if instruction == ']' {
			depth--
		}
	}
}

func (v *VirtualMachine) close() {
	for depth := 1; depth > 0; {
		v.instruction--
		instruction := v.source[v.instruction]
		if instruction == '[' {
			depth--
		} else if instruction == ']' {
			depth++
		}
	}
	v.instruction--
}

func (v *VirtualMachine) input() error {
	if len(v.inputData) < 1 {
		return ErrorNoInputData
	}

	element := v.inputData[0]
	v.inputData = v.inputData[1:]
	v.memory[v.pos] = element
	return nil
}
