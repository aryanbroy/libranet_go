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
	}
	
	return &Book{
		LibraryItem: libItem,
		Pages: pages,
	}
}

func (b *Book) Return(userId int64) error {
	var activeRecord *BorrowingRecord

	for _, record := range b.borrowRecord {
		if record.ReturnDate == nil && record.UserId == userId {
			activeRecord = record
			break
		}
	}

	if activeRecord == nil {
		err := fmt.Errorf("No borrowing record found")
		return err
	}

	duration := 4 * 24 * time.Hour
	returnDate := time.Now().Add(duration)

	activeRecord.ReturnDate = &returnDate
	activeRecord.Fine = CalculateFine(activeRecord.DueDate, returnDate)	
	b.IsAvailable = true

	return nil
}

func (b *Book) Borrow(userId int64, duration time.Duration) (*BorrowingRecord, error) {
	isAvailable := b.ValidateBorrowing()
	if !isAvailable {
		err := fmt.Errorf("Book Not available for borrowing")
		return nil, err
	}

	b.IsAvailable = false

	borrowDate := time.Now()
	dueDate := CalculateDueDate(borrowDate, duration)
	br := NewBorrowingRecord(b.ID, userId, borrowDate, dueDate)
	b.borrowRecord = append(b.borrowRecord, br)	

	return br, nil
}

