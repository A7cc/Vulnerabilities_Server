# 壹 Vulnerabilities_Server_Golang

这是一个用`Golang`写的`Web`靶场，该系统是以食谱菜单管理系统为场景去编写，一种实战化形式的安全漏洞靶场，其中存在多个安全漏洞，需要我们去探索和发现。该项目旨在帮助安全研究人员和爱好者了解和掌握关于`Golang`系统的渗透测试和代码审计知识。

当前代码为靶场的后端代码，使用`Golang`语言、`Gin`框架和`mysql`数据库，前端使用`Vue`框架。项目地址：https://github.com/A7cc/Vulnerabilities_Server

项目后面的设想是以这个场景为出发点扩展出其他语言的漏洞靶场，后面会持续更新，如果您觉得对你有些许帮助，请加个⭐，您的支持将是[`Vulnerabilities_Server`](https://github.com/A7cc/Vulnerabilities_Server)对你有些许帮助，请加个⭐，您的支持将是对你有些许帮助，请加个⭐，您的支持将是[`Vulnerabilities_Server`](https://github.com/A7cc/Vulnerabilities_Server)前进路上的最好的见证！

# 贰 Vulnerability

目前有这些漏洞，如果有好的`idea`漏洞，可以提个`issues`给我，我来加：

```bash
登录处存在：用户名枚举

验证码：万能验证码

密码修改处：任意密码修改

ping处：命令执行

登录处：暴力破解

订单查询处,添加菜品：SQL注入

所有的文件上传功能点：文件上传

多处存在：越权、未授权

角色的功能：越权漏洞

数据库文件下载和删除功能：文件下载、删除和读取

获取名言金句功能：SSRF

获取数据库文件功能：目录遍历

JWT：密钥为空

日志功能：敏感信息泄露，前端信息泄露

修改价格处：负值反冲

原生模板的测试功能：模板注入

测试性功能处：ZIP的漏洞
```

> 注意：可能会有其他漏洞，在写的时候由于突然的想法加但是没提出来，如果发现的话，帮忙提个`issues `（不是交`CVE`，用这个系统交`CVE`的是`SB`）。。。
> `Web`的用户名密码：`admin`/`hJ58F_Qr96LW`

# 叁 部署

- `Golang`后端

创建一个`vul_server_go`的`mysql`数据库，然后导入`dbdata`文件夹下的`vul_server_go.sql`数据即可完成数据库部署！

如果有`golang`环境的话，直接在`Go_Server\src`目录下运行：

```bash
go run .
```

如果没有`golang`环境的话，可根据不同操作系统下载对应的可执行文件，然后运行即可。

- `Vue`前端

如果有`node`环境的话直接，运行`npm install`下载组件即可。可能会出现下面这种情况，可以忽略：

![image-20240909180126928](../README/image-20240909180126928.png)

如果没有`node`环境的话，直接用后端的`swagger`即可运行。

```bash
http://localhost:8081/swagger/index.html
```

- `Docker`部署

后端直接在`Go_Server`目录下执行：

```bash
docker-compose up -d
```

前端的话，直接在`Vue_Server`目录下执行：

```bash
docker-compose up -d
```
