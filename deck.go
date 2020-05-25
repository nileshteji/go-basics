 package main


import "fmt"

type deck [] string



func main()  {
  cards  := deck{"Nilesh","teji"}
// this works as standard for loop
  for index := 0;  index< len(cards); index++ {
  fmt.Println(cards[index])
  }
//this works as a for  each loop in golang
  for _,index := range cards{
    fmt.Println(index)
  }
}
