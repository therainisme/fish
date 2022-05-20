package petpet

import (
	"fish/mirai"
	"fish/model"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
)

var server = "http://10.1.1.1:5701"
var reg = regexp2.MustCompile(`(?<=\[CQ:at,qq=)\d*(?=\]\s\u6478)`, 0)

func Petpet(event model.PostEvent) {
	log.Printf("user: %s message: %s\n", event.Sender.Nickname, event.Message)

	m, _ := reg.FindStringMatch(event.Message)
	if m == nil {
		return
	}

	qq := m.Groups()[0].Captures[0].String()
	writeQQAvatarToDisk(qq)
	log.Printf("[Touched Event] user %s touch %s\n", event.Sender.Nickname, qq)
	msg := fmt.Sprintf("[CQ:image,file=%s,subtype=0]", server+"/avatar?qq="+qq)
	mirai.SendToGroup(msg, event.GroupID)
}

func GetAvatar(ctx *gin.Context) {
	qq := ctx.Query("qq")

	f, err := os.Open(getQQGIFPath(qq))
	if err != nil {
		log.Printf("open petpet gif err: %s\n", err.Error())
		return
	}
	img, err := ioutil.ReadAll(f)
	if err != nil {
		log.Printf("read gif err: %s\n", err.Error())
		return
	}

	_, err = ctx.Writer.Write(img)
	if err != nil {
		log.Printf("return img err: %s\n", err.Error())
		return
	}
}

func writeQQAvatarToDisk(qq string) {
	jpgPath := getQQJPGPath(qq)
	gifPath := getQQGIFPath(qq)

	resp, err := http.Get(fmt.Sprintf("http://q1.qlogo.cn/g?b=qq&nk=%s&s=100", qq))
	if err != nil {
		log.Printf("get user avatar err: %s\n", err)
		return
	}
	defer resp.Body.Close()
	img, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read avatar response body err: %s\n", err)
		return
	}

	f, err := os.Create(jpgPath)
	if err != nil {
		log.Printf("open file %s err: %s\n", jpgPath, err.Error())
	}
	f.Write(img)
	f.Close()

	// call petpet
	cmd := exec.Command(
		filepath.Join(petpetExecPath, "petpet"),
		jpgPath,
		gifPath,
		"20",
	)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("create pipe err", err.Error())
	}
	defer stdout.Close() // 保证关闭输出流

	if err := cmd.Start(); err != nil { // 运行命令
		fmt.Println("petpet exec err", err.Error())
	}

	if _, err := ioutil.ReadAll(stdout); err != nil { // 读取输出结果
		fmt.Println("petpet err", err.Error())
	}

	cmd.Wait()
}
