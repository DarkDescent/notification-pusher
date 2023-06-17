-- create new value
-- name: CreateFirebaseToken :one
insert into firebase_token(token, cid, active, expiresAt) values($1, $2, $3, $4) returning *;

-- so we can set token as inactive
-- name: SetActivitity :one
update firebase_token set active=$1 where id=$2 returning *;

-- find token by external cid
-- name: FindTokenByCid :one
select * from firebase_token where cid=$1;

-- find token information by token and cid (for example to find if we already have active token for this pair)
-- requesting by cid and token, in this order, because of index
-- name: FindDuplicatedToken :one
select * from firebase_token where cid=$1 and token=$2;

-- find tokens that must be expired
-- name: FindExpiredToken :many
select * from firebase_token where active=true and expiresat < $1;