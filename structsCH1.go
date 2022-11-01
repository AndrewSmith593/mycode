package main

import (
    "fmt"
)

type Astro struct {
    name string
    age int
    mission string
    isneeded bool
}

func main(){
    buzz := Astro{"Buzz Lightyear", 42, "Infinity and beyond", false}
    woody := Astro{"Tom Hanks", 58, "Cowboy life", true}
    aldrin := Astro{"B. Aldrin", 79, "BAMF", true}

    fmt.Println(buzz)
    fmt.Println(woody)
    fmt.Println(aldrin)
}
