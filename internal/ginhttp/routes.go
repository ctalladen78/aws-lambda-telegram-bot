package ginhttp

import (
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ksopin/aws-lambda-telegram-bot/internal/app"
	"net/http"
)

func New() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	InitRoutes(r)
	return r
}

func Run() error {
	return New().Run(":80")
}

func InitRoutes(engine *gin.Engine) {
	engine.GET("/", info)
	engine.POST("/", post)
}

func post(c *gin.Context) {
	u := &tgbotapi.Update{}
	err := c.BindJSON(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg": err.Error(),
		})
		return
	}

	err = app.Get().Reply(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func info(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg": "Bot",
	})
}

