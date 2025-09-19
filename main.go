package main

import (
	"fmt"
	"log"
	"time"

)

func main() {
	
	user := NewUser(1, "Aryan")
	book := NewBook(1, "Something", "Broy", 10)

	duration := 3 * 24 * time.Hour
	err := book.Borrow(user.ID, duration)
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Printf("Book \"%s\" is borrowed by %s for %v days\n", book.Title, user.Name, duration.Hours()/24)
	}
}
