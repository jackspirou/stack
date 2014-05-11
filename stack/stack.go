// A stack built on a slice.
package stack

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack []interface{}

func NewStack() *Stack {
	return new(Stack)
}

func (s *Stack) Push(element interface{}) {
	*s = append(*s, element)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("Empty Stack.")
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return top, nil
}

func (s *Stack) Empty() bool {
	return len(*s) < 1
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) Length() int {
	return len(*s)
}

func (s *Stack) String() string {
	stack := ""
	if !s.Empty() {
		for key, value := range *s {
			stack += strconv.Itoa(key) + ":" + fmt.Sprintf("%v", value) + ", "
		}
		stack = stack[:len(stack)-1]
	}
	return "[" + stack + "]"
}
