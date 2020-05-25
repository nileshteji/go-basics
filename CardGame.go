package main
import ("fmt"
  "strings"
)
type deck []string

func main(){

cards := newDeck()
//fmt.Println(cards.hand(3))
fmt.Println(cards.toString())







}
func (d deck) toString() string {
  return strings.Join(d," ")
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
