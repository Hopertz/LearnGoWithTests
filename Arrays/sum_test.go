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
	t.Run("collection of zero size", func(t *testing.T) {
		arrayOfNumbers := []int{}
		sum := Sum(arrayOfNumbers)
		want := 0

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

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("Make the sum of some of the slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("Safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

}
