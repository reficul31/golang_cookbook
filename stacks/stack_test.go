package stacks

import "testing"

var err error

// Test the stack capacity allocation function NewStack
func TestStackCapacityAllocation(testCase *testing.T) {
	testCase.Log("To test that the stack size is allocated correctly")

	if s, err := NewStack(20); err == nil {
		if len(s.arr) != 20 {
			testCase.Errorf("Stack Error: Stack size not equal to capacity")
		}
	}
	
	testCase.Errorf("Stack Error: %s", err)
}

// Test the Push function of the Stack
func TestStackPush(testCase *testing.T) {
	testCase.Log("To test that the stack push function pushes numbers on top")
	
	s, _ := NewStack(2)
	if err = s.Push(1) ; err != nil {
		testCase.Errorf("Stack Error: %s", err)
	} else {
		if s.arr[s.top] != 1 {
			testCase.Errorf("Stack Error: Element not pushed")
		}
	}

	err = s.Push(2)
	if err = s.Push(3) ; err == nil {
		testCase.Errorf("Stack Error: Overflow not reported")
	}
}

// Test the Pop function of the Stack
func TestStackPop(testCase *testing.T) {
	testCase.Log("To test that the stack pop function pops the number on top")
	var element int

	s, _ := NewStack(3)
	_ = s.Push(3)
	_ = s.Push(2)
	_ = s.Push(1)

	if element, err = s.Pop() ; err != nil {
		testCase.Errorf("Stack Error: %s", err)
	}
	if element != 1 {
		testCase.Errorf("Stack Error: Element not popped")
	}
	
	_, err = s.Pop()
	_, err = s.Pop()

	if _, err = s.Pop() ; err == nil {
		testCase.Errorf("Stack Error: Stack Underflow not reported")
	}
}