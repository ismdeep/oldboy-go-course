package main

import (
	"fmt"
	"strings"
	"time"
)

type Book struct {
	Id          string    /* 图书编号 */
	Name        string    /* 书名 */
	Cnt         int       /* 副本数量 */
	Author      string    /* 作者 */
	PublishDate time.Time /* 出版日期 */
	AddedDate   time.Time /* 上架日期 */
}

func NewBook(name string, cnt int, author string, publishDate time.Time) *Book {
	return &Book{
		Name:        name,
		Cnt:         cnt,
		Author:      author,
		PublishDate: publishDate,
	}
}

func (book Book) String() string {
	return fmt.Sprintf("{Id: \"%s\", Name: \"%s\", Author: \"%s\", Cnt: %d, PublishDate: \"%s\", AddedDate: \"%s\"}", book.Id, book.Name, book.Author, book.Cnt, book.PublishDate.Format("2006-01-02"), book.AddedDate.Format("2006-01-02"))
}

func (book *Book) QueryValid(bookName string, authorName string, publishDateBegin time.Time, publishDateEnd time.Time) bool {
	if "" != bookName && strings.Index(book.Name, bookName) < 0 {
		return false
	}
	if "" != authorName && strings.Index(book.Author, authorName) < 0 {
		return false
	}
	if book.PublishDate.Before(publishDateBegin) || book.PublishDate.After(publishDateEnd) {
		return false
	}
	return true
}