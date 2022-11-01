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
