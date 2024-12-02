package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/uptrace/bun/driver/sqliteshim"
)

func main() {
	command := os.Args[1]
	switch command {
	case "migration":
		migration()
	case "migrate":
		migrate()
	}

}

func migrate() {
	db, err := sql.Open(sqliteshim.ShimName, "beholder.db")
	if err != nil {
		panic(err)
	}

	migrations, err := os.ReadDir("migrations")
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(migrations); i++ {
		name := migrations[i].Name()
		migration, err := os.ReadFile("migrations/" + name)
		if err != nil {
			panic(err)
		}
		sql := string(migration)

		db.Exec(sql)

	}

	db.Close()
}

func migration() {
	files, err := os.ReadDir("migrations")
	if err != nil {
		panic(err)
	}

	lastMigration := files[len(files)-1]

	lastID := strings.Split(lastMigration.Name(), "-")[0]
	integer, err := strconv.Atoi(lastID)
	if err != nil {
		panic(err)
	}
	cmd := exec.Command("touch", "migrations/"+fmt.Sprintf("0%d-%s.up.sql", integer+1, os.Args[2]))
	cmd.Run()
}
