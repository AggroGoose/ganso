package server

import (
	"database/sql"
	db "ganso-core/db/sqlc"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getCreateDeletePostArg struct {
	ID string `uri:"id" binding:"required"`
}

type ReturnPost struct {
	ID       string         `json:"id"`
	Slug     sql.NullString `json:"slug"`
	AudioUrl sql.NullString `json:"audio_url"`
	CommentCount int64 `json:"comment_count"`
	LikeCount int64 `json:"like_count"`
}
func (server *Server) getOrCreatePost(ctx *gin.Context) {
	var req getCreateDeletePostArg
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req.ID

	noPost := false
	post, err := server.store.GetPost(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			noPost = true
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	likeCount := int64(0)
	commentCount := int64(0)

	if noPost {
		post, err = server.store.CreatePost(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	} else {
		likes, err := server.store.PostLikeCount(ctx, arg)
		if err == nil {likeCount = likes}
		comments, err := server.store.CommentCount(ctx, arg)
		if err == nil {commentCount = comments}
	}

	returnPost := ReturnPost{
		ID: post.ID,
		Slug: post.Slug,
		AudioUrl: post.AudioUrl,
		LikeCount: likeCount,
		CommentCount: commentCount,
	}

	ctx.JSON(http.StatusOK, returnPost)
}

func (server *Server) deletePost(ctx *gin.Context) {
	var req getCreateDeletePostArg
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req.ID

	err := server.store.DeletePost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, "Post Deleted Successfully")
}

type updatePostAudioArg struct {
	ID       string         `json:"id" binding:"required"`
	AudioUrl string `json:"audio_url" binding:"required"`
}

func (server *Server) updatePostAudio(ctx *gin.Context) {
	var req updatePostAudioArg
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdatePostAudioParams{
		ID: req.ID,
		AudioUrl: sql.NullString{String: req.AudioUrl, Valid: true},
	}

	post, err := server.store.UpdatePostAudio(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, post)
}

type likeSaveArgs struct {
	UserID string `json:"user_id"`
	PostID string `json:"post_id"`
}

func (server *Server) likeSaveState(ctx *gin.Context) {
	var req likeSaveArgs
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	state := map[string]bool{
		"isLiked":false,
		"isSaved":false,
	}

	arg := db.GetPostLikeParams{
		UserID: req.UserID,
		PostID: req.PostID,
	}

	log.Println("LikeStateArgs", arg)

	liked, err := server.store.GetPostLike(ctx, arg)
	if err != nil {
		log.Println("GetLikeError", err)
	}

	log.Println("Liked:", liked)

	arg2 := db.GetPostSaveParams{
		UserID: req.UserID,
		PostID: req.PostID,
	}

	saved, err := server.store.GetPostSave(ctx, arg2)
	if err != nil {
		log.Println("GetSaveError", err)
	} 

	if len(liked) > 0 {
		state["isLiked"] = true
	}

	if len(saved) > 0 {
		state["isSaved"] = true
	}

	ctx.JSON(http.StatusOK, state)
}

func (server *Server) likePost(ctx *gin.Context) {
	var req likeSaveArgs
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.LikePostParams{
		UserID: req.UserID,
		PostID: req.PostID,
	}

    log.Println("User:",req.UserID)

	like, err := server.store.LikePost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, like)
}

func (server *Server) unlikePost(ctx *gin.Context) {
	var req likeSaveArgs
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.RemoveLikePostParams{
		UserID: req.UserID,
		PostID: req.PostID,
	}

	err := server.store.RemoveLikePost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, "Post like removed")
}

func (server *Server) savePost(ctx *gin.Context) {
	var req likeSaveArgs
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.SavePostParams{
		UserID: req.UserID,
		PostID: req.PostID,
	}

	save, err := server.store.SavePost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, save)
}

func (server *Server) removeSavePost(ctx *gin.Context) {
	var req likeSaveArgs
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.RemoveSavePostParams{
		UserID: req.UserID,
		PostID: req.PostID,
	}

	err := server.store.RemoveSavePost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, "Post save removed")
}

type userSavedArgs struct{
	UserID string `json:"user_id" binding:"required"`
	Limit  int32  `json:"limit" binding:"required,max=10"`
	PageNum int32  `json:"page_num" binding:"required,min=1"`
}

func (server *Server) getUserSavedPosts(ctx *gin.Context) {
	var req userSavedArgs
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	
	arg := db.GetUserSavesParams {
		UserID: req.UserID,
		Limit: req.Limit,
		Offset: (req.PageNum - 1) * req.Limit,
	}

	saves, err := server.store.GetUserSaves(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, saves)
}
