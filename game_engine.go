package main

import (
	"fmt"
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
}

type Event struct {
	ID     string
	Type   string
	Player string
	Depth  int
	Source *Card
	Target *Card
}

type GameEngine struct {
	maxDepth int
	state    *GameState
}

func (g *GameEngine) PlayCard(playerId string, cardId string) {
	//get player
	player := g.state.Players[playerId]
	//get card from hand
	var card *Card
	cardIndex := -1
	for i, c := range player.Hand {
		if c.ID == cardId {
			card = c
			cardIndex = i
			break
		}
	}
	//remove card from hand
	player.Hand = append(player.Hand[:cardIndex], player.Hand[cardIndex+1:]...)
	//execute all abilities
	for _, ability := range card.Abilities {
		if ability.Type == "instant" {
			g.ExecuteAbility(card, ability, 0)
		}
	}
	//process queue
	g.ProcessEventQueue()
	//move card to graveyard
	player.Graveyard = append(player.Graveyard, card)
	//change turn
	if g.state.Turn == "player1" {
		g.state.Turn = "player2"
	} else {
		g.state.Turn = "player1"
	}
}

func (g *GameEngine) ExecuteAbility(source *Card, ability Ability, depth int) {
	if depth > g.maxDepth {
		return
	}
	//check Ability Effect and handle appropriately
	switch ability.Effect {
	case "destroy_stack":
		g.DestroyStack(source, depth)
		break
	case "salvage_self":
		g.SalvageSelf(source, depth)
		break
	default:
		return
	}
}

func (g *GameEngine) ProcessEventQueue() {
	//while event queue is not empty
	//process event

	for len(g.state.EventQueue) > 0 {
		event := g.state.EventQueue[0]
		g.state.EventQueue = g.state.EventQueue[1:]
		g.ProcessEvent(event)
	}
}

func (g *GameEngine) ProcessEvent(event Event) {
	fmt.Println("ProcessEvent:", event.ID)
	if event.Depth > g.maxDepth {
		return
	}

	switch event.ID {
	case "destroy_card":
		g.DestroyCard(event)
	}
}
