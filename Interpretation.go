package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func NombreFourmi(filename string) int {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	v := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		if LigneNbrDeFourmi, OnPasse := strconv.Atoi(text); OnPasse == nil {
			if v == 0 {
				v = LigneNbrDeFourmi
			} else {
				log.Fatal("01.A movais format, metez qu'un nombre de fourmi")
			}
		}
	}
	if v == 0 {
		log.Fatal("01.B movais format, mettez un nombre de fourmi")
	}
	file.Close()
	return v
}
func SalleStart(filename string) string {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	Start := ""

	for scanner.Scan() {
		text := scanner.Text()
		valide := false

		if text == "##start" {
			valide = true
		}

		if valide {
			if Start == "" {
				scanner.Scan()
				Start = string(scanner.Text())

				if Start == "" {
					log.Fatal("02 mauvais format, rentrée des informations pour la salle start")
				}

				sep := " " // Séparateur
				decoupage := strings.Split(Start, sep)

				if len(decoupage) != 3 {
					log.Fatal("02.A mauvais format, metez que trois nombres au format 0 1 0 pour la salle start")
				}

				if !VerifNum(Start) {
					log.Fatal("02.B mauvais format, metez que trois nombres au format 0 1 0 pour la salle start")
				}
			} else {
				log.Fatal("02.C mauvais format, metez qu'une entrée")
			}
		}
	}

	file.Close()
	return Start
}
func SalleEnd(filename string) string {

	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	End := ""

	for scanner.Scan() {
		text := scanner.Text()
		valide := false

		if text == "##end" {
			valide = true
		}

		if valide {
			if End == "" {
				scanner.Scan()
				End = string(scanner.Text())

				if End == "" {
					log.Fatal("03 mauvais format, rentrée des informations pour la salle end")
				}

				sep := " " // Séparateur
				decoupage := strings.Split(End, sep)

				if len(decoupage) != 3 {
					log.Fatal("03.A mauvais format, metez que trois nombres au format 0 1 0 pour la salle end")
				}

				if !VerifNum(End) {
					log.Fatal("03.B mauvais format, metez que trois nombres au format 0 1 0 pour la salle end")
				}
			} else {
				log.Fatal("03.C mauvais format, metez qu'une sortie")
			}
		}
	}

	file.Close()
	return End
}
func TrouverSalle(filename string) [][]string {

	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	Salle := [][]string{}

	for scanner.Scan() {
		text := scanner.Text()
		sep := " "

		if text != "##start" && text != "##end" {
			if len(text) > 1 {
				decoupage := strings.Split(text, sep)

				if len(decoupage) != 3 {
					continue
				}
				if verifName(decoupage[0]) {
					if VerifNum(decoupage[1]) && VerifNum(decoupage[2]) {
						Salle = append(Salle, decoupage)
					}
				}
			}
		} else {
			scanner.Scan()
			continue
		}
	}

	file.Close()
	return Salle
}
func TrouverLiaison(filename string) [][]string {

	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	Liaison := [][]string{}

	for scanner.Scan() {
		text := scanner.Text()
		sep := "-"

		decoupage := strings.Split(text, sep)

		if len(decoupage) == 2 {
			if !verifName2(decoupage) {
				log.Fatal("05.A movait format ne métait pas deux foix le même nom dans la lieson")
			}
			if verifName(decoupage[0]) && verifName(decoupage[1]) {
				Liaison = append(Liaison, decoupage)
			}
		}
	}
	if len(Liaison) == 0 {
		log.Fatal("05 Metez des conection de salle sous le format nom-nom, avec des nom diférent")
	}
	if !verifName3(Liaison) {
		log.Fatal("05.B movait format, metez un nom de salle définit plutôt")
	}
	return Liaison
}
