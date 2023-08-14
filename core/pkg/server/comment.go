package server

import (
	"database/sql"
	db "ganso-core/db/sqlc"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createCommentRequest struct {
	UserID  string `json:"user_id" binding:"required"`
	PostID  string `json:"post_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type commentResponse struct {
	ID       int64     `json:"id"`
	UserID   string    `json:"user_id"`
	PostID   string    `json:"post_id"`
	Edited   bool      `json:"edited"`
	DateTime time.Time `json:"date_time"`
	Content  string    `json:"content"`
}

type createReplyRequest struct {
	UserID  string `json:"user_id" binding:"required"`
	CommentID  int64 `json:"comment_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type replyResponse struct {
	ID        int64     `json:"id"`
	UserID    string    `json:"user_id"`
	CommentID int64    `json:"comment_id"`
	Edited    bool      `json:"edited"`
	DateTime  time.Time `json:"date_time"`
	Content   string    `json:"content"`
}

func (server *Server) createComment(ctx *gin.Context) {
	var req createCommentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCommentParams{
		UserID: req.UserID,
		PostID: req.PostID,
		Content: req.Content,
	}

	comment, err := server.store.CreateComment(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code.Name() == "unique_violation" {
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := commentResponse(comment)
	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) createReply(ctx *gin.Context) {
	var req createReplyRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateReplyParams{
		UserID: req.UserID,
		CommentID: req.CommentID,
		Content: req.Content,
	}

	reply, err := server.store.CreateReply(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code.Name() == "unique_violation" {
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := replyResponse(reply)
	ctx.JSON(http.StatusOK, rsp)
}

type commentsForPost struct {
	PostID string `json:"post_id" binding:"required"`
	Limit  int32  `json:"limit" binding:"required,max=10"`
	PageNum int32  `json:"page_num" binding:"required,min=1"`
}

type ReturnComment struct {
	ID       int64          `json:"id"`
	Username sql.NullString `json:"username"`
	Image    sql.NullString `json:"image"`
	Content  string         `json:"content"`
	DateTime time.Time      `json:"date_time"`
	CountReply int	`json:"count_reply"`
	Replies []db.GetRepliesForCommentRow `json:"replies"`
}

func (server *Server) getCommentsForPost(ctx *gin.Context) {
var req commentsForPost
if err := ctx.ShouldBindJSON(&req); err != nil {
	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	return
}


arg := db.GetCommentsForPostParams {
	PostID: req.PostID,
	Limit: req.Limit,
	Offset: (req.PageNum - 1) * req.Limit,
}

comments, err := server.store.GetCommentsForPost(ctx, arg)
if err != nil {
	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	return
}

resComments := make([]ReturnComment, len(comments))

for i:=0;i<len(comments);i++ {
	c := comments[i]

	arg := db.GetRepliesForCommentParams {
		CommentID: c.ID,
		Limit: 3,
		Offset: 0,
	}

	commentReplies := make([]db.GetRepliesForCommentRow, 0)
	totalReplies := 0

	replies, _ := server.store.GetRepliesForComment(ctx, arg)
	if len(replies) > 0 {
		commentReplies = append(commentReplies, replies...) 
		totalReplies = int(replies[0].FullCount)
	}

	fullComment := ReturnComment {
		ID: c.ID,
		Username: c.Username,
		Image: c.Image,
		Content: c.Content,
		DateTime: c.DateTime,
		CountReply: totalReplies,
		Replies: commentReplies,
	}

	resComments[i] = fullComment
}

ctx.JSON(http.StatusOK, resComments)
}

type repliesForPost struct {
	CommentID int64 `json:"comment_id" binding:"required"`
	Limit  int32  `json:"limit" binding:"required"`
	PageNum int32  `json:"page_num" binding:"required,min=1"`
}
func (server *Server) getRepliesForComment(ctx *gin.Context) {
	var req repliesForPost
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetRepliesForCommentParams {
		CommentID: req.CommentID,
		Limit: req.Limit,
		Offset: (req.PageNum - 1) * req.Limit,
	}

	replies, err := server.store.GetRepliesForComment(ctx, arg)
if err != nil {
	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	return
}

ctx.JSON(http.StatusOK, replies)
}

type updateResponseArgs struct {
	ID      int64  `json:"id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func (server *Server) updateComment(ctx *gin.Context) {
	var req updateResponseArgs
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCommentParams {
		ID: req.ID,
		Content: req.Content,
	}

	update, err := server.store.UpdateComment(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, update)
}

func (server *Server) updateReply(ctx *gin.Context) {
	var req updateResponseArgs
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateReplyParams {
		ID: req.ID,
		Content: req.Content,
	}

	update, err := server.store.UpdateReply(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, update)
}

type deleteSingleArg struct {
	ID      int64  `json:"id" binding:"required"`
}

func (server *Server) deleteComment(ctx *gin.Context) {
	var req deleteSingleArg
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req.ID

	err := server.store.DeleteComment(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, "Comment Deleted")
}

func (server *Server) deleteReply(ctx *gin.Context) {
	var req deleteSingleArg
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req.ID

	err := server.store.DeleteReply(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, "Reply Deleted")
}