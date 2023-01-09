package data_structure

import "errors"

type Stack struct {
	Top    *Node
	length int
}

func NewStack() *Stack {
	return &Stack{
		length: 0,
	}
}

func (s *Stack) Push(data int) {
	item := Node{
		Data: data,
	}

	item.Next = s.Top
	s.Top = &item
	s.length++
}

func (s *Stack) Length() int {
	return s.length
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("the stack is empty")
	}

	return s.Top.Data, nil
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("the stack is empty")
	}

	item := s.Top
	s.Top = item.Next
	s.length--

	return item.Data, nil
}

func (s *Stack) IsEmpty() bool {
	return s.Top == nil
}
