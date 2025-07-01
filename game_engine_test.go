package main

import (
	"testing"
)

func InitializeTestGameScenario() (*GameEngine, *Card) {
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
	card3 := *card2
	card3.ID = "3"

	card4 := *card2
	card4.ID = "4"

	ge := InitializeGameEngine()

	ge.state.Players["player1"].Hand = append(ge.state.Players["player1"].Hand, card1)

	ge.state.Players["player2"].Stack = append(ge.state.Players["player2"].Stack, card2)
	ge.state.Players["player2"].Stack = append(ge.state.Players["player2"].Stack, &card3)
	ge.state.Players["player2"].Stack = append(ge.state.Players["player2"].Stack, &card4)
	//ge.state.Players["player2"].Stack = append(ge.state.Players["player2"].Stack, card2)
	//ge.state.Players["player2"].Stack = append(ge.state.Players["player2"].Stack, card2)
	return ge, card1
}

func TestGameEngine_ExecuteAbility(t *testing.T) {
	ge, card1 := InitializeTestGameScenario()

	err := ge.ExecuteAbility(card1, card1.Abilities[0], 0)
	if err != nil {
		t.Fatal(err)
	}
	err = ge.ProcessEventQueue()
	if err != nil {
		t.Fatal(err)
	}
	ge.PrintGameState()
}
