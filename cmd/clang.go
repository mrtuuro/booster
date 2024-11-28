package cmd

import (
    "fmt"
    "os"
)

type CLang struct {
    ProjectName string
    Domain      string
    Version     string
}

func (l *CLang) makeFileWrite() string {
    s := fmt.Sprintf("TARGET = bin/main\n\nSRC = $(wildcard src/*.c)\nOBJ = $(patsubst src/%%.c, obj/%%.o, $(SRC))\n\nrun: clean default\n\t@./$(TARGET)\n\ndefault: $(TARGET)\n\nclean:\n\t@rm -f obj/*.o\n\t@rm -f bin/*\n\t@rm -f *.db\n\n$(TARGET):\n\t@gcc -o bin/main src/main.c")
    return s

}

func (l *CLang) CreateMakeFile(path string) error {
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

func (l *CLang) CreateMainFile(path string) error {
    err := os.Chdir("./src")
    if err != nil {
        return err
    }
    mainFile, err := os.OpenFile("main.c", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0744)
    if err != nil {
        return err
    }

    mainFileStr := l.mainFileWrite()
    _, err = mainFile.WriteString(mainFileStr)
    if err != nil {
        return fmt.Errorf("writing main.c: %v", err)
    }

    return nil
}

func (l *CLang) mainFileWrite() string {
    s := fmt.Sprintf("#include<stdio.h>\n\nint main(int argc, char* args[]) {\n\tprintf(\"Hello from %s\\n\");\n}", l.ProjectName)
    return s
}

func(l *CLang) createCLangDirs(path string) error {

    err := os.Mkdir(path + "/src", 0755)
    if err != nil {
        return err
    }
    err = os.Mkdir(path + "/obj", 0755)
    if err != nil {
        return err
    }
    err = os.Mkdir(path + "/bin", 0755)
    if err != nil {
        return err
    }

    return nil
}

func (l *CLang) Run() error {
    wd, err := os.Getwd()
    if err != nil {
        return fmt.Errorf("retrieving working directory: %v", err)
    }

    dirName := fmt.Sprintf("%s/%s", wd, l.ProjectName)
    err = os.Mkdir(dirName, 0755)
    if err != nil {
        return fmt.Errorf("creating project directory: %v", err)
    }

    err = l.createCLangDirs(dirName)
    if err != nil {
        return fmt.Errorf("creating C language dirs: %v", err)
    }

    err = os.Chdir(dirName)
    if err != nil {
        return fmt.Errorf("changing working directory: %v", err)
    }

    err = l.CreateMakeFile(dirName)
    if err != nil {
        return fmt.Errorf("creating makefile: %v", err)
    }

    err = l.CreateMainFile(dirName)
    if err != nil {
        return fmt.Errorf("creating main.c file: %v", err)

    }

    return nil
}
