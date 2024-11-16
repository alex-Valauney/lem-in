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
	tadDeRelation := CreateRelation(TrouverLiaison(file), tabDeRoom)
	tabDeFourmie := CreateFourmie(NombreFourmi(file), tabDeRoom[0])

	fmt.Println()
	for _, r := range tabDeRoom {
		fmt.Println(r.AffichageRoom())
	}
	fmt.Println()
	for _, t := range tadDeRelation {
		fmt.Println(t.Affichagetunnel())
	}
	fmt.Println()
	fmt.Println(tabDeFourmie)

	fmt.Println(time.Since(timer))
}
