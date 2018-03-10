package main

import (
	"time"
	"errors"
	_ "github.com/lib/pq"
	"database/sql"
	"github.com/lib/pq"
)

type Student struct {
	ID        int
	Name      string
	Age       int16
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Create(s Student) error {
	query := `INSERT INTO 
				estudiantes (name, age, active)
				values($1, $2, $3)`
	db := GetConnection()
	defer db.Close()

	ps, err := db.Prepare(query)
	if err != nil{
		return err
	}
	defer ps.Close()

	r, err := ps.Exec(s.Name, s.Age, s.Active)
	if err != nil{
		return err
	}

	i, _ := r.RowsAffected()
	if i != 1{
		return errors.New("se esperaba una fila afectada")
	}
	return nil
}

func GetAll() (listStudents []Student, err error){
	var query string
	query = "SELECT * FROM estudiantes"

	db := GetConnection()
	defer db.Close()

	ps, err := db.Prepare(query)
	if err != nil{
		return listStudents, err
	}

	rs, err := ps.Query()
	if err != nil{
		return listStudents, err
	}
	defer rs.Close()

	var name sql.NullString
	var age sql.NullInt64
	var active sql.NullBool
	var created pq.NullTime
	var updated pq.NullTime

	for rs.Next(){
		student := Student{}
		err = rs.Scan(
			&student.ID,
			&name,
			&age,
			&active,
			&created,
			&updated,
		)

		if err != nil{
			return listStudents, err
		}

		student.Name = name.String
		student.Age = int16(age.Int64)
		student.Active = active.Bool
		student.CreatedAt = created.Time
		student.UpdatedAt = updated.Time

		listStudents = append(listStudents, student)

	}

	return listStudents, err

}

func GetByID(id int) (student Student, err error){
	var query string
	query = `SELECT * FROM ESTUDIANTES WHERE id = $1`
	db := GetConnection()
	defer db.Close()
	ps, err := db.Prepare(query)
	if err != nil{
		return
	}

	rs, err := ps.Query(id)
	defer rs.Close()

	if err != nil{
		return
	}

	name := sql.NullString{}
	age := sql.NullInt64{}
	active := sql.NullBool{}
	created := pq.NullTime{}
	updated := pq.NullTime{}


	if rs.Next(){
		err = rs.Scan(
			&student.ID,
			&name,
			&age,
			&active,
			&created,
			&updated,
		)
	}

	if err != nil{
		return
	}

	student.Name = name.String
	student.Age = int16(age.Int64)
	student.CreatedAt = created.Time
	student.UpdatedAt = updated.Time

	return
}

func Delete(id int) (err error){
	query := `DELETE from estudiantes where id = $1`
	db := GetConnection()
	defer db.Close()

	ps, err := db.Prepare(query)
	if err != nil{
		return
	}

	result , err := ps.Exec(id)
	if err != nil{
		return
	}

	rows, err := result.RowsAffected()
	if rows != 1{
		return errors.New("se esperaba 1 fila afectada")
	}
	return
}
