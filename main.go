package main

import (
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/vinchauhan/go-task/db"
)

func main() {

	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.Init(dbPath)
	if err != nil {
		panic(err)
	}
}
