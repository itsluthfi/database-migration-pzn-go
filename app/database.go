package app

import (
	"database/sql"
	"rest-api-pzn-go/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/migration_go_db")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

	// create migrate
	// migrate create -ext sql -dir db/migrations create_table_first
	// migrate create -ext sql -dir db/migrations create_table_second
	// migrate create -ext sql -dir db/migrations create_table_third
	// migrate create -ext sql -dir db/migrations sample_dirty_state

	// run migrate
	// migrate -database 'mysql://root:root@tcp(localhost:3306)/migration_go_db' -path db/migrations up
	// migrate -database 'mysql://root:root@tcp(localhost:3306)/migration_go_db' -path db/migrations down

	// force rollback
	// migrate -database 'mysql://root:root@tcp(localhost:3306)/migration_go_db' -path db/migrations version
	// migrate -database 'mysql://root:root@tcp(localhost:3306)/migration_go_db' -path db/migrations force 20250213063547
}
