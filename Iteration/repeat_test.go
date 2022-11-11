package Iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("got %q wanted %q ", repeated, expected)
	}

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)

	}
}

func ExampleRepeat() {
	repeat := Repeat("la", 3)
	fmt.Println(repeat)
	// Output: lalala
}
