package model

import "Hezzl_testing/internal/db"

const (
	createProjects = `create table if not exists projects (
    	id serial primary key,
    	name varchar(255) not null,
    	created_at timestamp
    )`

	createGoods = `create table if not exists goods (
    id serial primary key,
    project_id int references projects(id),
    name varchar(255) not null,
    description varchar(255),
    priority int,
    removed bool default false,
    created_at timestamp
	)`

	createIndexName = `create index if not exists index_name on goods(name)`
	createProject   = `insert into projects (id, name, created_at) values (1, 'first record', now())`
)

func InitSchema(s *db.Storage) error {
	tx := s.Db.MustBegin()

	tx.MustExec(createProjects)
	tx.MustExec(createGoods)
	tx.MustExec(createProject)
	tx.MustExec(createIndexName)

	return tx.Commit()
}
