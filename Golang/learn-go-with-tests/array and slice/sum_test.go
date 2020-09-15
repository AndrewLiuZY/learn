package main

import (
	"testing"
	"reflect"
)

func TestSum(t *testing.T)  {

	numbers:=[]int{1,2,3,4,5}

	got:=Sum(numbers)
	want:=15

	if got!=want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T)  {
	
	got:=SumAll([]int{1,2,3},[]int{2,3,4})
	want:=[]int{6,9}

	if !reflect.DeepEqual(got,want) {
		t.Errorf("got %v want %v",got ,want)
	}

}

func TestSumAllTails(t *testing.T){

	assert:=func(t *testing.T,got,want []int){
		if !reflect.DeepEqual(got,want) {
			t.Errorf("got %v want %v",got,want)
		}
	}

	t.Run("make the sums of some slice",func(t *testing.T){
		got:=SumAllTails([]int{1,2,3},[]int{1,1,1})
		want:=[]int{5,2}

		assert(t,got,want)
	})

	t.Run("safely sum empty slices",func(t *testing.T){
		got:=SumAllTails([]int{1,2,3},[]int{})
		want:=[]int{5,0}

		assert(t,got,want)
	})
}