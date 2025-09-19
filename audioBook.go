package main

import (
	"fmt"
	"time"
)

func NewAudioBook(id int64, title string, author string, duration time.Duration) *AudioBook{
	libItem := LibraryItem{
		ID: id,
		Title: title,
		Author: author,
		IsAvailable: true,
	}
	
	return &AudioBook{
		LibraryItem: libItem,
		isPlaying: false,
		Duration: duration,
	}
}

func (audioBook *AudioBook) Play() error{
	var isBorrowed bool
	var userName string 
	isBorrowed = false

	for _, borrowRecord := range audioBook.borrowRecord {
		if borrowRecord.ReturnDate == nil {
			userName = borrowRecord.User.Name
			isBorrowed = true
			break
		}
	}

	if isBorrowed == true {
		err := fmt.Errorf("Cannot play audioBook, it is borrowed by user: %v\n", userName)
		return err
	}
	audioBook.isPlaying = true	
	fmt.Printf("Playing audiobook, duration: %v\n", audioBook.Duration)
	return nil
}

func (audioBook *AudioBook) Pause() error{
	var isBorrowed bool
	var userName string 
	isBorrowed = false

	for _, borrowRecord := range audioBook.borrowRecord {
		if borrowRecord.ReturnDate == nil {
			userName = borrowRecord.User.Name
			isBorrowed = true
			break
		}
	}

	if isBorrowed == true {
		err := fmt.Errorf("Cannot pause audioBook, it is borrowed by user: %v\n", userName)
		return err
	}
	audioBook.isPlaying =  false
	fmt.Println("Paused audiobook")
	return nil
}
