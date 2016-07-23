package blackjack

import (
	"testing"
)

// Test Case - 1
func TestCardValueInDeck(t *testing.T){
	deck:= &Deck{}
	deck.InitDeck()
	testSuit("D", deck.D, t)
	testSuit("C", deck.C, t)
	testSuit("H", deck.H, t)
	testSuit("S", deck.S, t)
}

// test case -2
func TestDealerBursted(t *testing.T){
	user := &User{}
	dealer := &Dealer{}
	card1 := Card{"D",10,0,"J",false}
	card2 := Card{"C",9,0,"9",false}
	user.hit(card1)
	user.hit(card2)
	user.stand()

	card3 := Card{"H",4,0,"4",false}
	card4 := Card{"D",7,0,"7",false}
	card5 := Card{"C",11,0,"Q",false}
	dealer.hit(card3)
	dealer.hit(card4)
	dealer.hit(card5)

	game:=Game{user, dealer}
	
	p,_:=game.result();
	if(p!="U"){
		t.Errorf("Expected User to be winner", p)
	}
}

// test case -2
func TestUserBursted(t *testing.T){
	user := &User{}
	dealer := &Dealer{}
	card1 := Card{"D",10,0,"J",false}
	card2 := Card{"C",2,0,"2",false}
	card3 := Card{"H",9,0,"9",false}
	user.hit(card1)
	user.hit(card2)
	user.hit(card3)
	
	game:=Game{user, dealer}
	
	p,_:=game.result();
	if(p!="U"){
		t.Errorf("Expected User to be winner", p)
	}
}

// test case -3
func TestDealerStand(t *testing.T){
	user := &User{}
	dealer := &Dealer{}
	card1 := Card{"D",10,0,"J",false}
	card2 := Card{"C",8,0,"9",false}
	user.hit(card1)
	user.hit(card2)
	user.stand()

	card3 := Card{"D",7,0,"7",false}
	card4 := Card{"D",3,0,"3",false}
	card5 := Card{"C",9,0,"9",false}
	dealer.hit(card3)
	dealer.hit(card4)
	dealer.hit(card5)

	game:=Game{user, dealer}
	
	p,_:=game.result();
	if(p!="D"){
		t.Errorf("Expected Dealer to be winner", p)
	}
}

// test case -4
func TestUserHavingAce(t *testing.T){
	user := &User{}
	dealer := &Dealer{}
	card1 := Card{"D",2,0,"2",false}
	card2 := Card{"C",8,0,"8",false}
	card3 := Card{"H",1,11,"A",false}
	user.hit(card1)
	user.hit(card2)
	user.hit(card3)
		
	game:=Game{user, dealer}
	p,_:=game.result();
	if(p!="U"){
		t.Errorf("Expected User to be winner", p)
	}
}

func testSuit(cType string, suit [13]Card,t *testing.T){
	for _,card := range suit{
		if card.cardType != cType {
			t.Errorf("Expected %v but found %v", cType, card.cardType)
		}
	}
}

