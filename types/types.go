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
	Pword string `field:"pword" binding:"required,min=6"`
}

type SignInReq struct {
	Uname string `field:"uname" binding:"required"`
	Pword string `field:"pword" binding:"required"`
}

type PostTopicReq struct {
	Url     string `field:"url" binding:"required,min=3"`
	Title   string `field:"title" binding:"required"`
	Comment string `field:"comment"`
}

type TopicCommentUri struct {
	TopicId string `uri:"topicId" binding:"required"`
}

type PostCommentReq struct {
	Comment string `field:"comment" binding:"required"`
}

type CommentReplyUri struct {
	CommentId string `uri:"commentId" binding:"required"`
}

type PostReplyReq struct {
	Comment string `field:"comment" binding:"required"`
}
