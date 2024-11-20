package main

func TriDeChemin(ListeDeChemin [][]*Room) [][][]*Room {

	ListeGroupeChemin := [][][]*Room{}

	for i := 0; i < len(ListeDeChemin); i++ {

		GroupeChemin := [][]*Room{ListeDeChemin[i]}

		for j := i + 1; j < len(ListeDeChemin); j++ {
			if CheminIndependant(ListeDeChemin[i], ListeDeChemin[j]) {
				GroupeChemin = append(GroupeChemin, ListeDeChemin[j])
			}
		}
		ListeGroupeChemin = append(ListeGroupeChemin, GroupeChemin)
	}
	return ListeGroupeChemin
}

func CheminIndependant(CheminA []*Room, CheminB []*Room) bool {

	for i := 1; i < len(CheminA)-1; i++ {
		for j := 1; j < len(CheminB)-1; j++ {
			if CheminA[i].Name == CheminB[j].Name {
				return false
			}
		}
	}
	return true
}
