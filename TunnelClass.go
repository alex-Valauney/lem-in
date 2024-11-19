package main

import "fmt"

type Tunnels struct {
	SalleA *Room
	SalleB *Room
}

func (T Tunnels) GetInconue(SalleConnue *Room) *Room {

	if SalleConnue.Name != T.SalleA.Name {
		return T.SalleA
	}

	return T.SalleB
}

func (T Tunnels) Affichagetunnel() string {
	return fmt.Sprintf("%s-%s", T.SalleA.Name, T.SalleB.Name)
}
