package mirai

import (
	"fmt"
	"log"
	"net/http"
)

var URL = "http://10.1.1.52:5700"

func SendToGroup(msg string, group int) {
	str := fmt.Sprintf("%s/send_msg?message=%s&group_id=%d", URL, msg, group)
	log.Println(str)
	_, err := http.Get(str)
	if err != nil {
		log.Printf("send group message err: %s", err.Error())
	}
}