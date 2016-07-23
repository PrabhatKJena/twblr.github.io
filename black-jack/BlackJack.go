package blackjack

import (
	"fmt"
)

type Card struct{
	cardType string
	faceValue int
	alterValue int
	symbol string
	isHidden bool
}

type Deck struct{
	D,C,S,H [13]Card
}

func (deck *Deck) InitDeck(){
	deck.InitCard("D", &deck.D) // diamond
	deck.InitCard("C", &deck.C) //
	deck.InitCard("S", &deck.S) //
	deck.InitCard("H", &deck.H) // heart
}

func (deck *Deck) InitCard (cType string,suit *[13]Card){
	suit[0] = Card{cType,1,11,"A", false}
	i :=2
	for i<=10{ // 1..9
		suit[i-1] = Card{cType,i,0,fmt.Sprintf("%v",i), false}
		i++;
	}
	suit[10] = Card{cType,10,0,"J", false}
	suit[11] = Card{cType,11,0,"Q", false}
	suit[12] = Card{cType,12,0,"K", false}
}

type Player interface{
	value() int
	hit(card Card)
	stand()
	opt(card Card) int
}

func (player *User) opt(card Card) int {
	if player.val + card.alterValue < 21 {
		return card.alterValue
	}
	return card.faceValue
}
func (player *Dealer) opt(card Card) int {
	if player.val + card.alterValue < 21 {
		return card.alterValue
	}
	return card.faceValue	
}
func (player *User) hit(card Card){
	if card.symbol == "A"{
		player.val = player.val + player.opt(card)
	}else{
		player.val = player.val + card.faceValue
	}
	player.cardOnTbl = append(player.cardOnTbl, card)
}

func (player *User) value() int{
	return player.val
}
func (player *User) stand(){

}
func (player *Dealer) stand(){
	
}
func (player *Dealer) hit(card Card){
	if card.symbol == "A"{
		player.val = player.val + player.opt(card)
	}else{
		player.val = player.val + card.faceValue
	}
	player.cardOnTbl = append(player.cardOnTbl, card)
}

func (player *Dealer) value() int{
	return player.val
}

type User struct{
	cardOnTbl []Card
	val int
}

type Dealer struct{
	deck *Deck
	cardOnTbl []Card
	val int
}

type Game struct{
	user Player
	dealer Player
}
func (game *Game) result() (string, int){
	uVal := game.user.value()
	dVal := game.dealer.value()
	if(uVal > 21){
		return "D", 1
	}
	if (dVal > 21){
		return "U",1
	}
	if uVal > dVal {
		return "U", 1
	}else if uVal < dVal{
		return "D", 1
	}
	return	"", -1
}
func (game *Game) userTurn() (string, int){
	return game.result()
}
func (game *Game) dealerTurn() (string, int){
	return game.result()
}
func (game *Game) initGame(){
	/*deck:= &Deck{}
	deck.InitDeck()
	user:= *User{}
	dealer := *Dealer{deck}
	cTypeN:=rand.Intn(4)*/

}
func main(){
	fmt.Println("Welcome to Black Jack Game\n==========================")

}
