package main

import (
	"sync"
)

type Player struct {
	ID        string
	Hand      []*Card
	Stack     []*Card
	Graveyard []*Card
}
type Card struct {
	ID        string
	Type      string
	Owner     string
	Abilities []Ability
}

type Ability struct {
	Type    string
	Trigger string
	Effect  string
	Target  string
}

type GameState struct {
	Players    map[string]*Player
	Turn       string
	EventQueue []Event
	mutex      sync.RWMutex
}

type Event struct {
	ID     string
	Type   string
	Player string
	Depth  int
}

type GameEngine struct {
	state *GameState
}

func main() {

}
