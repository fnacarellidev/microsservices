// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package pgquery

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Diary struct {
	ID         uuid.UUID
	DiaryOwner uuid.UUID
	CreatedAt  pgtype.Timestamp
}

type Record struct {
	ID        uuid.UUID
	DiaryID   uuid.UUID
	Title     string
	Content   string
	CreatedAt pgtype.Timestamp
}

type User struct {
	ID       uuid.UUID
	Username string
	Password string
}
