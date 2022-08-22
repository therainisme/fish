package receiver

import (
	"encoding/json"
	"fish/function/petpet"
	"fish/model"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

func Listen(addr string) {

	r := gin.New()
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/"},
	}))
	r.Use(gin.Recovery())

	r.POST("/", dispatch)
	r.GET("/avatar", petpet.GetAvatar)
	r.Run(addr)
}

func dispatch(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Printf("read body err: %s", err.Error())
	}

	var postEvent model.PostEvent
	_ = json.Unmarshal(body, &postEvent)

	switch postEvent.PostType {
	case "message":
		petpet.Petpet(postEvent)
	default:
		// todo default
	}
}
