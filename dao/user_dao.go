package dao

import (
	"go-todo-apimodels"
)

type UserDAO struct {}

func (dao *UserDAO) GetById(id int) (*models.User, error)  {
	u := &models.User{}
	db, err := GetDBConnection()
	if(err != nil){
		return u, err
	}
	defer db.Close()


	query := "SELECT u.id, u.name, u.email, u.active FROM todo_user as u WHERE u.id = $1"
	dbQuery, err := db.Prepare(query)
	if(err != nil){
		return  u, err
	}
	defer dbQuery.Close()
	userRow := dbQuery.QueryRow(id)
	err = userRow.Scan(&u.ID, &u.Name, &u.Email, &u.Active)
	if(err != nil){
		return  u, err
	}

	return  u, nil
}

func (dao *UserDAO) Validate(data *models.LoginData) (*models.User, error){
	u := &models.User{}
	db, err := GetDBConnection()
	if(err != nil){
		return u, err
	}
	defer db.Close()

	query := "SELECT u.id, u.name, u.email, u.active FROM todo_user as u WHERE u.email = $1 AND  u.password = md5($2) AND u.active = true"
	dbQuery, err := db.Prepare(query)
	if(err != nil){
		return  u, err
	}
	defer dbQuery.Close()
	userRow := dbQuery.QueryRow(data.Email, data.Password)
	err = userRow.Scan(&u.ID, &u.Name, &u.Email, &u.Active)
	if(err != nil){
		return  u, err
	}

	return  u, nil
}
