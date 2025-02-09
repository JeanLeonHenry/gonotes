/*
Copyright © 2025 Jean-Léon HENRY
*/
package main

import (
	"context"
	"database/sql"
	_ "embed"
	"log"

	"github.com/JeanLeonHenry/gonotes/cmd"
	"github.com/JeanLeonHenry/gonotes/db"
	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var ddl string

func main() {
	ctx := context.Background()

	// INFO: will create the file if it doesn't exist
	DB, err := sql.Open("sqlite", "test.db")
	if err != nil {
		log.Fatalln(err)
	}

	// // FIX: create tables, will error if they exist
	if _, err := DB.ExecContext(ctx, ddl); err != nil {
		log.Println(err)
	}
	queries := db.New(DB)
	cmd.InitCLI(queries, ctx)
	cmd.Execute()
}
