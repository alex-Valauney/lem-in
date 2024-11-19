package main

import "fmt"

type Room struct {
	Name    string
	Xcoord  int
	Ycoord  int
	Tunnels []Tunnels
	Start   bool
	End     bool
}

func (R *Room) AffichageRoom() string {
	strFinal := fmt.Sprintf("%s -> ", R.Name)

	for _, s := range R.Tunnels {
		salle := s.GetInconue(R)
		strFinal += fmt.Sprintf("<- %s ", salle.Name)
	}

	return strFinal
}
