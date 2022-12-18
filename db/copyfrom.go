// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: copyfrom.go

package db

import (
	"context"
)

// iteratorForCreateLunchMenus implements pgx.CopyFromSource.
type iteratorForCreateLunchMenus struct {
	rows                 []CreateLunchMenusParams
	skippedFirstNextCall bool
}

func (r *iteratorForCreateLunchMenus) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForCreateLunchMenus) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].Resturant,
		r.rows[0].Date,
		r.rows[0].Language,
		r.rows[0].Name,
		r.rows[0].Menu,
	}, nil
}

func (r iteratorForCreateLunchMenus) Err() error {
	return nil
}

func (q *Queries) CreateLunchMenus(ctx context.Context, arg []CreateLunchMenusParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"lunch_menus"}, []string{"resturant", "date", "language", "name", "menu"}, &iteratorForCreateLunchMenus{rows: arg})
}
