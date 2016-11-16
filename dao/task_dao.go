package dao

import (
	"todoisAPI/models"
)

type TaskDAO struct {}

func (dao *TaskDAO) GetAll(email string) ([]models.Task, error)  {
	listTask := make([]models.Task,0)
	db, err := GetDBConnection()
	if(err != nil){
		return listTask, err
	}
	defer db.Close()

	query := "SELECT  t.id, t.description,t.complete, t.created_at, t.complete_at, u.id, u.name, u.email, u.active FROM todo_task as t inner join todo_user u on t.user_id = u.id	WHERE u.email = $1"
	dbQuery, err := db.Prepare(query)
	if(err != nil){
		return  listTask, err
	}
	defer dbQuery.Close()
	taskRows, err := dbQuery.Query(email)
	if(err != nil){
		return listTask, err
	}

	for taskRows.Next(){
		t := models.Task{}
		err = taskRows.Scan(&t.ID, &t.Description, &t.Complete,  &t.CreatedAt, &t.CompletedAt, &t.User.ID, &t.User.Name, &t.User.Email, &t.User.Active)
		if(err != nil){
			return  listTask, err
		}
		listTask = append(listTask, t)
	}

	return  listTask, nil
}

func (dao *TaskDAO) Create(t models.Task, email string) (models.Task, error)  {
	task := models.Task{}
	db, err := GetDBConnection()
	if(err != nil){
		return task, err
	}
	defer db.Close()

	query := "INSERT INTO todo_task(description, user_id)VALUES ($1, (SELECT u.id FROM todo_user as u WHERE u.email = $2)) RETURNING id "
	dbQuery, err := db.Prepare(query)
	if(err != nil){
		return  task, err
	}
	taskRow := dbQuery.QueryRow(t.Description, email)
	taskRow.Scan(&task.ID)

	query = "SELECT  t.id, t.description,t.complete, t.created_at, t.complete_at, u.id, u.name, u.email, u.active FROM todo_task as t inner join todo_user u on t.user_id = u.id WHERE t.id = $1"
	dbQuery, err = db.Prepare(query)
	if(err != nil){
		return  task, err
	}
	defer dbQuery.Close()

	taskRow = dbQuery.QueryRow(&task.ID)
	err = taskRow.Scan(&t.ID, &t.Description, &t.Complete,  &t.CreatedAt, &t.CompletedAt, &t.User.ID, &t.User.Name, &t.User.Email, &t.User.Active)
	if(err != nil){
		return task, err
	}

	return task, nil
}

func (dao *TaskDAO) Delete(id int) (error)  {
	db, err := GetDBConnection()
	if(err != nil){
		return err
	}
	defer db.Close()

	query := "DELETE FROM todo_task WHERE id=$1"
	dbQuery, err := db.Prepare(query)
	if(err != nil){
		return  err
	}

	_, err = dbQuery.Exec(id)
	if(err != nil){
		return  err
	}
	defer dbQuery.Close()

	return  nil
}

func (dao *TaskDAO) Update(id int) (error)  {
	db, err := GetDBConnection()
	if(err != nil){
		return err
	}
	defer db.Close()

	query := "UPDATE todo_task SET complete=true,complete_at=now() WHERE id = $1"
	dbQuery, err := db.Prepare(query)
	if(err != nil){
		return  err
	}

	_, err = dbQuery.Exec(id)
	if(err != nil){
		return  err
	}
	defer dbQuery.Close()

	return  nil
}