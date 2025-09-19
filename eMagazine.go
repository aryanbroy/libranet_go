package main

import "fmt"

func NewEMagazine(id int64, title string, author string, issueNumber int) *EMagazine{
	libItem := LibraryItem{
		ID: id,
		Title: title,
		Author: author,
		IsAvailable: true,
	}
	
	return &EMagazine{
		LibraryItem: libItem,
		IssueNumber: issueNumber,
		IsArchive: false,
	}
}

func (audioBook *AudioBook) ArchiveIssue(issueNumber int) {
	fmt.Printf("Archieved issue with issue number: %v\n", issueNumber)
}
