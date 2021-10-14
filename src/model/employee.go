package model

import (
	"intro_web/src/utils"
)

// Employee 结构体
type Employee struct {
	Id           string
	Superior     string
	DepartmentId string
	Name         string
	BirthDate    string
	Address      string
	Sex          int8
	Salary       float64
}

func (e *Employee) Info() (err error) {
	// SQL
	sql := "SELECT * FROM employee WHERE id = ?"
	row := utils.Db.QueryRow(sql, e.Id)
	err = row.Scan(&e.Id, &e.Superior, &e.DepartmentId, &e.Name, &e.BirthDate, &e.Address, &e.Sex, &e.Salary)
	return
}
