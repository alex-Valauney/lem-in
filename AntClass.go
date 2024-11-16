package main

import "fmt"

type Ant struct {
	Name          int
	SalleActuelle *Room
}

func (fourmie *Ant) GetSalle() []Room {

	stock := []Room{}

	for _, t := range fourmie.SalleActuelle.Tunnels {
		stock = append(stock, t.GetInconue(*fourmie.SalleActuelle))
	}

	return stock
}

func (fourmie *Ant) Deplacer(nouvSalle Room) bool {

	for _, s := range fourmie.GetSalle() {
		if s.Name == nouvSalle.Name {
			fourmie.SalleActuelle = &nouvSalle

			fmt.Printf("L%d-%s", fourmie.Name, fourmie.SalleActuelle.Name)
			return true
		}
	}

	return false
}
