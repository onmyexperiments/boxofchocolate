package boxofchocolate

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type chocolate struct {
	name   string
	nature string
}

type chocolateBox struct {
	name       string
	chocolates []chocolate
}

// GetMeAChocolate is an HTTP Cloud Function.
func GetMeAChocolate(w http.ResponseWriter, r *http.Request) {

	chocolateBox := getChocolateBox()
	c := chocolateBox.getRandomChocolate()
	smile := "You have got a " + c.name + " " + c.nature + " chocolate.. Enjoy till last bite :)"
	fmt.Fprint(w, smile)
	return

}

func getChocolateBox() chocolateBox {
	cb := chocolateBox{
		name:       "LoveForChocolate",
		chocolates: getChocolates(),
	}
	cb.shuffle()
	return cb
}

func getChocolates() []chocolate {

	chocolates := []chocolate{}

	chocolateTypes := []string{"Dark", "White", "Milk", "Ruby"}
	chocolateTastes := []string{"Bittersweet", "Bitter", "Sweet"}

	for _, chocolatetype := range chocolateTypes {
		for _, chocolatetaste := range chocolateTastes {
			c := chocolate{chocolatetaste, chocolatetype}
			chocolates = append(chocolates, c)
		}
	}
	return chocolates
}

func (c chocolateBox) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range c.chocolates {
		newPosition := r.Intn(len(c.chocolates) - 1)

		c.chocolates[i], c.chocolates[newPosition] = c.chocolates[newPosition], c.chocolates[i]
	}
}

func (c chocolateBox) getRandomChocolate() chocolate {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return c.chocolates[r.Intn(len(c.chocolates)-1)]
}
