package repository

import (
	"Hezzl_testing/internal/db"
	"Hezzl_testing/internal/service/dto"
	"database/sql"
	"errors"
	"log"
	"time"
)

const (
	createGoodQuery = `insert into
    goods (project_id, name, description, priority, created_at)
	values ($1, $2, $3, $4, $5)
	returning id, project_id, name, priority, removed, created_at`

	countMaxPriorityQuery = `select COALESCE(MAX(priority), 0)+1 from goods`

	selectForGoodQuery = `select name, description from goods where id=$1 and project_id=$2 for update`

	updateForGoodQuery = `update goods set name=$1, description=$2 where id=$3 and project_id=$4
 				returning id, project_id, name, description, priority, removed, created_at`

	deleteGoodQuery = `update goods set removed=true where id=$1 and project_id=$2 returning id, project_id, removed`

	getGoodsQuery = `select id, project_id, name, COALESCE(description),priority,removed,created_at from goods limit $1 offset $2`

	updatePriorityQuery = `update goods set priority=$1 where id=$2 and project_id=$3 order by id asc returning id, priority`
	patchGoodQuery      = `update goods set priority=$1 where id=$2 and project_id=$3 returning id, priority`
	patchGoodsQuery     = `update goods set priority=priority+1 where priority>=$1 and id<>$2`

	countRemovedQuery = `select COUNT(id) from goods where removed=true`
	countAllQuery     = `select COUNT(id) from goods`
)

type ProjectGoodRepo struct {
	Db db.DB
}

func NewRepository(db db.DB) *ProjectGoodRepo {
	return &ProjectGoodRepo{Db: db}
}

func (repo *ProjectGoodRepo) Create(projectId int, name string) (*dto.CreateGoodResponse, error) {
	var good dto.CreateGoodResponse
	now := time.Now()
	description := ""
	priority, err := repo.maxPriority()
	if err != nil {
		return nil, err
	}

	if err := repo.Db.QueryRowx(
		createGoodQuery,
		projectId,
		name,
		description,
		priority,
		now,
	).Scan(
		&good.Id,
		&good.ProjectId,
		&good.Name,
		&good.Priority,
		&good.Removed,
		&good.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &good, nil
}

func (repo *ProjectGoodRepo) Update(id, projectId int, name, description string) (*dto.UpdateGoodResponse, error) {
	var good dto.UpdateGoodResponse
	tx := repo.Db.MustBegin()

	defer tx.Rollback()

	if _, err := tx.Exec(selectForGoodQuery, id, projectId); err != nil {
		return nil, err
	}

	if err := tx.QueryRowx(
		updateForGoodQuery,
		name,
		description,
		id,
		projectId,
	).Scan(
		&good.Id,
		&good.ProjectId,
		&good.Name,
		&good.Description,
		&good.Priority,
		&good.Removed,
		&good.CreatedAt,
	); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &good, nil
}

func (repo *ProjectGoodRepo) Delete(id, projectId int) (*dto.DeleteGoodResponse, error) {
	var resp dto.DeleteGoodResponse

	if err := repo.Db.QueryRowx(deleteGoodQuery, id, projectId).Scan(&resp.Id, &resp.ProjectId, &resp.Removed); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (repo *ProjectGoodRepo) Get(limit, offset int) (*dto.GetGoodsResponse, error) {
	rows, err := repo.Db.Query(getGoodsQuery, limit, offset-1)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var resp dto.GetGoodsResponse

	for rows.Next() {
		var good dto.Good
		if err := rows.Scan(
			&good.Id,
			&good.ProjectId,
			&good.Name,
			&good.Description,
			&good.Priority,
			&good.Removed,
			&good.CreatedAt,
		); err != nil {
			return nil, err
		}

		resp.Goods = append(resp.Goods, good)
	}

	countRemoved, err := repo.countRemoved()
	if err != nil {
		return nil, err
	}

	all, err := repo.countAll()
	if err != nil {
		return nil, err
	}

	resp.Meta.Total = all
	resp.Meta.Removed = countRemoved
	resp.Meta.Limit = limit
	resp.Meta.Offset = offset

	return &resp, nil
}

func (repo *ProjectGoodRepo) PatchPriority(id, projectId, newPriority int) (*dto.ReprioritizeResponse, *dto.UpdateGoodsResponse, error) {
	tx := repo.Db.MustBegin()

	defer tx.Rollback()

	var resp dto.ReprioritizeResponse

	if err := tx.QueryRowx(patchGoodQuery, newPriority, id, projectId).Scan(&resp.Id, &resp.Priority); err != nil {
		log.Println(err)
		return nil, nil, err
	}

	rows, err := tx.Query(patchGoodsQuery, newPriority, id)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	defer rows.Close()

	var updatedGoods dto.UpdateGoodsResponse

	for rows.Next() {
		var good dto.Good
		if err := rows.Scan(
			&good.Id,
			&good.ProjectId,
			&good.Name,
			&good.Description,
			&good.Priority,
			&good.Removed,
			&good.CreatedAt,
		); err != nil {
			log.Println(err)
			return nil, nil, err
		}

		updatedGoods.Goods = append(updatedGoods.Goods, good)
	}

	if err := tx.Commit(); err != nil {
		return nil, nil, err
	}

	return &resp, &updatedGoods, nil
}

func (repo *ProjectGoodRepo) countRemoved() (int, error) {
	var count int
	if err := repo.Db.QueryRowx(countRemovedQuery).Scan(&count); err != nil {
		log.Println(err)
		return 0, err
	}

	return count, nil
}

func (repo *ProjectGoodRepo) countAll() (int, error) {
	var count int
	if err := repo.Db.QueryRowx(countAllQuery).Scan(&count); err != nil {
		log.Println(err)
		return 0, err
	}

	return count, nil
}

func (repo *ProjectGoodRepo) maxPriority() (int, error) {
	var count int

	if err := repo.Db.QueryRowx(countMaxPriorityQuery).Scan(&count); err != nil && errors.Is(err, sql.ErrNoRows) {
		log.Println(err)
		return 0, err
	}

	return count, nil
}
