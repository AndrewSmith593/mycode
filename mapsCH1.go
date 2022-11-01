package main

import (
    "fmt"
)

func main() {

    languageMap := map[string]string{
        "Python": ".py",
        "Golang": ".go",
        "Java": ".java",
        "Ansible": ".yaml",
        "C++": ".cpp",
    }

    fmt.Printf("languageMap: %v\n", languageMap)

    delete(languageMap, "C++")

    languageMap["Julia"] = ".jl"

    fmt.Printf("languageMap: %v\n", languageMap)
}
