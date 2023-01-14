package stack_min

import "errors"

type Stack struct {
	items   []int
	minimum []int
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) push(data int) {
	s.items = append(s.items, data)

	last := 0
	length := len(s.minimum)

	if length > 0 {
		last = s.minimum[length-1]

		if last < data {
			s.minimum = append(s.minimum, last)
			return
		}
	}

	s.minimum = append(s.minimum, data)
}

func (s *Stack) pop() (int, error) {
	data, err := s.peek()

	if err != nil {
		return 0, err
	}

	s.items = s.items[0 : len(s.items)-1]
	s.minimum = s.minimum[0 : len(s.minimum)-1]

	return data, nil
}

func (s *Stack) peek() (int, error) {
	length := len(s.items)

	if length > 0 {
		return s.items[length-1], nil
	}

	return 0, errors.New("the stack is empty")
}

func (s *Stack) min() (int, error) {
	length := len(s.minimum)

	if length > 0 {
		return s.minimum[length-1], nil
	}

	return 0, errors.New("the stack is empty")
}
