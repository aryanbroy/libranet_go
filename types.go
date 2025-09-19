package main 

import "time"

type User struct {
	ID int64
	Name string
	MaxBorrow int
}

type BorrowingRecord struct {
	ItermId int64
	UserId int64
	BorrowDate time.Time
	DueDate time.Time
	ReturnDate *time.Time
	Fine float64
}

type LibraryItem struct {
	ID int64
	Title string
	Author string
	IsAvailable bool
	BorrowDate *time.Time
	ReturnDate *time.Time
}

type Book struct {
	LibraryItem
	Pages int
}

type AudioBook struct {
	LibraryItem
	Duration time.Duration
}

type EMagazine struct {
	LibraryItem
	IssueNumbers int
}
