package data

import (
	"context"

	"github.com/uptrace/bun"
)

type CrudRepository struct {
	Model interface{}
	db    *bun.DB
}

type QueryFilter struct {
	Query string
	Args  []interface{}
	Skip  *int
	Take  *int
}

func (r *CrudRepository) Create(dest ...interface{}) *error {
	_, err := r.db.NewInsert().Model(r.Model).Returning("*").Exec(context.Background(), dest...)
	return &err
}

func (r *CrudRepository) CreateWhere(input QueryFilter, dest ...interface{}) *error {
	_, err := r.db.NewUpdate().Model(r.Model).Where(input.Query, input.Args...).Returning("*").Exec(context.Background(), dest)
	return &err
}

func (r *CrudRepository) Update(dest ...interface{}) *error {
	_, err := r.db.NewUpdate().Model(r.Model).Returning("*").Exec(context.Background(), dest)
	return &err
}

func (r *CrudRepository) UpdateWhere(input QueryFilter, dest ...interface{}) *error {
	_, err := r.db.NewUpdate().Model(r.Model).Where(input.Query, input.Args...).Returning("*").Exec(context.Background(), dest)
	return &err
}

func (r *CrudRepository) SelectOne(input QueryFilter, dest ...interface{}) *error {
	_, err := r.db.NewSelect().Where(input.Query, input.Args...).Limit(1).Model(r.Model).Exec(context.Background(), dest)
	return &err
}

func (r *CrudRepository) Select(input QueryFilter, dest ...interface{}) *error {
	_, err := r.db.NewSelect().Where(input.Query, input.Args...).Model(r.Model).Exec(context.Background(), dest)
	return &err
}

func (r *CrudRepository) SelectCount(input QueryFilter, dest ...interface{}) *error {
	_, err := r.db.NewSelect().Where(input.Query, input.Args...).Model(r.Model).Count(context.Background())
	return &err
}

func (r *CrudRepository) SelectPaginated(input QueryFilter, dest ...interface{}) *error {
	_, err := r.db.NewSelect().Where(input.Query, input.Args...).Offset(*input.Skip).Limit(*input.Take).Model(r.Model).Exec(context.Background(), dest)
	return &err
}

func (r *CrudRepository) Delete(input QueryFilter, dest ...interface{}) *error {
	_, err := r.db.NewDelete().Where(input.Query, input.Args...).Model(r.Model).Exec(context.Background(), dest)
	return &err
}
