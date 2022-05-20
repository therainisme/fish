package function

import (
	"fish/mirai"
	"fish/model"
	"fmt"
	"github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

var reg = regexp2.MustCompile(`(?<=\[CQ:at,qq=)\d*(?=\]\s\u6478)`, 0)

func Petpet(event model.PostEvent) {
	log.Printf("user: %s message: %s\n", event.Sender.Nickname, event.Message)

	m, _ := reg.FindStringMatch(event.Message)
	if m == nil {
		return
	}

	qq := m.Groups()[0].Captures[0].String()
	writeQQAvatarToStore(qq, "./")
	log.Printf("user %s touch %s\n", event.Sender.Nickname, qq)
	msg := fmt.Sprintf("[CQ:image,file=%s,subtype=0]", "http://10.1.1.77:5701/avatar?qq="+qq)
	mirai.SendToGroup(msg, event.GroupID)
}

func GetAvatar(ctx *gin.Context) {
	qq := ctx.Query("qq")
	img, _ := avatarMap.Load(qq)

	_, err := ctx.Writer.Write(img.([]byte))
	if err != nil {
		log.Printf("return img err: %s\n", err.Error())
		return
	}
}

func writeQQAvatarToStore(qq string, path string) {
	//path := filepath.Join(path, qq+".png")
	resp, err := http.Get(fmt.Sprintf("http://q1.qlogo.cn/g?b=qq&nk=%s&s=100", qq))
	if err != nil {
		log.Printf("get user avatar err: %s\n", err)
		return
	}
	img, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read avatar response body err: %s\n", err)
		return
	}

	// todo call petpet .exe
	avatarMap.Store(qq, img)
}
