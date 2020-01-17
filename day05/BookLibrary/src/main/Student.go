package main

import "fmt"

type Student struct {
	Id     string /* 身份证ID */
	Name   string /* 姓名 */
	Grade  string /* 年级 */
	Gender string /* 性别 */
}

func (student Student) String() string {
	return fmt.Sprintf("{Id: \"%s\", Name: \"%s\", Grade: \"%s\", Gender: \"%s\"}", student.Id, student.Name, student.Grade, student.Gender)
}

func (student *Student) UpdateName(name string, grade string, gender string) {
	student.Name = name
	student.Grade = grade
	student.Gender = gender
}
