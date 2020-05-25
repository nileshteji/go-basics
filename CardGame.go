package main
import "fmt"
type deck []string

func main(){

cards := newDeck()
fmt.Println(cards.hand(3))







}


func (d deck) hand(variable int) (deck,deck) {
return d[:variable],d[variable:]
}

func newDeck() deck {


cardSuites := deck{"Spades","Hearts","Diamonds",}
cardValues := deck{"One","Two","Three","Four","Five","Six","Seven","Eight","Nine","Ten","Jacks","Queen","King"}
cards := deck{}



for _,cardS := range cardSuites{
 for _,cardV := range cardValues  {
    cards=append(cards,cardS+" of "+ cardV)
  }
}



return cards

}
