package main 



func Perimeter(rectangle Rectangle ) float64{

     return  2 * (rectangle.Width + rectangle.Height)

}

func Area (rectangle Rectangle) float64{
	return rectangle.Height * rectangle.Width
}







//declaring a struct

type Rectangle struct{
	Width float64
	Height float64
}