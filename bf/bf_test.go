package bf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testsReturnInt struct {
	source string
	want   int
}

type testsReturnRune struct {
	source string
	want   rune
}

type testsReturnRunes struct {
	source    string
	want      []rune
	inputData string
	err       error
}

func TestOperatorNext(t *testing.T) {
	tests := []testsReturnInt{
		{source: ">", want: 1},
		{source: ">>>>", want: 4},
		{source: ">>", want: 2},
		{source: ">>>>>", want: 0},
	}

	for _, c := range tests {
		vm := New(c.source, 5, "")
		err := vm.Run()
		assert.Nil(t, err)
		assert.Equal(t, c.want, vm.pos)
	}
}

func TestOperatorPrevious(t *testing.T) {
	tests := []testsReturnInt{
		{source: "<", want: 4},
		{source: "<<<<<", want: 0},
		{source: "<<", want: 3},
		{source: "<<<<<<<<", want: 2},
	}

	for _, c := range tests {
		vm := New(c.source, 5, "")
		err := vm.Run()
		assert.Nil(t, err)
		assert.Equal(t, c.want, vm.pos)
	}
}

func TestOperatorAdd1(t *testing.T) {
	tests := []testsReturnRune{
		{source: "+", want: 1},
		{source: "++++++", want: 6},
		{source: "++++++++++++++++++++", want: 20},
	}

	for _, c := range tests {
		vm := New(c.source, 5, "")
		err := vm.Run()
		assert.Nil(t, err)
		assert.Equal(t, c.want, vm.memory[0])
	}
}

func TestOperatorSub1(t *testing.T) {
	tests := []testsReturnRune{
		{source: "-", want: -1},
		{source: "------", want: -6},
		{source: "--------------------", want: -20},
	}

	for _, c := range tests {
		vm := New(c.source, 5, "")
		err := vm.Run()
		assert.Nil(t, err)
		assert.Equal(t, c.want, vm.memory[0])
	}
}

func TestOperatorShow(t *testing.T) {
	tests := []testsReturnRunes{
		{source: "+++++++++++++++++++++++++++++++++.", want: []rune{33}},
		{source: "++++++++++++++++++++++++++++++++.+++++++++++++++++++++++++++++++++.", want: []rune{' ', 'A'}},
		{source: "+++++++++++++++++++++++++++++++++.>+++++++++++++++++++++++++++++++++.", want: []rune{'!', '!'}},
		{source: "++++++++++++++++++++++++++++++++.+.+..", want: []rune{' ', '!', '"', '"'}},
	}

	for _, c := range tests {
		vm := New(c.source, 5, "")
		err := vm.Run()
		assert.Nil(t, err)
		assert.Equal(t, c.want, vm.Output)
	}
}

func TestOperatorOpenAndClose(t *testing.T) {
	tests := []testsReturnRunes{
		{source: "+++.[-.].", want: []rune{3, 2, 1, 0, 0}},
		{source: "-----.[+.].", want: []rune{-5, -4, -3, -2, -1, 0, 0}},
		{source: "[].", want: []rune{0}},
	}

	for _, c := range tests {
		vm := New(c.source, 5, "")
		err := vm.Run()
		assert.Nil(t, err)
		assert.Equal(t, c.want, vm.Output)
	}
}

func TestOperatorInput(t *testing.T) {
	tests := []testsReturnRunes{
		{source: ",", want: nil, err: ErrorNoInputData},
		{source: ",.", want: []rune{'b'}, err: nil, inputData: "b"},
	}

	for _, c := range tests {
		vm := New(c.source, 5, c.inputData)
		err := vm.Run()
		assert.Equal(t, c.err, err)
		assert.Equal(t, c.want, vm.Output)
	}
}

func TestRunFunction(t *testing.T) {
	tests := []testsReturnRunes{
		{ // adapted from https://en.wikipedia.org/wiki/Brainfuck#Hello_World!
			source: "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.",
			want:   []rune{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '!'},
			err:    nil,
		},
		// The following tests are collect from http://esoteric.sange.fi/brainfuck/bf-source/prog/tests.b
		{ // "esoteric array is big"
			source: "++++[>++++++<-]>[>+++++>+++++++<<-]>>++++<[[>[[>>+<<-]<]>>>-]>-[>+>+<<-]>]+++++[>+++++++<<++>-]>.<<.",
			want:   []rune{'#', '\n'},
			err:    nil,
		},
		{ // "esoteric obscure errors"
			source: "[]++++++++++[>++++++++++++++++++>+++++++>+<<<-]A;?@![#>>+<<]>[>++<[-]]>.>.",
			want:   []rune{'H', '\n'},
			err:    nil,
		},
		{ // esoteric new line input
			source:    ">,>+++++++++,>+++++++++++[<++++++<++++++<+>>>-]<<.>.<<-.>.>.<<.",
			want:      []rune{'L', 'F', '\n', 'L', 'F', '\n'},
			inputData: string([]rune{'\n', 4}),
			err:       nil,
		},
	}

	for _, c := range tests {
		vm := New(c.source, 5000, c.inputData)
		err := vm.Run()
		assert.Equal(t, c.err, err)
		assert.Equal(t, c.want, vm.Output)
	}
}
