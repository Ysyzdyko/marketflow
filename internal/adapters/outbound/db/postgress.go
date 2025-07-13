package db

import (
	"database/sql"
	"fmt"
	"log"
	"marketflow/internal/config"
)

type Postgres struct {
	db *sql.DB
	Repo
}

func NewPostgres(cfg config.DBConfig) *Postgres {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName,
	)

	db, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		log.Fatalf("Failed to open DB connection: %v", err)
	}

	log.Println("Connected to DB:", cfg.DBName)

	return &Postgres{
		db:   db,
		Repo: Repo{Conn: db},
	}
}

func (p *Postgres) Close() {
	if err := p.db.Close(); err != nil {
		log.Printf("Error closing the database connection: %v", err)
	}
}
