package shapes

import (
	"testing"
)

func TestPerimeter(t *testing.T)  {
	rectangle:=Rectangle{10.0,10.0}
	got:=Perimeter(rectangle)
	want:=40.0

	if got!=want {
		t.Errorf("got %.2f want %.2f",got,want)
	}
}

func TestArea(t *testing.T){

	areaTests:=[]struct{
		name string
		shape Shape
		want float64
	}{
		{name : "rectangle" , shape : Rectangle{10.0,4.0} , want : 40.0},
		{name : "circle" , shape : Circle{10.0} , want : 314.1592653589793},
		{name : "triangle" , shape : Triangle{10.0,4.0} , want : 20},
	}

	for _, areaTest := range areaTests {
		t.Run(areaTest.name , func(t *testing.T) {
			got := areaTest.shape.Area()
			if got != areaTest.want {
				t.Errorf("got %.2f want %.2f",got,areaTest.want)
			}
		})
	}
}