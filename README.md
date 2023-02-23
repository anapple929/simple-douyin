# simple-douyin

**简介：** gorm+gin+go-micro+mysql+redis 极简版抖音，实现了基础功能，互动功能和社交功能。侧重于互动功能。

**队号：** 1122233

**开发接口文档：** https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof#6QCRJV

**项目答辩文档：** https://dzs47lqpfu.feishu.cn/docx/Q48XdJaBxosEwSxCtWNc2pCLnQc

**线上演示视频：** https://www.bilibili.com/video/BV16D4y1G7Ca/?vd_source=140298c09a850b3c5413933606102807

**项目地址：** 43.138.51.56:4000

---
### 运行
1. `git clone https://github.com/anapple929/simple-douyin`
2. 进入每个微服务`go mod tidy`，
3. 将 `etcd` 运行起来
4. 运行每个微服务的`main.go`
5. `postman`或前端项目测试，网关端口：`4000`

---
### 微服务

|   微服务    | 对网关暴露的接口                               | 为其他微服务提供的接口 |
|:--------:|:---------------------------------------|:------------|
|   **user**   | 注册<br/>登录<br/>根据用户id查询用户信息             | 通过批量用户id查询批量user<br/>修改关注数<br/>修改粉丝数<br/>修改发布数<br/>修改获赞数<br/>修改点赞数            |
| **publish**  | 视频发布<br/>  视频列表                        | 通过批量视频id查询视频信息<br/>修改点赞数<br/>修改评论数            |
|   **feed**   | 视频流                                    |             |
| **favorite** | 点赞<br/>取消点赞<br/>点赞列表                   |  根据 用户id 和 视频id 判断是否有点赞关系<br/>根据 用户id 视频id 点赞关系 构成的结构体，批量查询点赞关系<br/>           |
| **comment**  | 评论<br/>删除评论<br/>评论列表                   |             |
| **relation** | 关注<br/>取消关注<br/>关注列表<br/>粉丝列表<br/>好友列表 |  根据 用户id 和 用户id 判断是否有关注关系<br/>根据 用户id1 用户id2 关注关系 构成的结构体，批量查询关注关系           |
| **message**  | 聊天记录<br/>发消息                           | 根据用户id1和用户id2，查找二人聊天最新一条消息实体            |


---
### mysql

|    表     |   拥有权限的微服务   |
|:--------:|:------------:|
|   user   |     user     |
| comment  |   comment    |
| favorite |   favorite   |
| relation |   relation   |
|  video   | publish,feed |
| message  |   message    |

---
### redis

热点数据

* db0: userid -> user
* db1: videoid -> video
* db2: userid+videoid -> bool  是否点赞 

---
### 实现的接口

基础接口，互动接口，社交接口

* /douyin/user/login/
* /douyin/user/register/
* /douyin/user/


* /douyin/publish/action/
* /douyin/publish/list/
* /douyin/feed/


* /douyin/favorite/action/
* /douyin/favorite/list/


* /douyin/comment/action/
* /douyin/comment/list/


* /douyin/relation/action/
* /douyin/relation/follow/list/
* /douyin/relation/follower/list/
* /douyin/relation/friend/list/


* /douyin/message/chat/
* /douyin/message/action/

