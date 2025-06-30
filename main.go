package main

import (
	"fmt"
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
	Source *Card
	Target *Card
}

type GameEngine struct {
	state *GameState
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
			g.ExecuteAbility(card, ability)
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

func (g *GameEngine) ExecuteAbility(source *Card, ability Ability) {
	//check Ability Effect and handle appropriately
	switch ability.Effect {
	case "destroy_stack":
		for _, card := range g.state.Players["player2"].Stack {
			event := Event{
				ID:     "destroy card",
				Source: source,
				Target: card,
			}
			g.state.EventQueue = append(g.state.EventQueue, event)
		}
		break
	case "salvage_self":
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

	ge.state.Players["player1"].Hand = append(ge.state.Players["player1"].Hand, card1)

	ge.state.Players["player2"] = &Player{
		ID:        "player2",
		Hand:      make([]*Card, 0),
		Stack:     make([]*Card, 0),
		Graveyard: make([]*Card, 0),
	}

	ge.state.Players["player2"].Stack = append(ge.state.Players["player2"].Stack, card2)

	ge.PlayCard("player1", "1")
}
