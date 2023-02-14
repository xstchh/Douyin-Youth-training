# Douyin-Youth-training - 142857Group

> ## Constructure
> ### basic (基本完善)
>
> 基础功能
>
> ### interaction (初具雏形)
> 互动方向
>
> ### socialize (开发中...)
> 社交方向

---



## update informations - 2.12
此次push旨在解决目前项目结构中的问题

+ 一部分代码采用`gin`框架，而另一部分采用`Hertz`，两者并不统一，现在统一采用`Hertz`作为框架
+ 移动`common.go`导致的问题：在移动`common.go`后并且重新定义了`common`包，这个举动让原本代码中的部分`section`无法运行，此次在其中添加了`import`代码前置，以便流畅运行项目。



## update informations - 2.14

优化项目结构，直接将`common.go`移动到`controller`下，方便调用，随后补全了`common.go`；

添加`public`包，用于存储静态资源；

添加`service`包，用于处理后台逻辑；

添加`comment.go` `demodata.go` `message.go` `relation.go`用于完善业务逻辑；

补充`README.md`文件的代码模块简介部分和成员介绍部分；



---

## Code Overview：







### comment.go

对于`func CommentAction`

**更具体地**： 

CommentAction 函数接收一个 gin.Context 参数，

并从中解析出请求中的 token 和 action_type 两个参数。 

然后判断这个 token 是否在 usersLoginInfo 这个 map 中存在，

如果存在，再根据 action_type 的值来决定返回不同的响应。

如果 token 不存在或者 action_type 的值不是 "1"，

则返回一个状态码为0的响应。







### favorite.go

这个程序定义了两个处理视频收藏夹的函数

对于`func FavoriteAction`

**更具体地**：

第一个函数 FavoriteAction 使用*gin.Context作为参数，

它是一个指向包含当前HTTP请求及其上下文信息的结构体的指针。

该函数从请求中提取一个 token 查询参数，

并检查它是否存在于 usersLoginInfo 映射中。 

如果是，该函数将发送一个状态码为0的JSON响应，表示成功。 

如果 token 不存在，

该函数将发送一个带有状态代码1和状态消息**用户不存在**的JSON响应。







### feed.go

这段代码提供了一个Feed接口，

返回一个固定的视频列表给每个请求。









### message.go

这段代码实现了用户之间的聊天功能，

包括发送消息和获取聊天记录。

对于`func MessageAction`

**更具体地**： 

MessageAction 函数是一个处理用户发送消息的接口，

它会接收一个 token、to_user_id 和 content 参数。 

如果 token 是有效的，那么就会创建一条新的消息，

并将其追加到临时消息列表中。

对于`func MessageChat`

**更具体地**： 

MessageChat 函数是一个获取与另一个用户的聊天记录的接口，

它会接收一个 token 和 to_user_id 参数。 

如果 token 是有效的，那么就会返回与该用户的聊天记录。







### publish.go

主要实现了发布视频的功能，

包括上传视频文件到服务器并存储在指定目录，

还有获取所有用户发布的视频列表。







### relation.go

模拟用户关系相关的操作，

并不是真正的实现用户关系的逻辑，

只是返回了固定的模拟数据。







### user.go

实现了用户相关的功能，包括用户注册、登录、查询用户信息等。

---

## Group member

xs

gyf

dyz

ywb

lp

ll

（个人简介待补充）

---

感谢阅读
