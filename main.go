package main

import (
    "log"

    "github.com/mrtuuro/booster/cmd"
    "github.com/mrtuuro/booster/config"
)

func main() {

    cfg := config.NewConfig()
    l := cmd.New(cfg)

    err := l.CreateDir()
    if err != nil {
        log.Fatal(err)
    }
    return
}
