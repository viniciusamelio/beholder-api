package infra

import (
	"database/sql"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func GetDB() *bun.DB {
	sqldb, err := sql.Open(sqliteshim.ShimName, "file:beholder.db?cache=shared")
	if err != nil {
		panic(err)
	}

	sqldb.Exec("INSERT INTO organizations(id,name) VALUES(1,'teste')")

	db := bun.NewDB(sqldb, sqlitedialect.New())
	return db
}

func RunMigrations() {
	migrations, err := os.ReadDir("migrations")
	if err != nil {
		panic(err)
	}

	db := GetDB()

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
