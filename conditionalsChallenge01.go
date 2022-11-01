package main 

import (
    "fmt"
    "time"
    "math/rand"
)

func main() {

    rand.Seed(time.Now().UnixNano())

    flavors := []string{"choco", "razzberry blitz", "vanilla cream dream", "obama type"}

    if choice := rand.Intn(3); choice == 3 {
        fmt.Printf("Yummm you picked the flavor %v .\n", flavors[3])
    } else if choice == 2 {
        fmt.Printf("Yummm you picked the flavor %v.\n", flavors[2])   
    } else if choice == 1 {
        fmt.Printf("Yummm you picked the flavor %v.\n", flavors[1])
    } else {
        fmt.Printf("Yummm you picked the flavor %v.\n", flavors[0])
    }
}
