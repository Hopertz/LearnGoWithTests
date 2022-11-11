package main

import "testing"

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
