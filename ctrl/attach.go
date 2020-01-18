package ctrl

import (
	"chat/util"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func init() {
	os.MkdirAll("./mnt", os.ModePerm)
}
func Upload(w http.ResponseWriter, r *http.Request) {
	UploadLocal(w, r)
	//UploadOss(w,r)
}

//1.存储位置 ./mnt,需要确保已经创建好
//2.url格式 /mnt/xxxx.png  需要确保网络能访问/mnt/
func UploadLocal(writer http.ResponseWriter,
	request *http.Request) {
	//todo 获得上传的源文件s
	srcfile, head, err := request.FormFile("file")
	if err != nil {
		util.RespFail(writer, err.Error())
	}

	//todo 创建一个新文件d
	suffix := ".png"
	//如果前端文件名称包含后缀 xx.xx.png
	ofilename := head.Filename
	tmp := strings.Split(ofilename, ".")
	if len(tmp) > 1 {
		suffix = "." + tmp[len(tmp)-1]
	}
	//如果前端指定filetype
	//formdata.append("filetype",".png")
	filetype := request.FormValue("filetype")
	if len(filetype) > 0 {
		suffix = filetype
	}
	//time.Now().Unix()
	filename := fmt.Sprintf("%d%04d%s",
		time.Now().Unix(), rand.Int31(),
		suffix)
	dstfile, err := os.Create("./mnt/" + filename)
	if err != nil {
		util.RespFail(writer, err.Error())
		return
	}

	//todo 将源文件内容copy到新文件
	_, err = io.Copy(dstfile, srcfile)
	if err != nil {
		util.RespFail(writer, err.Error())
		return
	}
	//todo 将新文件路径转换成url地址

	url := "/mnt/" + filename
	//todo 响应到前端
	util.RespOk(writer, url, "")
}

//即将删掉,定期更新
const (
	AccessKeyId     = "5p2RZKnrUanMuQw9"
	AccessKeySecret = "bsNmjU8Au08axedV40TRPCS5XIFAkK"
	EndPoint        = "oss-cn-shenzhen.aliyuncs.com"
	Bucket          = "winliondev"
)

//权限设置为公共读状态
//需要安装
func UploadOss(writer http.ResponseWriter,
	request *http.Request) {
	//todo 获得上传的文件
	srcfile, head, err := request.FormFile("file")
	if err != nil {
		util.RespFail(writer, err.Error())
		return
	}

	//todo 获得文件后缀.png/.mp3

	suffix := ".png"
	//如果前端文件名称包含后缀 xx.xx.png
	ofilename := head.Filename
	tmp := strings.Split(ofilename, ".")
	if len(tmp) > 1 {
		suffix = "." + tmp[len(tmp)-1]
	}
	//如果前端指定filetype
	//formdata.append("filetype",".png")
	filetype := request.FormValue("filetype")
	if len(filetype) > 0 {
		suffix = filetype
	}

	//todo 初始化ossclient
	client, err := oss.New(EndPoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		util.RespFail(writer, err.Error())
		return
	}
	//todo 获得bucket
	bucket, err := client.Bucket(Bucket)
	if err != nil {
		util.RespFail(writer, err.Error())
		return
	}
	//todo 设置文件名称
	//time.Now().Unix()
	filename := fmt.Sprintf("mnt/%d%04d%s",
		time.Now().Unix(), rand.Int31(),
		suffix)
	//todo 通过bucket上传
	err = bucket.PutObject(filename, srcfile)
	if err != nil {
		util.RespFail(writer, err.Error())
		return
	}
	//todo 获得url地址
	url := "http://" + Bucket + "." + EndPoint + "/" + filename

	//todo 响应到前端
	util.RespOk(writer, url, "")
}
