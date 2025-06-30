package main

import (
	"fmt"
	"strings"
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

func (g *GameEngine) PrintGameState() {
	state := g.state

	fmt.Println("\n=== GAME STATE ===")
	for playerID, player := range state.Players {
		fmt.Printf("\n%s:\n", strings.ToUpper(playerID))
		fmt.Printf("  Hand (%d): ", len(player.Hand))
		for _, card := range player.Hand {
			fmt.Printf("%s ", card.ID)
		}
		fmt.Printf("\n  Stack (%d): ", len(player.Stack))
		for _, card := range player.Stack {
			fmt.Printf("%s ", card.ID)
		}
		fmt.Printf("\n  Graveyard (%d): ", len(player.Graveyard))
		for _, card := range player.Graveyard {
			fmt.Printf("%s ", card.ID)
		}
		fmt.Println()
	}
	fmt.Println("==================")
}

func (g *GameEngine) PlayCard(playerId string, cardId string) error {
	//get player
	player, ok := g.state.Players[playerId]
	if !ok {
		return fmt.Errorf("player %s not found", playerId)
	}

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
	if card == nil {
		return fmt.Errorf("card %s not found in hand", cardId)
	}
	//remove card from hand
	player.Hand = append(player.Hand[:cardIndex], player.Hand[cardIndex+1:]...)
	//execute all abilities
	for _, ability := range card.Abilities {
		if ability.Type == "instant" {
			err := g.ExecuteAbility(card, ability, 0)
			if err != nil {
				return err
			}
		}
	}
	//process queue
	err := g.ProcessEventQueue()
	if err != nil {
		return err
	}
	//move card to graveyard
	g.state.Players[playerId].Graveyard = append(g.state.Players[playerId].Graveyard, card)
	//change turn
	if g.state.Turn == "player1" {
		g.state.Turn = "player2"
	} else {
		g.state.Turn = "player1"
	}
	return nil
}

func (g *GameEngine) ExecuteAbility(source *Card, ability Ability, depth int) error {
	if depth > g.maxDepth {
		return fmt.Errorf("max depth of %d exceeded", g.maxDepth)
	}
	//check Ability Effect and handle appropriately
	switch ability.Effect {
	case "destroy_stack":
		g.DestroyStack(source, depth)
		break
	case "salvage_self":
		err := g.SalvageSelf(source, depth)
		if err != nil {
			return err
		}
		break
	default:
		return fmt.Errorf("unknown ability effect: %s", ability.Effect)
	}
	return nil
}

func (g *GameEngine) ProcessEventQueue() error {
	//while event queue is not empty
	//process event

	for len(g.state.EventQueue) > 0 {
		event := g.state.EventQueue[0]
		g.state.EventQueue = g.state.EventQueue[1:]
		err := g.ProcessEvent(event)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *GameEngine) ProcessEvent(event Event) error {
	fmt.Println("ProcessEvent:", event.ID)
	if event.Depth > g.maxDepth {
		return fmt.Errorf("max depth of %d exceeded", g.maxDepth)
	}

	switch event.ID {
	case "destroy_card":
		err := g.DestroyCard(event)
		if err != nil {
			return err
		}
	}
	return nil
}
