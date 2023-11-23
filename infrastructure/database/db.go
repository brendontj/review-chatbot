package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

type Database struct {
	conn *pgx.Conn
}

func New() *Database {
	return &Database{}
}

func (d *Database) Connect() {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:pg123@localhost:5432/chatbot")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	d.conn = conn
}

func (d *Database) Disconnect() {
	d.conn.Close(context.Background())
}
