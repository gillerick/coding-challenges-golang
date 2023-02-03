package main

type EmptyStackException struct{}

func (e EmptyStackException) Error() string {
	return "stack is empty"
}

type ArrayStack struct {
	data []interface{}
}

func (s *ArrayStack) Len() int {
	return len(s.data)
}

func (s *ArrayStack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *ArrayStack) Push(element interface{}) {
	s.data = append(s.data, element)
}

func (s *ArrayStack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, EmptyStackException{}
	}
	return s.data[len(s.data)-1], nil
}

func (s *ArrayStack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, EmptyStackException{}
	}

	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top, nil

}
