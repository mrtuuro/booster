package config

import (
    "flag"
    "os"
)

type Config struct {
    ProjectName string
    Language    string
    Domain      string
}

func NewConfig() *Config {
    cfg := &Config{}
    cfg.Init()
    return cfg
}

func (cfg *Config) Init() {
    args := os.Args
    projectName := flag.String("name", "", "Name of the project to be created")
    language := flag.String("lang", "", "Language to be used with project")
    domain := flag.String("domain", "", "Github repo to the repository")
    flag.Parse()
    if len(args) < 2 {
        flag.Usage()
        os.Exit(0)
    }

    cfg.ProjectName = *projectName
    cfg.Language = *language
    cfg.Domain = *domain
}
