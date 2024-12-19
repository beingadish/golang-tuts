package main

import "fmt"

// ====================== GENERICS EXPLAINATION USING GENERIC STACK IMPLEMENTATION ===============================

// Only String & Int Type Slices are allowed
func printSlice[T string | int](nums ...T) {
	for _, val := range nums {
		fmt.Println(val)
	}
}

// Stack of Any Type can be made using this Generic Stack struct

type stack[T interface{}] struct {
	index    int
	elements []T
}

func (s *stack[T]) push(element T) {
	if len(s.elements) == 0 {
		s.index = -1
	}
	s.elements = append(s.elements, element)
	s.index += 1
	s.elements[s.index] = element
}

func (s *stack[T]) pop() {
	s.elements = s.elements[0:s.index]
	s.index -= 1
	if s.index == -1 {
		s.elements = []T{}
	}
}

func (s *stack[T]) top() T {
	if s.index == -1 {
		panic("Stack Undeflow")
	} else {
		return s.elements[s.index]
	}
}

func main() {
	// Golang Tuts on Generics

	myStack := stack[int]{}

	myStack.push(5)
	myStack.push(1)
	myStack.push(7)
	myStack.pop()
	myStack.pop()
	myStack.pop()
	myStack.push(100)
	fmt.Println(myStack.top())

}
