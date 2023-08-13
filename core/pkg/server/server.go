package server

import (
	"encoding/json"
	"net/http"

	db "ganso-core/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "What's good, dawg!?")
	})

	/*
	#####################
	## Comment Actions ##
	#####################
	*/
	router.POST("/comment/GetCommentPost", server.getCommentsForPost)
	router.POST("/comment/GetReplyComment", server.getRepliesForComment)
	router.POST("/comment/CreateComment", server.createComment)
	router.POST("/comment/CreateReply", server.createReply)
	router.PUT("/comment/EditComment", server.updateComment)
	router.PUT("/comment/EditReply", server.updateReply)
	router.DELETE("/comment/DeleteComment", server.deleteComment)
	router.DELETE("/comment/DeleteReply", server.deleteReply)
	
	server.router = router
	return server
}

func writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}