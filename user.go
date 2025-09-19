package main

const maxBorrow = 10

func NewUser(id int64, name string) *User{
	return &User{
		ID: id,
		Name: name,	
		MaxBorrow: maxBorrow,
	}
}
