package main

func (l *LibraryItem) ValidateBorrowing() bool{
	return l.IsAvailable
}
