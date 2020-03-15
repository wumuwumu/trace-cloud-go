package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/wumuwumu/trace-cloud/storage"
	"net/http"
)



func createTrace(c *gin.Context){
	var trace storage.Trace
	err := c.BindJSON(&trace)
	if err != nil{
		logrus.WithError(err).Error("参数解析错误")
		c.JSON(http.StatusOK,gin.H{"code": 1, "error": "参数错误"})
		return
	}
	id,err:=storage.CreateTrace(trace)
	c.JSON(http.StatusOK,gin.H{"code": 1, "data": id})
}

func getTrace(c *gin.Context){
	traceId := c.Param("traceId")
	trace,err := storage.GetTrace(traceId)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{"code": 1, "error": "获取出错"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":1,"data":trace})
}
