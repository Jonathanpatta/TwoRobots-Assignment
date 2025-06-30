package main

import (
	"errors"
	"fmt"
)

func (g *GameEngine) DestroyStack(source *Card, depth int) {

	//player:= "player1"
	enemy := "player2"
	if source.Owner == "player2" {
		//player = "player2"
		enemy = "player1"
	}
	fmt.Println("Destroying Stack of ", enemy)
	for _, card := range g.state.Players["player2"].Stack {
		event := Event{
			ID:     "destroy_card",
			Source: source,
			Target: card,
			Depth:  depth,
		}
		g.state.EventQueue = append(g.state.EventQueue, event)
	}
}

func (g *GameEngine) SalvageSelf(source *Card, depth int) error {
	fmt.Println("Salvaging Self of ", source)
	player := g.state.Players[source.Owner]
	player.Hand = append(player.Hand, source)
	for i, graveyardCard := range player.Graveyard {
		if graveyardCard.ID == source.ID {
			player.Graveyard = append(player.Graveyard[:i], player.Graveyard[i+1:]...)
			break
		}
	}
	return nil
}

func (g *GameEngine) DestroyCard(event Event) error {
	affectedCard := event.Target

	fmt.Println("Destroying Card:", affectedCard)
	player := g.state.Players[event.Target.Owner]
	found := false
	for i, stackCard := range player.Stack {
		if stackCard.ID == affectedCard.ID {
			player.Stack = append(player.Stack[:i], player.Stack[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		return errors.New("card not found in player stack")
	}
	for _, ability := range affectedCard.Abilities {
		if ability.Trigger == "on_destroy" {
			fmt.Println("On Destroy Ability Triggered")
			err := g.ExecuteAbility(affectedCard, ability, event.Depth+1)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
