package main

import (
    "fmt"

    "github.com/mrtuuro/booster/config"
)

func main() {

    cfg := config.NewConfig()
    fmt.Printf("projectName: %s\nlanguage: %s\ndomain: %s\n", cfg.ProjectName, cfg.Language, cfg.Domain)

    fmt.Println("We gonna boost you")
}
