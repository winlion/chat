package ctrl

import (
	"chat/model"
	"chat/service"
	"chat/util"
	"net/http"
)

func UserLogin(writer http.ResponseWriter,
	request *http.Request) {
	//数据库操作
	//逻辑处理
	//restapi json/xml返回
	//1.获取前端传递的参数
	//mobile,passwd
	//解析参数
	//如何获得参数
	//解析参数
	request.ParseForm()

	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")

	//模拟
	user, err := userService.Login(mobile, passwd)

	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOk(writer, user, "")
	}

}

var userService service.UserService

func UserRegister(writer http.ResponseWriter,
	request *http.Request) {

	//request.ParseForm()
	//
	//mobile := request.PostForm.Get("mobile")
	//
	//plainpwd := request.PostForm.Get("passwd")
	//nickname := fmt.Sprintf("user%06d",rand.Int31())
	//avatar :=""
	//sex := model.SEX_UNKNOW
	//有了数据绑定方法,不需要其他的啦
	var user model.User
	util.Bind(request, &user)
	user, err := userService.Register(user.Mobile, user.Passwd, user.Nickname, user.Avatar, user.Sex)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOk(writer, user, "")

	}

}

//解析一下
func FindUserById(writer http.ResponseWriter,
	request *http.Request) {
	var user model.User
	util.Bind(request, &user)
	user = userService.Find(user.Id)
	if user.Id == 0 {
		util.RespFail(writer, "该用户不存在")
	} else {
		util.RespOk(writer, user, "")

	}

}
