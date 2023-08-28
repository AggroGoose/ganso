package server

import (
	"database/sql"
	"log"
	"net/http"

	db "ganso-core/db/sqlc"

	"github.com/gin-gonic/gin"
)

type userIDArg struct {
	ID string `uri:"id" binding:"required"`
}

func (server *Server) getOrCreateUser(ctx *gin.Context) {
	var req userIDArg
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req.ID

	noUser := false
	user, err := server.store.GetUser(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			noUser = true
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	if noUser {
		user, err = server.store.CreateUser(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}
	ctx.JSON(http.StatusOK, user)
}

func (server *Server) deleteUser(ctx *gin.Context) {
	var req userIDArg
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req.ID

	err := server.store.DeleteUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, "User Deleted Successfully")
}

type userIntakeParams struct {
	ID       string         `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Image    string `json:"image" binding:"required"`
}

func (server *Server) userCompleteIntake(ctx *gin.Context) {
	var req userIntakeParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateUserIntakeParams{
		ID: req.ID,
		Username: sql.NullString{String: req.Username, Valid: true},
		Image: sql.NullString{String: req.Image, Valid: true},
	}

	user, err := server.store.UpdateUserIntake(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, user)
}

type updateUsernameArgs struct {
	ID       string         `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type updateUserImageArgs struct {
	ID    string         `json:"id" binding:"required"`
	Image string `json:"image" binding:"required"`
}

func (server *Server) updateUsername(ctx *gin.Context) {
	var req updateUsernameArgs
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateUserNameParams{
		ID: req.ID,
		Username: sql.NullString{String: req.Username, Valid: true},
	}

	user, err := server.store.UpdateUserName(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, user)
}

func (server *Server) updateUserImage(ctx *gin.Context) {
	var req updateUserImageArgs
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateUserImageParams{
		ID: req.ID,
		Image: sql.NullString{String: req.Image, Valid: true},
	}

	user, err := server.store.UpdateUserImage(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, user)
}

type usernameArg struct {
	Username string `uri:"username" binding:"required"`
}

func (server *Server) checkUsername(ctx *gin.Context) {
	var req usernameArg
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	log.Println("Username is:", req.Username)

	arg := sql.NullString{
		String: req.Username,
		Valid: true,
	}
	isUsed := true

	_, err := server.store.CheckUsername(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			isUsed = false
		} else {
			log.Println("Undefined error checking username:", err)
		}
	}
	ctx.JSON(http.StatusOK, map[string]bool{ "isUsed": isUsed })
}