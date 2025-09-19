package main

import (
	"fmt"
	"time"
)
func (l *LibraryItem) ValidateBorrowing() bool{
	return l.IsAvailable
}

func (l *LibraryItem) Return(userId int64) error {
	var activeRecord *BorrowingRecord

	for _, record := range l.borrowRecord {
		if record.ReturnDate == nil && record.User.ID == userId {
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
	l.IsAvailable = true

	return nil
}

func (l *LibraryItem) Borrow(user *User, duration time.Duration) (*BorrowingRecord, error) {
	isAvailable := l.ValidateBorrowing()
	if !isAvailable {
		err := fmt.Errorf("Book Not available for borrowing")
		return nil, err
	}

	l.IsAvailable = false

	borrowDate := time.Now()
	dueDate := CalculateDueDate(borrowDate, duration)

	br := NewBorrowingRecord(l, user, borrowDate, dueDate)
	l.borrowRecord = append(l.borrowRecord, br)

	return br, nil
}

