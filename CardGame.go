package main
import "fmt"
type deck []string

func main(){

newDeck()






}

func (d deck) hand(variable int) deck{
handDeck := deck{};
for i:=0;i< int;i++
{
  handDeck=append(handDeck,d[i])
}
return handDeck

}

func newDeck() {


cardSuites := deck{"Spades","Hearts","Diamonds",}
cardValues := deck{"One","Two","Three","Four","Five","Six","Seven","Eight","Nine","Ten","Jacks","Queen","King"}
cards := deck{}



for _,cardS := range cardSuites{
 for _,cardV := range cardValues  {
    cards=append(cards,cardS+" of "+ cardV)
  }
}



for _,card:=range cards{
  fmt.Println(card)
}


}
