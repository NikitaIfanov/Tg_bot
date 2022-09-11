package Db

import (
	"fmt"
)

func Show(db map[string]bool) string {

	msg := ""
	for key, value := range db {
		if value == true {
			msg += fmt.Sprintf("%s\n", key)
		}
	}
	return msg
}
