package main

import (
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func main() {

	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.Init(dbPath)
}
