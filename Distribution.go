package main

func Distrib(GroupeChemin [][]*Room, Fourmies []Ant) {

	ComteurFourmi := 0
	comteurChemin := 0

	for ComteurFourmi < len(Fourmies) {
		for i := 0; i < len(GroupeChemin); i++ {
			if len(GroupeChemin[i]) < comteurChemin && ComteurFourmi < len(Fourmies) {
				Fourmies[ComteurFourmi].Chemin = GroupeChemin[i]
				ComteurFourmi++
			}
		}
		comteurChemin++
	}
}
