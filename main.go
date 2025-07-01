package main

import "fmt"

func InitializeGameEngine() *GameEngine {
	ge := &GameEngine{
		maxDepth: 3,
		state: &GameState{
			Players:    make(map[string]*Player),
			EventQueue: make([]Event, 0),
			Turn:       "player1",
		},
	}

	ge.state.Players["player1"] = &Player{
		ID:        "player1",
		Hand:      make([]*Card, 0),
		Stack:     make([]*Card, 0),
		Graveyard: make([]*Card, 0),
	}

	ge.state.Players["player2"] = &Player{
		ID:        "player2",
		Hand:      make([]*Card, 0),
		Stack:     make([]*Card, 0),
		Graveyard: make([]*Card, 0),
	}
	return ge
}

func InitializeGameScenario() *GameEngine {
	card1 := &Card{
		ID:    "1",
		Owner: "player1",
		Abilities: []Ability{
			{
				Name:   "Stack Destroyer",
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
				Name:    "Restoration",
				Type:    "triggered",
				Trigger: "on_destroy",
				Effect:  "salvage_self",
				Target:  "player_hand",
			},
		},
	}

	ge := InitializeGameEngine()
	ge.state.Players["player1"].Hand = append(ge.state.Players["player1"].Hand, card1)
	ge.state.Players["player2"].Stack = append(ge.state.Players["player2"].Stack, card2)
	return ge
}

func main() {
	ge := InitializeGameScenario()

	ge.PrintGameState()
	err := ge.PlayCard("player1", "1")
	if err != nil {
		fmt.Println(err)
	}

	ge.PrintGameState()
}
