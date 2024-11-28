package cmd

import (
    "fmt"
    "os"
    "os/exec"
)

type LangGo struct {
    Version     string
    ProjectName string
    Domain      string
}

func (l *LangGo) makeFileWrite() string {
    s := fmt.Sprintf(".PHONY: all build run clean test\n\nPROJECT_NAME=%s\nbuild:\n\t@go build -o ./bin/$(PROJECT_NAME) ./*.go\n\nrun: build\n\t@clear\n\t@./bin/$(PROJECT_NAME)\n\nclean:\n\t@echo 'Cleanin up...'\n\t@go clean\n\t@rm -rf ./bin\n\ntest:\n\t@go test ./...", l.ProjectName)
    return s

}

func (l *LangGo) CreateMakeFile(path string) error {
    makeFile, err := os.OpenFile("Makefile", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0744)
    if err != nil {
        return err
    }

    makeFileStr := l.makeFileWrite()
    _, err = makeFile.WriteString(makeFileStr)
    if err != nil {
        return fmt.Errorf("writing makefile: %v", err)
    }

    return nil
}

func (l *LangGo) CreateMainFile(path string) error {
    mainFile, err := os.OpenFile("main.go", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0744)
    if err != nil {
        return err
    }

    mainFileStr := l.mainFileWrite()
    _, err = mainFile.WriteString(mainFileStr)
    if err != nil {
        return fmt.Errorf("writing makefile: %v", err)
    }

    return nil
}

func (l *LangGo) mainFileWrite() string {
    s := fmt.Sprintf("package main\nimport (\n\t\"fmt\"\n)\n\nfunc main(){\n\tfmt.Println(\"Hello from %s\")\n}\n", l.ProjectName)
    return s
}

func (l *LangGo) Run() error {
    wd, err := os.Getwd()
    if err != nil {
        return fmt.Errorf("retrieving working directory: %v", err)
    }

    dirName := fmt.Sprintf("%s/%s", wd, l.ProjectName)
    err = os.Mkdir(dirName, 0755)
    if err != nil {
        return fmt.Errorf("creating directory: %v", err)
    }

    err = os.Chdir(dirName)
    if err != nil {
        return fmt.Errorf("changing working directory: %v", err)
    }

    err = l.CreateMakeFile(dirName)
    if err != nil {
        return fmt.Errorf("creating makefile: %v", err)
    }

    command := exec.Command("go", "mod", "init", l.Domain)
    if err := command.Run(); err != nil {
        return fmt.Errorf("running go mod init command: %v", err)
    }

    err = l.CreateMainFile(dirName)
    if err != nil {
        return fmt.Errorf("creating main.go file: %v", err)

    }

    return nil
}
