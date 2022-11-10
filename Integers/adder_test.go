package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(5, 6)
	expected := 11

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}

}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
