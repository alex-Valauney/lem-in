package main

import (
	"fmt"
	"os"
)

func VerrifArgs() bool { // vérifi qu'il y a 1 argument
	if len(os.Args) <= 1 {
		fmt.Println("Fichier non rentrée")
		return false
	}

	if len(os.Args) > 2 {
		fmt.Println("trop d'argument")
		return false
	}
	return true
}
func VerifNum(text string) bool {
	for i := 0; i < len(text); i++ {
		if (text[i] < '0' || text[i] > '9') && (text[i] != ' ' && text[i] != '-') {
			return false
		}
	}
	return true
}
func verifName(text string) bool { //vérifie si le nom ne contien pas L et #

	if text[0] == 'L' && text[0] == '#' {
		return false
	}

	return true
}
func verifName2(text []string) bool { //vérifi les doublon sur un même liaison

	if text[0] == text[1] {
		return false
	}

	return true
}
func verifName3(text [][]string) bool { //vérifi si il utilise que des nom donner précédament

	file := os.Args[1]

	ListName := []string{}

	Start := SalleStart(file)
	End := SalleEnd(file)

	NameStart := string(Start[0])
	NameEnd := string(End[0])

	ListName = append(ListName, NameStart, NameEnd)

	Autre := TrouverSalle(file)

	for a := 0; a < len(Autre); a++ {
		ListName = append(ListName, Autre[a][0])
	}

	for i := 0; i < len(text); i++ {
		for j := 0; j < len(text[i]); j++ {
			valide := false
			for k := 0; k < len(ListName); k++ {
				if text[i][j] == ListName[k] {
					valide = true
				}
			}
			if valide == false {
				return false
			}
		}
	}

	return true
}
