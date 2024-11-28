package cmd

import (
    "fmt"

    "github.com/mrtuuro/booster/config"
)

type ILanguage interface {
    Run() error
}

func New(cfg *config.Config) (ILanguage, error) {
    switch cfg.Language {
    case "go":
        l := &LangGo{
            Version:     cfg.Version,
            ProjectName: cfg.ProjectName,
            Domain:      cfg.Domain,
        }
        return l, nil
    case "c":
        l := &CLang{
            Version:     cfg.Version,
            ProjectName: cfg.ProjectName,
            Domain:      cfg.Domain,
        }
        return l, nil
    default:
        return nil, fmt.Errorf("no other languages are supported at this version")
    }


}
