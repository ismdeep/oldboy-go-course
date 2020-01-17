package main

import (
	"fmt"
	"time"
)

type Rent struct {
	Id        string    /* 借阅编号 */
	StudentId string    /* 学生身份证 ID */
	BookId    string    /* 图书编号 */
	RentDate  time.Time /* 借出时间 */
	Status    int8      /* 图书借出状态 0已借出 1已归还 */
}

func (rent Rent) String() string {
	return fmt.Sprintf("{Id: \"%s\", StudentId: \"%s\", BookId: \"%s\", RentDate: \"%s\", Status: %d}", rent.Id, rent.StudentId, rent.BookId, rent.RentDate, rent.Status)
}
