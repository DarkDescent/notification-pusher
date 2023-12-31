// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFirebaseToken = `-- name: CreateFirebaseToken :one
insert into firebase_token(token, cid, active, expiresAt) values($1, $2, $3, $4) returning id, token, cid, active, expiresat
`

type CreateFirebaseTokenParams struct {
	Token     string
	Cid       uuid.UUID
	Active    bool
	Expiresat time.Time
}

// create new value
func (q *Queries) CreateFirebaseToken(ctx context.Context, arg CreateFirebaseTokenParams) (FirebaseToken, error) {
	row := q.db.QueryRowContext(ctx, createFirebaseToken,
		arg.Token,
		arg.Cid,
		arg.Active,
		arg.Expiresat,
	)
	var i FirebaseToken
	err := row.Scan(
		&i.ID,
		&i.Token,
		&i.Cid,
		&i.Active,
		&i.Expiresat,
	)
	return i, err
}

const findDuplicatedToken = `-- name: FindDuplicatedToken :one
select id, token, cid, active, expiresat from firebase_token where cid=$1 and token=$2
`

type FindDuplicatedTokenParams struct {
	Cid   uuid.UUID
	Token string
}

// find token information by token and cid (for example to find if we already have active token for this pair)
// requesting by cid and token, in this order, because of index
func (q *Queries) FindDuplicatedToken(ctx context.Context, arg FindDuplicatedTokenParams) (FirebaseToken, error) {
	row := q.db.QueryRowContext(ctx, findDuplicatedToken, arg.Cid, arg.Token)
	var i FirebaseToken
	err := row.Scan(
		&i.ID,
		&i.Token,
		&i.Cid,
		&i.Active,
		&i.Expiresat,
	)
	return i, err
}

const findExpiredToken = `-- name: FindExpiredToken :many
select id, token, cid, active, expiresat from firebase_token where active=true and expiresat < $1
`

// find tokens that must be expired
func (q *Queries) FindExpiredToken(ctx context.Context, expiresat time.Time) ([]FirebaseToken, error) {
	rows, err := q.db.QueryContext(ctx, findExpiredToken, expiresat)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FirebaseToken
	for rows.Next() {
		var i FirebaseToken
		if err := rows.Scan(
			&i.ID,
			&i.Token,
			&i.Cid,
			&i.Active,
			&i.Expiresat,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findTokenByCid = `-- name: FindTokenByCid :one
select id, token, cid, active, expiresat from firebase_token where cid=$1
`

// find token by external cid
func (q *Queries) FindTokenByCid(ctx context.Context, cid uuid.UUID) (FirebaseToken, error) {
	row := q.db.QueryRowContext(ctx, findTokenByCid, cid)
	var i FirebaseToken
	err := row.Scan(
		&i.ID,
		&i.Token,
		&i.Cid,
		&i.Active,
		&i.Expiresat,
	)
	return i, err
}

const setActivitity = `-- name: SetActivitity :one
update firebase_token set active=$1 where id=$2 returning id, token, cid, active, expiresat
`

type SetActivitityParams struct {
	Active bool
	ID     int64
}

// so we can set token as inactive
func (q *Queries) SetActivitity(ctx context.Context, arg SetActivitityParams) (FirebaseToken, error) {
	row := q.db.QueryRowContext(ctx, setActivitity, arg.Active, arg.ID)
	var i FirebaseToken
	err := row.Scan(
		&i.ID,
		&i.Token,
		&i.Cid,
		&i.Active,
		&i.Expiresat,
	)
	return i, err
}
