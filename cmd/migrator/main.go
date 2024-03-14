package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var storagePath, migrationsPath string

	flag.StringVar(&storagePath, "storage-path", "", "path to database")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.Parse()

	if storagePath == "" {
		panic("storage-path is required")
	}
	if migrationsPath == "" {
		panic("migrations-path is required")
	}
	migratePath := "file://" + migrationsPath
	storPath := fmt.Sprintf("sqlite3://%s", storagePath)
	fmt.Printf("\n migrate Path: %s \n stor path: %s \n", migratePath, storPath)

	m, err := migrate.New(
		migratePath,
		storPath,
	)
	if err != nil {
		panic(err)
	}
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")
			return
		}
		panic(err)
	}
	fmt.Println("Migrations apply")
}
