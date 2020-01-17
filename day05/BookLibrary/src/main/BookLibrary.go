package main

import (
	"time"
)

type BookLibrary struct {
	Books    []Book
	Rents    []Rent
	Students []Student
}

func (bookLib *BookLibrary) AddBook(book Book) {
	bookLib.Books = append(bookLib.Books, book)
}

func (bookLib *BookLibrary) AddRent(rent Rent) {
	bookLib.Rents = append(bookLib.Rents, rent)
}

func (bookLib *BookLibrary) AddStudent(student Student) {
	bookLib.Students = append(bookLib.Students, student)
}

func (bookLib *BookLibrary) QueryBooks(bookName string, authorName string, publishDateBegin time.Time, publishDateEnd time.Time) []Book {
	var books = []Book{}
	for _, book := range bookLib.Books {
		if book.QueryValid(bookName, authorName, publishDateBegin, publishDateEnd) {
			books = append(books, book)
		}
	}
	return books
}

func (bookLib *BookLibrary) QueryStudentRents(studentId string) []Rent {
	var rents = []Rent{}
	for _, rent := range bookLib.Rents {
		if rent.StudentId == studentId {
			rents = append(rents, rent)
		}
	}
	return rents
}