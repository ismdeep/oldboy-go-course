package main

import (
	"fmt"
	tsgutils "github.com/typa01/go-utils"
	"time"
)

func main() {
	var bookLibrary = BookLibrary{
		Books:    []Book{},
		Rents:    []Rent{},
		Students: []Student{},
	}
	/* Add Books*/
	bookLibrary.AddBook(Book{
		Id:          "9780596550714",
		Name:        "The C Programming Language",
		Cnt:         10,
		Author:      "K&R",
		PublishDate: NewDate("2013-01-02"),
		AddedDate:   time.Now(),
	})
	bookLibrary.AddBook(Book{
		Id:          "9781305075771",
		Name:        "Invitation to Computer Science (7th)",
		Cnt:         10,
		Author:      "G. Michael Schneider",
		PublishDate: NewDate("2015-02-04"),
		AddedDate:   time.Time{},
	})

	/* Add Students */
	bookLibrary.AddStudent(Student{
		Id:     "362334199108191211",
		Name:   "Del Cooper",
		Grade:  "2018",
		Gender: "m",
	})
	
	bookLibrary.AddStudent(Student{
		Id:     "362334199902191214",
		Name:   "John Levy",
		Grade:  "2018",
		Gender: "m",
	})

	/* Add Rents */
	bookLibrary.AddRent(Rent{
		Id:        tsgutils.GUID(),
		StudentId: "362334199108191211",
		BookId:    "9781305075771",
		RentDate:  time.Now(),
		Status:    0,
	})
	bookLibrary.AddRent(Rent{
		Id:        tsgutils.GUID(),
		StudentId: "362334199108191211",
		BookId:    "9780596550714",
		RentDate:  time.Now(),
		Status:    0,
	})
	bookLibrary.AddRent(Rent{
		Id:        tsgutils.GUID(),
		StudentId: "362334199902191214",
		BookId:    "9781305075771",
		RentDate:  time.Now(),
		Status:    0,
	})
	bookLibrary.AddRent(Rent{
		Id:        tsgutils.GUID(),
		StudentId: "362334199902191214",
		BookId:    "9780596550714",
		RentDate:  time.Now(),
		Status:    0,
	})
	//fmt.Println(bookLibrary.QueryBooks("Computer","", NewDate("1970-01-01"),time.Now()))
	//fmt.Println(bookLibrary.QueryBooks("","K&R", NewDate("1970-01-01"),time.Now()))
	//fmt.Println(bookLibrary)
	fmt.Println(bookLibrary.QueryStudentRents("362334199108191211"))
}
