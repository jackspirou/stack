package gostack

import (
	"errors"
	"fmt"
	"reflect"
)

type StackItem struct {
	Item interface{}
}

func NewStackItem(i interface{}) StackItem {
	return StackItem{
		Item: i,
	}
}

func (s StackItem) String() string {
	return fmt.Sprintf("%v ", s.Item)
}

type Stack struct {
	list []StackItem
}

func NewStack() *Stack {
	return &Stack{
		list: make([]StackItem, 0),
	}
}

func BuildStack(slice interface{}) *Stack {
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		panic("BuildStack only takes a slice.")
	}
	buff := reflect.ValueOf(slice)
	s := &Stack{
		list: make([]StackItem, buff.Len()),
	}
	for i := 0; i < buff.Len(); i++ {
		s.list[i] = NewStackItem(buff.Index(i).Interface())
	}
	return s
}

func (s *Stack) Push(element interface{}) {
	s.list = append(s.list, NewStackItem(element))
}

func (s *Stack) Pop() (StackItem, error) {
	if s.Empty() {
		return StackItem{}, errors.New("Empty Stack.")
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

func (s *Stack) Length() int {
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
