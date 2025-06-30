package main

import (
	"fmt"
)

func InitializeGameScenario() (*Card, *GameEngine) {
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
		Owner: "player2",
		Abilities: []Ability{
			{
				Type:    "triggered",
				Trigger: "on_destroy",
				Effect:  "salvage_self",
				Target:  "player_hand",
			},
		},
	}

	ge := &GameEngine{
		maxDepth: 3,
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

	ge.state.Players["player1"].Hand = append(ge.state.Players["player1"].Hand, card1)

	ge.state.Players["player2"] = &Player{
		ID:        "player2",
		Hand:      make([]*Card, 0),
		Stack:     make([]*Card, 0),
		Graveyard: make([]*Card, 0),
	}
	return card2, ge
}

func main() {
	card2, ge := InitializeGameScenario()

	ge.state.Players["player2"].Stack = append(ge.state.Players["player2"].Stack, card2)
	fmt.Println(ge.state.Players["player1"].Hand, ge.state.Players["player2"].Hand)
	fmt.Println(ge.state.Players["player1"].Stack, ge.state.Players["player2"].Stack)
	fmt.Println(ge.state.Players["player1"].Graveyard, ge.state.Players["player2"].Graveyard)
	ge.PlayCard("player1", "1")
	fmt.Println(ge.state.Players["player1"].Hand, ge.state.Players["player2"].Hand)
	fmt.Println(ge.state.Players["player1"].Stack, ge.state.Players["player2"].Stack)
	fmt.Println(ge.state.Players["player1"].Graveyard, ge.state.Players["player2"].Graveyard)
}
