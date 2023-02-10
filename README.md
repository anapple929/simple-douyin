# simple-douyin
gorm+gin+go-micro+mysql 极简版抖音

---
* 注册中心:etcd
* 网关 - api-gateway 4000
* 用户微服务 - user 8082
* token 8085
* publish 8083
* feed 8084
* favorite 8086
* comment 8087
---
### mysql 
43.138.51.56:3306  
数据库名:simple-douyin 
用户名:root 
密码:jhr292023 

* user 用户表
* comment 评论表
* favorite 用户对视频点赞表
* relation 用户和用户之间关注表
* user 用户表
* video 视频表

---
### 运行
1. `git clone XXX`
2. 在 goland 中打开，进入每个微服务`go mod tidy`，
3. 将 etcd 运行起来
4. 运行每个微服务的main.go
5. postman或前端项目测试接口，入口是统一的网关：localhost:4000/douyin/xxx
---
### 实现的接口
* /douyin/user/login/
* /douyin/user/register/
* /douyin/user/
* /douyin/publish/action/
* /douyin/publish/list/
* /douyin/feed/
* /douyin/favorite/action/
* /douyin/favorite/list/
---
### 注意
测试视频上传功能时，电脑要下载ffmpeg，并配置到环境变量中。