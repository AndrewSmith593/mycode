package main

import (
    "fmt"
)

type Astro struct {
    name     string
    age      int
    mission  string
    isneeded bool
}

func main() {

    ast1 := Astro{"Megan McArthur", 35, "ISS", false}
    ast2 := Astro{"Barry Wilmore", 41, "Boeing Crew Flight Test (CFT)", true}
    ast3 := Astro{"Raja Chari", 39, "SpaceX Crew-3", true}

    p := []Astro{ast1, ast2, ast3}

    fmt.Println(p)

    fmt.Println(p[2].mission)
}
