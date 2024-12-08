# trade
the backend code for the second-hand trading website project
## 目录
- api : 用于定义接口函数,也就是controller层

- cmd : 程序启动

- conf : 用于存储配置文件

- middleware : 应用中间件

- pkg/ctl : 封装响应

- pkg/e : 封装业务状态码

- pkg/util : 工具函数(日志打印)

- repository: 仓库放置所有存储

- repository/cache: 放置redis缓存

- repository/db: 持久层MySQL仓库

- repository/db/dao: 对db进行操作的dao层

- repository/db/model: 定义所有持久层数据库表结构的model层

- routes : 路由逻辑处理

- service : 接口函数的实现

- types : 放置请求/响应信息的结构体

## 安装和启动redis
下载window用的 5.0 版本Redis-x64-5.0.14.1.zip
https://github.com/tporadowski/redis/releases

解压后进入到目录下，打开命令行启动服务：

`redis-server.exe redis.windows.conf`

## 数据库配置
将 config/conf/config.yaml中

mysql的userName和password改为自己数据库的信息

