package types

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Avatar struct {
	Id         pgtype.UUID
	Uname      pgtype.Text
	CreatedOn  pgtype.Timestamp
	Pebbles    pgtype.Numeric
	VerifiedOn pgtype.Timestamp
	Email      pgtype.Text
}

type SignUpReq struct {
	Uname string `field:"uname" binding:"required"`
	Pword string `field:"pword" binding:"required"`
}
