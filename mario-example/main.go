package main

import (
    "fmt"
    "models/mario"
)

func main() {

	m := mario.Player{3, 1, []string{"mushroom"}}

	fmt.Printf("Mario details: %v\n", m)

	m.Greenmushroom()

	m.Pickup("tanooki suit")

	m.CanWhistle()

	m.Pickup("whistle")

	m.CanWhistle()

}
