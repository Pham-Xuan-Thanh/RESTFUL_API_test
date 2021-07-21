package configs

import (
	"fmt"
	"golang-example/controllers"
	"golang-example/docs"
	"golang-example/repositories"
	"golang-example/services"
	"os"
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type Server struct {
	db     *gorm.DB
	router *gin.Engine
}

func InitServer(db *gorm.DB) *Server {
	url_docs := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Egde247 docs API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = url_docs
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	server := &Server{db: db}
	router := gin.New()

	// init module

	deviceReposi := repositories.NewDeviceRepository(db)
	deviceService := services.NewDeviceService(deviceReposi)
	deviceController := controllers.NewDeviceController(deviceService)

	// init logger
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT","GET","POST", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		// MaxAge: 12 * time.Hour,
	}))
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.Use(gin.Recovery())

	deviRouter := router.Group("/api/device")
	{
		deviRouter.POST("/", deviceController.InsertDevice)
		deviRouter.GET("/", deviceController.GetAll)
		deviRouter.PUT("/", deviceController.UpdateDevice)
		deviRouter.GET("/search", deviceController.GetAllDeviceByFilter)
		deviRouter.DELETE("/",deviceController.DeleteDevice)

	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.router = router
	return server
}

func (server *Server) Start() error {
	address := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	return server.router.Run(address)
}
