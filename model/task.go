package model

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	Id        int    `json:"id"`
	Desc      string `json:"desc"`
	Completed bool   `json:"completed"`
}

type TasksList struct {
	Tasks []Task `json:"tasks"`
}

func GetTask(db *sql.DB) TasksList {
	qry := "select * from tasks"
	cur, err := db.Query(qry)
	if err != nil {
		panic(err)
	}

	tasks := TasksList{}

	for cur.Next() {
		tmp := Task{}
		err := cur.Scan(&tmp.Id, &tmp.Desc, &tmp.Completed)
		if err != nil {
			panic(err)
		}
		tasks.Tasks = append(tasks.Tasks, tmp)
	}
	return tasks
}

func AddTask(db *sql.DB, desc string, staus bool) (int64, error) {
	qry, err := db.Prepare("insert into tasks(desc,completed) Values(?,?)")
	if err != nil {
		panic(err)
	}
	op, er := qry.Exec(desc, staus)
	if er != nil {
		panic(er)
	}
	return op.LastInsertId()

}

func DeleteTask(db *sql.DB, id int) (int64, error) {
	qry, err := db.Prepare("delete from tasks where id=?")
	if err != nil {
		panic(err)
	}
	op, er := qry.Exec(id)
	if er != nil {
		panic(er)
	}
	return op.RowsAffected()
}

func UpdateStatus(db *sql.DB, id int) (int64, error) {
	stmt, e := db.Prepare("select completed from tasks where id=?")
	if e != nil {
		panic(e)
	}
	cur, e1 := stmt.Query(id)
	if e1 != nil {
		panic(e1)
	}
	var new_status bool

	for cur.Next() {
		tmp := Task{}
		err := cur.Scan(&tmp.Completed)
		if err != nil {
			panic(err)
		}
		new_status = !tmp.Completed
	}

	qry, err := db.Prepare("update tasks set completed=? where id=?")
	if err != nil {
		panic(err)
	}
	op, er := qry.Exec(new_status, id)
	if er != nil {
		panic(er)
	}
	return op.LastInsertId()
}
