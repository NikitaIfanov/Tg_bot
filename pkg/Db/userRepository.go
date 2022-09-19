package Db

import (
	"fmt"
	"log"
)

type UserRepository struct {
	store *Store
}

func (r UserRepository) AddUser(ChatId int) {
	query := fmt.Sprintf("INSERT INTO Users (UserChatID) VALUES ('%d')", ChatId)
	if err := r.store.db.QueryRow(query).Scan(); err != nil {
		log.Print(err)
	}

}

func (r UserRepository) RemoveUser(ChatId int) {
	query := fmt.Sprintf("DELETE FROM Users WHERE (UserChatID='%d')", ChatId)
	if err := r.store.db.QueryRow(query).Scan(); err != nil {
		log.Print(err)
	}
}

func (r UserRepository) AddPair(ChatId int, file []byte) {
	query := fmt.Sprintf("INSERT INTO Users (Pairs) WHERE (UserChatID='%d') VALEUS('%q')", ChatId, file)
	if err := r.store.db.QueryRow(query).Scan(); err != nil {
		log.Print(err)
	}
}
