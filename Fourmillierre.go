package main

import (
	"log"
	"strconv"
	"strings"
)

func CreateRoom(SalleStart string, SalleEnd string, TabSalle [][]string) []Room {

	TabRoom := []Room{}

	TabRoom = append(TabRoom, CreateStart(SalleStart), CreateEnd(SalleEnd))

	for i := 0; i < len(TabSalle); i++ {

		CoX, _ := strconv.Atoi(TabSalle[i][1])
		CoY, _ := strconv.Atoi(TabSalle[i][2])

		TabRoom = append(TabRoom, CreateTheRoom(TabSalle[i][0], CoX, CoY))
	}

	return TabRoom
}
func CreateRelation(TabRelation [][]string, TabSalle []Room) []Tunnels {

	TabTunnel := []Tunnels{}

	for i := 0; i < len(TabRelation); i++ {

		RoomA := &Room{}
		RoomB := &Room{}

		for j := 0; j < len(TabSalle); j++ {
			if TabSalle[j].Name == TabRelation[i][0] {
				RoomA = &TabSalle[j]
			}
			if TabSalle[j].Name == TabRelation[i][1] {
				RoomB = &TabSalle[j]
			}
		}

		LaLiaison := CreateTheLiaison(RoomA, RoomB)

		RoomA.Tunnels = append(RoomA.Tunnels, LaLiaison)
		RoomB.Tunnels = append(RoomB.Tunnels, LaLiaison)

		TabTunnel = append(TabTunnel, LaLiaison)

	}

	for a := 0; a < len(TabTunnel); a++ {
		for b := a + 1; b < len(TabTunnel); b++ {
			if TabTunnel[a] == TabTunnel[b] {
				log.Fatal("05.D movais format, ne metez pas plusieur fois le même tunel")
			}
		}
	}

	return TabTunnel
}
func CreateFourmie(NbrFourmi int, start Room) []Ant {
	TabFourmi := []Ant{}

	for i := 0; i < NbrFourmi; i++ {
		TabFourmi = append(TabFourmi, CreateKevin(TabFourmi, start))
	}
	return TabFourmi
}
func CreateTheRoom(name string, X int, Y int) Room {
	return Room{name, X, Y, nil, false, false}
}
func CreateStart(salle string) Room {

	decoupage := strings.Split(salle, " ")

	name := decoupage[0]

	CoX, _ := strconv.Atoi(decoupage[1])
	CoY, _ := strconv.Atoi(decoupage[2])

	SalleStart := CreateTheRoom(name, CoX, CoY)

	SalleStart.Start = true

	return SalleStart
}
func CreateEnd(salle string) Room {

	decoupage := strings.Split(salle, " ")

	name := decoupage[0]

	CoX, _ := strconv.Atoi(decoupage[1])
	CoY, _ := strconv.Atoi(decoupage[2])

	SalleEnd := CreateTheRoom(name, CoX, CoY)

	SalleEnd.End = true

	return SalleEnd
}
func CreateTheLiaison(SA *Room, SB *Room) Tunnels {
	return Tunnels{SA, SB}
}
func CreateKevin(TabFourmi []Ant, Start Room) Ant {
	return Ant{len(TabFourmi) + 1, &Start, nil}
}
