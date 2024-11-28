package main

import (
    "log"
    "os"

    "github.com/mrtuuro/booster/cmd"
    "github.com/mrtuuro/booster/config"
)

func main() {

    cfg := config.NewConfig()
    l, err := cmd.New(cfg)
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }

    err = l.Run()
    if err != nil {
        log.Fatal(err)
    }
    return
}
