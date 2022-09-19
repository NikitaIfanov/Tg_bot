package Db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

func (s *Store) Open() error {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/Db")
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Store) Close() error {
	if err := s.db.Close(); err != nil {
		return err
	}
	return nil
}

func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}
