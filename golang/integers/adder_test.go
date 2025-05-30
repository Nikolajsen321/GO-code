package integers

import (
	"fmt"
	"testing"
)

// TestAdder tester Add-funktionen.
func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// ExampleAdd demonstrerer Add-funktionen.
// Dette vil blive vist i dokumentationen.
func ExampleAdd() {
	sum := Add(10, 15)
	fmt.Println(sum)
	// Output: 150
}
