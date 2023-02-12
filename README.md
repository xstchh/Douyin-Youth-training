# Douyin-Youth-training

## Constructure
### basic (开发中...)
基础功能

### interaction (开发中...)
互动方向

### socialize (开发中...)
社交方向

## update informations
此次push旨在解决目前项目结构中的问题

+ 一部分代码采用gin框架，而另一部分采用Hertz，两者并不统一，现在统一采用Hertz作为框架
+ 移动common.go导致的问题：在移动common.go后并且重新定义了common包，这个举动让原本代码中的部分section无法运行，此次在其中添加了import代码前置，以便流畅运行项目。