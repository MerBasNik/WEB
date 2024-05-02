package handler

import (
	"fmt"

	"github.com/MerBasNik/rndmCoffee/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/MerBasNik/rndmCoffee/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(wsHandler *service.HandlerWS) *gin.Engine {
	router := gin.New()

	router.Use(CORSMiddleware())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/forgot-password", h.forgotPassword)
		auth.PUT("/reset-password/:token", h.resetPassword)
	}

	api := router.Group("/api", h.userIdentity)
	{
		profile := api.Group("/profile")
		{
			profile.POST("/create_profile", h.createProfile)
			profile.PUT("/edit_profile/:prof_id", h.editProfile)
			profile.POST("/init_all_hobby", h.initAllHobby)
			profile.GET("/get_profile/:prof_id", h.getProfile)
			profile.PUT("/upload_avatar", h.uploadAvatar)

			hobby := profile.Group(":prof_id/hobby")
			{
				hobby.POST("/create_hobby", h.createHobby)
				hobby.GET("/get_hobby", h.getAllHobby)
				hobby.DELETE("/delete_hobby/:hobby_id", h.deleteHobby)
			}
		}

		chats := api.Group("/chats")
		{
			chats.POST("/create_chat", h.createList)
			chats.GET("/get_all_chats", h.getAllLists)
			chats.POST("/find_chats_users", h.findUsersByTime)
			chats.POST("/find_chats_users_by_hobby", h.findUsersByHobby)
			chats.GET("/get_chat/:chat_id", h.getListById)
			chats.PUT("/update_chat/:chat_id", h.updateList)
			chats.PUT("/rename_chat/:chat_id", h.renameChat)
			chats.DELETE("/delete_chat/:chat_id", h.deleteList)
			chats.DELETE("/delete_find_users", h.deleteFindUsers)

			items := chats.Group(":chat_id/items")
			{
				items.POST("/create_item", h.createItem)
				items.GET("/get_all_items", h.getAllItems)
				items.GET("/get_users", h.getUsers)
				items.GET("/get_item/:item_id", h.getItemById)
				items.PUT("/update_item/:item_id", h.updateItem)
				items.DELETE("/delete_item/:item_id", h.deleteItem)
			}
		}

		webSocketApi := api.Group("/ws")
		{
			webSocketApi.POST("/createRoom", wsHandler.CreateRoom)
			webSocketApi.GET("/joinRoom/:roomId", func(c *gin.Context) {
				//wsHandler.JoinRoom()
			})
		}
	}

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		fmt.Println(c.Request.Method)

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
