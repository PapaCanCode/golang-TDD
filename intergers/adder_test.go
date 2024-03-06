package intergers

import (
	"fmt"
	"testing"
)

// Adds two integers and returns the sum
func Add(x, y int) int {
	return x + y
}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' got '%d' ", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	//Output is 6
}
