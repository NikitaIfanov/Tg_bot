package Db

import (
	"encoding/json"
	"log"
	"main.go/pkg/Exchange"
)

func Marshal(forDb *Exchange.ForDb) []byte {
	file, err := json.MarshalIndent(forDb, "", "    ")
	if err != nil {
		log.Print(err)
		return nil
	}
	return file
}
