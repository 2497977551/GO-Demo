package main

import (
	"log"
)

func getOne(name string) (test, error) {
	t := test{}
	err := db.QueryRow("SELECT * FROM test WHERE name=?", name).Scan(&t.id, &t.name, &t.sex)

	return t, err
}

func getWhole() (tests []test, err error) {
	rows, err := db.Query("SELECT * FROM test")
	for rows.Next() {
		t := test{}
		err = rows.Scan(&t.id, &t.name, &t.sex)
		if err != nil {
			log.Fatalln(err.Error())
		}
		tests = append(tests, t)
	}
	return
}

func (t *test) update() (err error) {
	_, err = db.Exec("UPDATE test SET name = ? WHERE id = ?", t.name, t.id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return
}
