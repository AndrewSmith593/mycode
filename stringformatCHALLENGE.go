package main

import (
    "fmt"
)

func main(){
    uri := "https://example.org:6001/v2/snacks?"
    r, q, s := "req=snickers", "quantity=1", "size=king"

    url := fmt.Sprintf("%s%s%s&%s", uri, r, q, s)

    fmt.Println(url)
}
