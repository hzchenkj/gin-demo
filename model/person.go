package model

import (
	"fmt"
	db "gin-demo/database"
	"log"
)

type Person struct {
	Id        int    `json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

func (p *Person) AddPerson() (id int64, err error) {
	sql := "insert into person(first_name,last_name) values (?,?)"
	rs, err := db.MySqlDB.Exec(sql, p.FirstName, p.LastName)
	if err != nil {
		log.Fatalln(err)
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (p *Person) GetPerson() (person Person, err error) {
	sql := "select id,first_name,last_name from person where id=?"
	row := db.MySqlDB.QueryRow(sql, p.Id)
	err = row.Scan(&person.Id, &person.FirstName, &person.LastName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

func (p *Person) GetPersons() (persons []Person, err error) {
	persons = make([]Person, 0)
	sql := "select id,first_name,last_name from person"
	rows, err := db.MySqlDB.Query(sql)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		log.Fatalln(err)
		return
	}
	return
}

func (p *Person) DelPerson() (ra int64, err error) {
	rs, err := db.MySqlDB.Exec("delete from person where id=?", p.Id)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return
}

func (p *Person) UpdatePerson() (ra int64, err error) {
	rs, err := db.MySqlDB.Exec("update person set first_name=?,last_name=? where id=?", p.FirstName, p.LastName, p.Id)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return
}
