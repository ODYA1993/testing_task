package store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

// Store ...
type Store struct {
	config *ConfigDB
	db     *sql.DB
}

// NewStore ...
func NewStore(config *ConfigDB) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		s.config.Host, s.config.Port, s.config.UserName, s.config.Password, s.config.DBName, s.config.SslMode)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

// Close ...
func (s *Store) Close() {
	err := s.db.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
