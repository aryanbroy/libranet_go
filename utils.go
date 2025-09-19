package main

import (
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
)

func NewBorrowingRecord(itemId int64, userId int64, borrowDate time.Time, dueDate time.Time) *BorrowingRecord{
	id := uuid.New().String()

	br := &BorrowingRecord{
		ID: id,
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
	fmt.Println("DueDate: ",dueDate)
	return dueDate
}

func CalculateFine(dueDate, returnDate time.Time) float64 {
	var fine float64 

	if returnDate.Before(dueDate) {
		fmt.Println("Returned on or before due date")
		fmt.Println("0 fine implemented")
		return 0
	}

	fmt.Println("Returned after due date")
	duration := returnDate.Sub(dueDate)
	days := math.Round(duration.Hours()/24)
	fine = days * 10
	fmt.Printf("Returned after %v days, fine amount implemented: %v\n", days, fine)
	return fine
}
