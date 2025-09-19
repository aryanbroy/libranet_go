package main

import (
	"fmt"
	"log"
	"time"

)

func main() {
	
	user := NewUser(1, "Aryan")

	fmt.Println()
	fmt.Println("Book: ")
	fmt.Println()

	book := NewBook(1, "Something", "Broy", 10)
	duration := 3 * 24 * time.Hour
	br, err := book.Borrow(user, duration)
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Printf("Book \"%s\" is borrowed by %s for %v days, with borrowing id: %s\n", book.Title, user.Name, duration.Hours()/24, br.ID)
	}

	fmt.Println()

	err = book.Return(user.ID)
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Printf("Book \"%s\" was returned by %s\n", book.Title, user.Name)
	}

	fmt.Println()
	fmt.Println("AudioBook:")
	fmt.Println()

	audioBook := NewAudioBook(1, "Some song", "Author1", 3 * time.Minute)
	ar, err := audioBook.Borrow(user, duration)
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Printf("AudioBook \"%s\" is borrowed by %s for %v days, with borrowing id: %s\n", audioBook.Title, user.Name, duration.Hours()/24, ar.ID)
	}

	audioBook.Play()
	fmt.Println()
	audioBook.Pause()

	err = audioBook.Return(user.ID)
		if err != nil {
		log.Panic(err)
	} else {
		fmt.Printf("AudioBook \"%s\" was returned by %s\n", audioBook.Title, user.Name)
	}

	fmt.Println()
	fmt.Println("Emagazine:")
	fmt.Println()

	user2 := NewUser(2, "Someone other than me")
	emagazine := NewEMagazine(1, "Times Now", "Me ofc", 10)
	emr, err := emagazine.Borrow(user2, duration)
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Printf("Emagazine \"%s\" is borrowed by %s for %v days, with borrowing id: %s\n", emagazine.Title, user2.Name, duration.Hours()/24, emr.ID)
	}

	err = emagazine.Return(user2.ID)
		if err != nil {
		log.Panic(err)
	} else {
		fmt.Printf("Emagazine \"%s\" was returned by %s\n", emagazine.Title, user2.Name)
	}
	fmt.Println()
	emagazine.ArchiveIssue(10)
}
