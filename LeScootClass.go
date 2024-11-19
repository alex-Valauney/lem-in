package main

import (
	"bufio"
	"fmt"
	"os"
)

type Scoot struct {
	Position      *Room
	Chemin        []*Room
	ListeDeChemin [][]*Room
}

func (fourmie *Scoot) ListeDeSalle() []*Room {

	stock := []*Room{}
	for _, t := range fourmie.Position.Tunnels {
		stock = append(stock, t.GetInconue(fourmie.Position))
	}

	return stock
}
func (fourmie *Scoot) Eclerage(nouvSalle *Room) {
	fourmie.Position = nouvSalle
	fourmie.Chemin = append(fourmie.Chemin, nouvSalle)

	for _, s := range fourmie.Chemin {
		fmt.Print(s.Name + " ")
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if fourmie.Position.End {
		fmt.Println("End Reached")
		fourmie.ListeDeChemin = append(fourmie.ListeDeChemin, fourmie.Chemin)
		return
	} else {
		for _, s := range fourmie.ListeDeSalle() {
			valid := true
			for _, c := range fourmie.Chemin {
				if s == c {
					valid = false
				}
			}
			if valid {
				temp := fourmie.Chemin
				fourmie.Eclerage(s)
				fourmie.Chemin = temp
				fourmie.Position = nouvSalle
			}
		}
	}
}
