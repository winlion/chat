# 郑重声明

# 近段时间发现很多人将本系统用作生产环境，在此本人在此郑重声明：本系统只能用作学习使用，不得用作生产环境。
# 如因将本系统用做生产环境导致出现的一系列的损失，本人不承担任何责任！



# 先上效果
![界面效果图](https://www.techidea8.com/app/chat/3.jpg)
# 安装方法
本系统升级到golang1.12,请开启如下支持
```
#开启go mod支持
export GO111MODULE=on
#使用代理
export GOPROXY=https://goproxy.io

```
## 1.下载项目
```bash
git clone https://github.com/winlion/chat.git
```


## 2.项目配置，非常重要
### 2.1 ide配置

+ goland IDE

ADD Configuration->左上角+->go build类型,右侧填写
```
files :{你的项目路径}/hellox.x/main.go
workdir:{你的项目路径}/hellox.x/
```
+ vscode
安装golang插件即可,无需配置,推荐使用

### 2.2 数据库配置
修改service/init.go 中数据库配置文件
```cgo
const (
	driveName = "mysql"  //数据库类型,不要动
	dsName    = "root:root@(127.0.0.1:3306)/tech-chat?charset=utf8"  //tech-chat是数据库名称,请先创建
	showSQL   = true  //是否显示sql语句
	maxCon    = 10  //最大连接数
	NONERROR  = "noerror" //一个字符串标记常量
)
```
为你自己的数据库以及密码,格式如下
```
用户名:密码@(ip:port)/数据库名称?charset=utf8
```

### 2.3 配置子网掩码
修改ctrl/chat.go  179行左右
```cgo
func udpsendproc() {
	log.Println("start udpsendproc")
	//todo 使用udp协议拨号
	con, err := net.DialUDP("udp", nil,
		&net.UDPAddr{
			IP:   net.IPv4(192, 168, 0, 255),
			Port: 3000,
	})
    //....
}

```
其中`IP:net.IPv4(192, 168, 0, 255)`, 改为你当前应用所在服务器的子网掩码,
举个简单一点的例子,比如当前应用所安装环境是`192.168.2.106`，则需要修改参数为`net.IPv4(192, 168, 2, 255)`
`Port: 3000`为通信端口。本系统依赖于UPD进行分布式部署。因此需要在防火墙内开放该端口。

### 2.4 分布式部署
本系统支持分布式部署,要求是将当前应用部署在同一个网段中。代码修改同2.3

### 2.5 页面入口地址
```
http://127.0.0.1:8080/user/login.shtml
```

## 3.依赖包安装

使用go mod 自动处理安装包

## 4. 操作说明
关注如下公众号找到`im系统10万并发` 认真阅读
![界面效果图](https://www.techidea8.com/betaidea.png)
![界面效果图](https://www.techidea8.com/techidea8-2.jpg)
