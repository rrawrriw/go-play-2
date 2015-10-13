package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	tapas = []string{
		"Warme Chorizo-Tomaten Spieße",
		"Fleischbällchen im Kräuter-Kapern Mantel",
		"Curry-Hähnchenspieße mit Bananendip",
		"Tomaten-Thunfisch Päckchen",
		"Hähnchenleber in Sherrysauce",
		"Tapas Trio auf Weißbrot",
		"Sherry Garnelen",
		"Riesengarnelen mit Datteln",
		"Andalusische Salmorejo",
	}

	friends = []string{
		"Bob",
		"Alice",
		"Peter",
		"Bille",
		"Bobl",
	}

	wg sync.WaitGroup
)

func main() {
	for _, name := range friends {
		enjoy(name)
	}
	wg.Wait()
	fmt.Println("Muy buena comida!")

}

func enjoy(name string) {
	// Number of Tapas
	no := r(4) + 5

	for x := 0; x < no; x++ {
		wg.Add(1)

		// How long x eat
		ti := d(30, 300)
		// Pick random tapas
		tap := r(len(tapas))

		time.AfterFunc(ti, func() {
			defer wg.Done()
			fmt.Println(name, "enjoying", tapas[tap], "for", ti)
		})
	}
}

func d(s, e int) time.Duration {
	t := r(e-s) + s
	return time.Duration(t) * time.Millisecond

}

func r(m int) int {
	t := int64(time.Now().Nanosecond())
	s := rand.NewSource(t)
	r := rand.New(s)

	return r.Intn(m - 1)
}
