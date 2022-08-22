package mirai

import (
	"fish/config"
	"fmt"
	"log"
	"net/http"
)

func SendToGroup(msg string, group int) {
	str := fmt.Sprintf("%s/send_group_msg?message=%s&group_id=%d", *config.BotAddress, msg, group)
	log.Println(str)
	_, err := http.Get(str)
	if err != nil {
		log.Printf("send group message err: %s", err.Error())
	}
}
