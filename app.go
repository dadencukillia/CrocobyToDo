package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

type Task struct {
	Id   int
	Desc string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// TO-DO
func getNewReleativePath(file string) string {
	str, _ := os.Executable()
	return filepath.Join(filepath.Dir(str), strings.ReplaceAll(file, "\\", "/"))
}

func (a *App) SaveTasks(content string) bool {
	err := os.WriteFile(getNewReleativePath("tasks.data"), []byte(content), 0644)
	if err == nil {
		fmt.Println("Succeful writed: " + getNewReleativePath("tasks.data"))
		return false
	}
	return true
}

func (a *App) LoadTasks() string {
	cont, err := os.ReadFile(getNewReleativePath("tasks.data"))
	if err != nil {
		return ""
	}
	return string(cont)
}
