package handler

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ovasquezbrito/base-app/pkg/service"
	"github.com/ovasquezbrito/base-app/token"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/ovasquezbrito/base-app/docs"
)

type Handler struct {
	services   *service.Service
	tokenMaker token.Maker
}

func NewHandler(services *service.Service, tokenMaker token.Maker) *Handler {
	return &Handler{services: services, tokenMaker: tokenMaker}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization, Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1)),
	)

	auth := router.Group("/user")
	{
		auth.POST("/login", h.login)
		auth.POST("/register", h.register)
		auth.POST("/verify-token", h.verifyToken)
	}

	//api := router.Group("/api", h.userIdentity)
	api := router.Group("/api", h.authMiddleware(h.tokenMaker))
	{
		roles := api.Group("/roles")
		{
			roles.GET("/", h.getAllRoles)
			roles.GET("/:id/show", h.getRoleById)
		}

		menus := api.Group("/menus")
		{
			menus.GET("/:id/show", h.getAllMenuRoleUserById)
		}

	}

	return router
}

func responseError(err error) gin.H {
	return gin.H{"error": err.Error()}
}
