package receiver

import (
	"encoding/json"
	"fish/function/petpet"
	"fish/model"
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

func Listen(addr string) {
	r := gin.Default()
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
