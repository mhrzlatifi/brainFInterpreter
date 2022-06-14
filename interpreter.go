package brainFInterpreter

import (
	"fmt"
)

type tape struct {
	possition   int
	memoryblock []byte
}

func newTape() *tape {
	tape := new(tape)
	tape.memoryblock = make([]byte, 4096)

	return tape
}

func (t tape) getMemoryBlockValue() byte {
	return t.memoryblock[t.possition]
}

func (t *tape) setValueToMemoryBlock(value byte) {
	t.memoryblock[t.possition] = value
}

func (t *tape) increment() {
	t.memoryblock[t.possition]++
}

func (t *tape) decrement() {
	t.memoryblock[t.possition]--
}

func (t *tape) moveForward() {
	t.possition++

	if len(t.memoryblock) <= t.possition {
		t.memoryblock = append(t.memoryblock, 0)
	}
}

func (t *tape) moveBackward() {
	t.possition--
}

func parse(str string) (string, map[int]int) {

	parsed := make([]byte, 0)
	stack := make([]int, 0)
	loopMap := make(map[int]int, 128)

	for i, char := range str {

		switch char {
		case '>', '<', '+', '-', '.', ',', '[', ']':
			parsed = append(parsed, byte(char))
			if char == '[' {
				stack = append(stack, i)
			} else if char == ']' {
				last := len(stack) - 1
				left := stack[last]
				stack = stack[:last]
				right := i
				loopMap[right] = left
				loopMap[left] = right
			}
		}

	}
	return string(parsed), loopMap

}

func TranslateThis(str string) {

	parsedStr, loopMap := parse(str)

	index := 0
	tape := newTape()
	var value byte

	for index < len(parsedStr) {

		switch parsedStr[index] {

		case '>':
			tape.moveForward()
		case '<':
			tape.moveBackward()
		case '+':
			tape.increment()
		case '-':
			tape.decrement()
		case '.':
			fmt.Printf("%c", tape.getMemoryBlockValue())
		case ',':
			fmt.Scanf("%c", &value)
			tape.setValueToMemoryBlock(value)
		case '[':
			if tape.getMemoryBlockValue() == 0 {
				index = loopMap[index]
			}
		case ']':
			if tape.getMemoryBlockValue() != 0 {
				index = loopMap[index]
			}

		}
		index++

	}

}
