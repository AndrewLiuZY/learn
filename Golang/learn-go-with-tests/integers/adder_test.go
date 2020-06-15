package integers

import (
	"testing"
	"fmt"
)

func TestAdder(t *testing.T) {
	sum := Add(1 , 2)	
	expected := 3

	if sum!=expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd(){
	sum:=Add(10,5)
	fmt.Println(sum)
	//Output: 15
}