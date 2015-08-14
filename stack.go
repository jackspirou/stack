package stack

import (
	"errors"
	"fmt"
	"reflect"
)

// Stack represents a golang implimentation of the abstract stack data type
// that accepts an empty interface to ensure struct compabilility.
type Stack struct {
	list []interface{}
}

func New() *Stack {
	return &Stack{}
}

func Build(slice interface{}) *Stack {
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		panic("BuildStack only takes a slice.")
	}
	buff := reflect.ValueOf(slice)
	s := &Stack{
		list: make([]interface{}, buff.Len()),
	}
	for i := 0; i < buff.Len(); i++ {
		s.list[i] = buff.Index(i).Interface()
	}
	return s
}

func (s *Stack) Push(item interface{}) {
	s.list = append(s.list, item)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("Stack Error: empty stack")
	}
	top := s.list[len(s.list)-1]
	s.list = s.list[0 : len(s.list)-1]
	return top, nil
}

func (s *Stack) Empty() bool {
	return len(s.list) < 1
}

func (s *Stack) Size() int {
	return len(s.list)
}

func (s *Stack) String() string {
	stack := ""
	if !s.Empty() {
		for _, value := range s.list {
			stack += fmt.Sprintf("%v ", value)
		}
		stack = stack[:len(stack)-2]
	}
	return "[" + stack + "]"
}
