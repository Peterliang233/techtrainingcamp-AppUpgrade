# techtrainingcamp-AppUpgrade
字节跳动后端训练营第五组代码提交(仅供学习参考，请勿恶意使用)

### 项目排期
[项目要求](https://docs.qq.com/doc/DTGFPVWRrRVZMWlVX)

[项目排期](https://docs.qq.com/sheet/DTGRLV3Rja0Rrb0Fi?tab=BB08J2)

### 规则相关接口redis处理流程图
+ 图1

![流程图](./source/img_1.png)
+ 图2

![流程图](./source/img_2.png)
+ 图3

![流程图](./source/img_3.png)

### api文档

<a href="https://documenter.getpostman.com/view/16170518/UVC8DmBE">在线api文档</a>

### 基本使用
+ git克隆项目 `git clone git@github.com:Peterliang233/techtrainingcamp-AppUpgrade.git`
+ 进入项目根目录 `cd techtrainingcamp-AppUpgrade`
+ 拉取依赖 `go mod download`
+ 修改数据库配置，进入config文件夹里面，对config.ini文件进行修改，主要修改mysql和redis相关配置
+ 进入数据库执行数据库脚本
+ 运行项目 `go run main.go`

### 部署
#### 后端部署
+ 直接进入deploy文件夹，执行`docker-compose -f docker-compose.yml up --build -d`就可以部署后端项目。

#### 前端部署
+ 我们进入服务器，安装nginx，在/etc/nginx/conf.d/目录下面添加一个配置文件，设置监听端口，配置文件的root配置修改为项目font目录下面的dist的目录的路径。然后保存退出。执行`nginx -s reload`即可。


### Thanks for free JetBrains license

感谢JetBrains免费提供免费使用

<a href="https://www.jetbrains.com" target="_blank"><img src="https://gitee.com/wejectchan/ginblog/raw/master/upload/jet.png" height="200" /></a>