package main

import "fmt"

type Ant struct {
	Name          int
	SalleActuelle *Room
	Chemin        []*Room
}

func (fourmie *Ant) GetSalle() []*Room {

	stock := []*Room{}

	for _, t := range fourmie.SalleActuelle.Tunnels {
		stock = append(stock, t.GetInconue(fourmie.SalleActuelle))
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
func (fourmie *Ant) OnAvence() {

	for i := 0; i < len(fourmie.Chemin); i++ {
		if fourmie.Chemin[i].Name == fourmie.SalleActuelle.Name && !fourmie.SalleActuelle.End {
			fourmie.Deplacer(*fourmie.Chemin[i+1])
		}
	}
}
