package server

import (
	"net/http"
	"time"

	db "ganso-core/db/sqlc"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	  }))

	

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "What's good, dawg!?")
	})
	/*
	##################
	## Post Actions ##
	##################
	*/
	router.POST("/post/GetorCreate/:id", server.getOrCreatePost)
	router.POST("/post/LikeSaveState", server.likeSaveState)
	router.POST("/post/LikePost", server.likePost)
	router.POST("/post/SavePost", server.savePost)
	router.POST("/post/GetSavedPosts", server.getUserSavedPosts)
	router.PUT("/post/UpdateAudio", server.updatePostAudio)
	router.DELETE("/post/UnlikePost", server.unlikePost)
	router.DELETE("/post/RemoveSavePost", server.removeSavePost)
	/*
	##################
	## User Actions ##
	##################
	*/
	router.GET("/user/CheckUsername/:username", server.checkUsername)
	router.POST("/user/GetorCreate/:id", server.getOrCreateUser)
	router.PUT("/user/IntakeComplete", server.userCompleteIntake)
	router.PUT("/user/UpdateUsername", server.updateUsername)
	router.PUT("/user/UpdateUserImage", server.updateUserImage)
	router.DELETE("/user/DeleteUser/:id", server.deleteUser)
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

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}