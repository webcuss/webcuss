package types

import "time"

type Avatar struct {
	Id        string
	Uname     string
	Pword     string
	CreatedOn time.Time
}

type SignUpReq struct {
	Uname string `field:"uname" binding:"required"`
	Pword string `field:"pword" binding:"required"`
}
