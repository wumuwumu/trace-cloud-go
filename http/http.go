package http

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/web"
	"net/http"
)

func Init(service web.Service)  {
	engine := gin.Default()
	engine.HandleMethodNotAllowed = true
	engine.NoMethod(func(context *gin.Context) {
		context.JSON(http.StatusMethodNotAllowed,gin.H{
			"result":false,
			"error":"404",
		})
	})
	engine.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound,gin.H{"result":false,"error":"404"})
	})
	router(engine)
	service.Handle("/",engine)
}

func router(r *gin.Engine){
	r.Group("/")
	{
		r.GET("/trace/:traceId",getTrace)
		api :=r.Group("/api")
		{
			api.POST("/trace",createTrace)
		}

	}

}
