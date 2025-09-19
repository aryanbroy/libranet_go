package main

import (
	"time"
)

func NewBorrowingRecord(itemId int64, userId int64, borrowDate time.Time, dueDate time.Time) *BorrowingRecord{
	br := &BorrowingRecord{
		ItermId: itemId,
		UserId: userId,
		BorrowDate: borrowDate,
		DueDate: dueDate,
		ReturnDate: nil,
		Fine: 0,
	}

	return br
}

func CalculateDueDate(borrowDate time.Time, duration time.Duration) time.Time {
	dueDate := borrowDate.Add(duration)
	return dueDate
}
