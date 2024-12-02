package main

import (
	"fmt"
	"github.com/stantowne/DeckMaker/deckMaker"
	"github.com/stantowne/badFunctions/badAdder"
	"github.com/stantowne/badFunctions/badOther"
)

func main() {
	badSum := badAdder.BadAdder(11, 45)
	veryBadSum := badAdder.VeryBadAdder(11, 45)
	evenWorse := badOther.BadMultiplier(4, 3)
	fmt.Println(badSum, veryBadSum, evenWorse)
	newDecks := deckMaker.DeckMaker(77)
	fmt.Println(len(newDecks))

}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
