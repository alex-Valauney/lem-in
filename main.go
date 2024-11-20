package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.Now()
	if !VerrifArgs() {
		return
	}
	file := Lecture()
	tabDeRoom := CreateRoom(SalleStart(file), SalleEnd(file), TrouverSalle(file))
	CreateRelation(TrouverLiaison(file), tabDeRoom)
	CreateFourmie(NombreFourmi(file), tabDeRoom[0])

	kevin := Scoot{nil, nil, nil}
	kevin.Eclerage(tabDeRoom[0])
	for _, i := range kevin.ListeDeChemin {
		for _, j := range i {
			fmt.Print(j.Name, " ")
		}
		fmt.Println()
	}

	fmt.Println(TriDeChemin(kevin.ListeDeChemin))

	fmt.Println()
	fmt.Println(time.Since(timer))
	fmt.Println()
}
