package configs

import (
	"golang-example/controllers"
	"golang-example/docs"
	"golang-example/repositories"
	"golang-example/services"
	"fmt"
	"os"
	"time"

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
	vodReposi := repositories.NewVODRepository(db)

	vodService := services.NewVODService(vodReposi)

	vodController := controllers.NewVODController(vodService)
	// init logger
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

	vodRouter := router.Group("/api/vod")
	{
		vodRouter.POST("/", vodController.AddRecordFile)
		vodRouter.GET("/", vodController.GetAllRecordFile)
		vodRouter.PUT("/", vodController.UpdateRecordFile)
		vodRouter.GET("/search", vodController.GetAllVodByFilter)
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.router = router
	return server
}

func (server *Server) Start() error {
	address := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	return server.router.Run(address)
}
