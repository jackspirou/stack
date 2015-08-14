package stack

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func Test_Stack_Of_Ints(t *testing.T) {

	fmt.Println("")
	fmt.Println("/*")
	fmt.Println("--------------------------")
	fmt.Println("---- TESTING INTEGERS ----")
	fmt.Println("--------------------------")
	fmt.Println("*/")
	fmt.Println("")
	fmt.Println("")

	// Generate sudo random integers
	fmt.Println("// Generating Sudo Random Ints")
	var randomInts [20]int
	for key := range randomInts {
		randomInts[key] = rand.Intn(len(randomInts))
	}

	// Test NewStack
	fmt.Println("// int - NewStack()")
	newStack := NewStack()
	for key, value := range randomInts {
		newStack.Push(randomInts[key])
		fmt.Println(newStack.String() + " Pushed: " + strconv.Itoa(value))
	}
	for !newStack.Empty() {
		element, err := newStack.Pop()
		if err != nil {
			t.Errorf("Int Stack Error:" + err.Error())
		}
		fmt.Println(newStack.String() + " Popped: " + element.String())
	}

	// Test var stack.Stack
	fmt.Println("// int - stack.Stack")
	var vStack Stack
	for key, value := range randomInts {
		vStack.Push(randomInts[key])
		fmt.Println(vStack.String() + " Pushed: " + strconv.Itoa(value))
	}
	for !vStack.Empty() {
		element, err := vStack.Pop()
		if err != nil {
			t.Errorf("Int Stack Error:" + err.Error())
		}
		fmt.Println(vStack.String() + " Popped: " + element.String())
	}

	// Test BuildStack
	fmt.Println("// int - BuildStack")
	dataSlice := []int{1, 2, 3, 4, 5, 6, 43, 423}
	buildStack := BuildStack(dataSlice)
	for key, value := range randomInts {
		buildStack.Push(randomInts[key])
		fmt.Println(buildStack.String() + " Pushed: " + strconv.Itoa(value))
	}
	for !buildStack.Empty() {
		element, err := buildStack.Pop()
		if err != nil {
			t.Errorf("Int Stack Error:" + err.Error())
		}
		fmt.Println(buildStack.String() + " Popped: " + element.String())
	}
}

func Test_Stack_Of_Strings(t *testing.T) {

	fmt.Println("")
	fmt.Println("/*")
	fmt.Println("-------------------------")
	fmt.Println("---- TESTING STRINGS ----")
	fmt.Println("-------------------------")
	fmt.Println("*/")
	fmt.Println("")
	fmt.Println("")

	// Generate sudo random strings
	fmt.Println("// Generating Strings")
	randomStrings := [20]string{
		"disadvantageous",
		"sarthe",
		"nonsatiable",
		"marguerite",
		"rediscussion",
		"fiducial",
		"unmagnetic",
		"respecter",
		"amphoricity",
		"floreated",
		"heathenising",
		"dextro",
		"sartorius",
		"hydrarch",
		"vaporizable",
		"baclava",
		"reclothe",
		"entrust",
		"conspicuously",
		"remint",
	}

	// Test NewStack
	fmt.Println("// string - NewStack()")
	newStack := NewStack()
	for key, value := range randomStrings {
		newStack.Push(randomStrings[key])
		fmt.Println(newStack.String() + " Pushed: " + value)
	}
	for !newStack.Empty() {
		element, err := newStack.Pop()
		if err != nil {
			t.Errorf("String Stack Error:" + err.Error())
		}
		fmt.Println(newStack.String() + " Popped: " + element.String())
	}

	// Test var stack.Stack
	fmt.Println("// string - stack.Stack")
	var vStack Stack
	for key, value := range randomStrings {
		vStack.Push(randomStrings[key])
		fmt.Println(vStack.String() + " Pushed: " + value)
	}
	for !vStack.Empty() {
		element, err := vStack.Pop()
		if err != nil {
			t.Errorf("String Stack Error:" + err.Error())
		}
		fmt.Println(vStack.String() + " Popped: " + element.String())
	}

	// Test BuildStack
	fmt.Println("// string - BuildStack")
	dataSlice := []string{"1", "2", "3", "4", "5", "6", "43", "423"}
	buildStack := BuildStack(dataSlice)
	for key, value := range randomStrings {
		buildStack.Push(randomStrings[key])
		fmt.Println(buildStack.String() + " Pushed: " + value)
	}
	for !buildStack.Empty() {
		element, err := buildStack.Pop()
		if err != nil {
			t.Errorf("Int Stack Error:" + err.Error())
		}
		fmt.Println(buildStack.String() + " Popped: " + element.String())
	}
}

type Node struct {
	lit string
}

func NewNode(lit string) *Node {
	return &Node{
		lit: lit,
	}
}

func (n *Node) String() string {
	return n.lit
}

func Test_Stack_Of_Stucts(t *testing.T) {

	fmt.Println("")
	fmt.Println("/*")
	fmt.Println("-------------------------")
	fmt.Println("---- TESTING STRUCTS ----")
	fmt.Println("-------------------------")
	fmt.Println("*/")
	fmt.Println("")
	fmt.Println("")

	// Generate sudo random strings
	fmt.Println("// Generating STRUCTS")
	randomStructs := [20]*Node{
		NewNode("disadvantageous"),
		NewNode("sarthe"),
		NewNode("nonsatiable"),
		NewNode("marguerite"),
		NewNode("rediscussion"),
		NewNode("fiducial"),
		NewNode("unmagnetic"),
		NewNode("respecter"),
		NewNode("amphoricity"),
		NewNode("floreated"),
		NewNode("heathenising"),
		NewNode("dextro"),
		NewNode("sartorius"),
		NewNode("hydrarch"),
		NewNode("vaporizable"),
		NewNode("baclava"),
		NewNode("reclothe"),
		NewNode("entrust"),
		NewNode("conspicuously"),
		NewNode("remint"),
	}

	// Test NewStack
	fmt.Println("// structs - NewStack()")
	newStack := NewStack()
	for key, value := range randomStructs {
		newStack.Push(randomStructs[key])
		fmt.Println(newStack.String() + " Pushed: " + value.String())
	}
	for !newStack.Empty() {
		element, err := newStack.Pop()
		if err != nil {
			t.Errorf("String Stack Error:" + err.Error())
		}
		fmt.Println(newStack.String() + " Popped: " + element.String())
	}

	// Test var stack.Stack
	fmt.Println("// struct - stack.Stack")
	var vStack Stack
	for key, value := range randomStructs {
		vStack.Push(randomStructs[key])
		fmt.Println(vStack.String() + " Pushed: " + value.String())
	}
	for !vStack.Empty() {
		element, err := vStack.Pop()
		if err != nil {
			t.Errorf("String Stack Error:" + err.Error())
		}
		fmt.Println(vStack.String() + " Popped: " + element.String())
	}

	// Test BuildStack
	fmt.Println("// struct - BuildStack")
	dataSlice := []*Node{
		NewNode("1"),
		NewNode("2"),
		NewNode("3"),
		NewNode("4"),
		NewNode("5"),
		NewNode("6"),
		NewNode("43"),
		NewNode("423"),
	}
	buildStack := BuildStack(dataSlice)
	for key, value := range randomStructs {
		buildStack.Push(randomStructs[key])
		fmt.Println(buildStack.String() + " Pushed: " + value.String())
	}
	for !buildStack.Empty() {
		element, err := buildStack.Pop()
		if err != nil {
			t.Errorf("Int Stack Error:" + err.Error())
		}
		fmt.Println(buildStack.String() + " Popped: " + element.String())
	}
}
