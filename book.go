package main

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

func (book *Book) GetPageCount() int {
	return book.Pages
}
