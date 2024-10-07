package main

import (
	"fmt"

	"github.com/per1cycle/miui/tui"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func openSqlite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("miui.db"), &gorm.Config{})
	if err != nil {
		return db, fmt.Errorf("unable to open database: %w", err)
	}
	return db, nil
}

func main() {
	db, err := openSqlite()
	tui.TuiFoo()
}