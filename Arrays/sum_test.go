package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		arrayOfNumbers := []int{1, 2, 3, 4, 5}
		sum := Sum(arrayOfNumbers)
		want := 15

		if sum != want {
			t.Errorf("got %d want %d given, %v ", sum, want, arrayOfNumbers)
		}

	})

}
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
