package main

import "fmt"
type shape interface{
  calculateArea() int
}

type triangle struct{
height float64
base float64
}

type square struct{
  sideLength float64
}


func main()  {

  t := triangle{10,5}
  s := square{10}
  print(t)
  print(s)

}

func print(b shape)  {
  fmt.Println(b.calculateArea())
}
func (t triangle) calculateArea() int {
 return int (0.5 * t.base * t.height)
}


func (s square) calculateArea() int{
return int(s.sideLength * s.sideLength)
}
