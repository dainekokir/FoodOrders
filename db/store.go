package dbDef

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

type DbStore struct {
	DB *pgx.Conn
}

func InitDb() *DbStore {
	return &DbStore{}
}

func (store *DbStore) Open(DATABASE_URL string) {
	var err error
	store.DB, err = pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
}

func (store *DbStore) Close() {
	store.DB.Close(context.Background())
}

func (store *DbStore) AddRequest(smg string) {
	rows, err := store.DB.Query(context.Background(), `insert into public."request" (text) values ($1)`, smg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "add to database: %v\n", err)
	}
	defer rows.Close()
}
