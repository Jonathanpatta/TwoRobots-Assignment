package main

import "fmt"

func (g *GameEngine) DestroyStack(source *Card, depth int) {
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

func (g *GameEngine) SalvageSelf(source *Card, depth int) {
	fmt.Println("salvage_self")
	player := g.state.Players[source.Owner]
	player.Hand = append(player.Hand, source)
	for i, graveyardCard := range player.Graveyard {
		if graveyardCard.ID == source.ID {
			player.Graveyard = append(player.Graveyard[:i], player.Graveyard[i+1:]...)
			break
		}
	}
}

func (g *GameEngine) DestroyCard(event Event) {
	affectedCard := event.Target

	fmt.Println("AffectedCard:", affectedCard)
	player := g.state.Players[event.Target.Owner]
	for i, stackCard := range player.Stack {
		if stackCard.ID == affectedCard.ID {
			player.Stack = append(player.Stack[:i], player.Stack[i+1:]...)
			break
		}
	}
	for _, ability := range affectedCard.Abilities {
		if ability.Trigger == "on_destroy" {
			g.ExecuteAbility(affectedCard, ability, event.Depth+1)
		}
	}
}
