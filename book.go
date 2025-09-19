package main

import (
	"fmt"
	"time"
)

func NewBook(id int64, title string, author string, pages int) *Book{
	libItem := LibraryItem{
		ID: id,
		Title: title,
		Author: author,
		IsAvailable: true,
		BorrowDate: nil,
		ReturnDate: nil,
	}
	
	return &Book{
		LibraryItem: libItem,
		Pages: pages,
	}
}

func (b *Book) Borrow(userId int64, duration time.Duration) error {
	isAvailable := b.ValidateBorrowing()
	if !isAvailable {
		err := fmt.Errorf("Book Not available for borrowing")
		return err
	}
	b.IsAvailable = false
	borrowDate := time.Now()
	dueDate := CalculateDueDate(borrowDate, duration)
	NewBorrowingRecord(b.ID, userId, borrowDate, dueDate)
	
	return nil
}

func (b *Book) Return(userId int64) error {
	return nil
}
