package main

import (
	"fmt"
)

type Player struct {
	Lives     int
	Stage     int
	Inventory []string
}

func (p *Player) Greenmushroom() {
	p.Lives++
	fmt.Printf("Mario got a mushroom and now has %v lives!\n", p.Lives)
}

func (p *Player) Pickup(item string) {
	fmt.Printf("Mario's inventory: %v\n", p.Inventory)

	p.Inventory = append(p.Inventory, item)
	fmt.Printf("Mario picked up a %v!\n", item)

	fmt.Printf("Mario's inventory: %v\n", p.Inventory)
}

func (p Player) CanWhistle() bool {
	if p.Contains("whistle") {
		return true
	}
	return false
}

func (p Player) Contains(s string) bool {
	for _, item := range p.Inventory {
		if item == s {
			fmt.Println("Mario CAN whistle.")
			return true
		}
	}
	fmt.Println("Mario CANNOT whistle.")
	return false
}

func main() {

	mario := Player{3, 1, []string{"mushroom"}}

	fmt.Printf("Mario details: %v\n", mario)

	mario.Greenmushroom()

	mario.Pickup("tanooki suit")

	mario.CanWhistle()

	mario.Pickup("whistle")

	mario.CanWhistle()

}

