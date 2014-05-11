package test

import (
	"fmt"
	"github.com/jackspirou/gostack/stack"
	"math/rand"
	"strconv"
)

func Ints() {
	s := stack.NewStack()
	var randomInts [20]int
	for key := range randomInts {
		randomInts[key] = rand.Intn(len(randomInts))
	}
	for key, value := range randomInts {
		s.Push(randomInts[key])
		fmt.Println(s.String() + " Pushed: " + strconv.Itoa(value))
	}
	for !s.Empty() {
		element, err := s.Pop()
		if err != nil {
			panic(err)
		}
		fmt.Println(s.String() + " Popped: " + strconv.Itoa(element.(int)))
	}

}
