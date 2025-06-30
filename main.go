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

func (g *GameEngine) PlayCard(playerId string, cardId string) {
	//get player
	//get card from hand
	//remove card from hand
	//execute all abilities
	//process queue
	//move card to graveyard
	//change turn
}

func (g *GameEngine) ExecuteAbility(source *Card, ability Ability) {
	//check Ability Effect and handle appropriately
}
func (g *GameEngine) ProcessEventQueue() {
	//while event queue is not empty
	//process event
}

func main() {
	card1 := &Card{
		ID:    "1",
		Owner: "player1",
		Abilities: []Ability{
			{
				Type:   "instant",
				Effect: "destroy_stack",
				Target: "enemy_stack",
			},
		},
	}

	card2 := &Card{
		ID:    "2",
		Owner: "player1",
		Abilities: []Ability{
			{
				Type:   "triggered",
				Effect: "salvage_self",
				Target: "player_hand",
			},
		},
	}

	ge := &GameEngine{
		state: &GameState{
			Players:    make(map[string]*Player),
			EventQueue: make([]Event, 0),
		},
	}

	ge.state.Players["player1"] = &Player{
		ID:        "player1",
		Hand:      make([]*Card, 0),
		Stack:     make([]*Card, 0),
		Graveyard: make([]*Card, 0),
	}

	ge.state.Players["player1"].Stack = append(ge.state.Players["player1"].Stack, card1)

	ge.state.Players["player2"] = &Player{
		ID:        "player2",
		Hand:      make([]*Card, 0),
		Stack:     make([]*Card, 0),
		Graveyard: make([]*Card, 0),
	}

	ge.state.Players["player2"].Stack = append(ge.state.Players["player2"].Stack, card2)

	ge.PlayCard("player1", "1")
}
