# 壹 Vulnerabilities_Server_Python

这是一个用`Python`写的`Web`靶场，该系统是以食谱菜单管理系统为场景去编写，一种实战化形式的安全漏洞靶场，其中存在多个安全漏洞，需要我们去探索和发现。该项目旨在帮助安全研究人员和爱好者了解和掌握关于`Python`系统的渗透测试和代码审计知识。

当前代码为靶场的后端代码，使用`python3`语言、`Django`框架和`mysql`数据库，前端使用`Vue`框架。项目地址：https://github.com/A7cc/Vulnerabilities_Server

项目后面的设想是以这个场景为出发点扩展出其他语言的漏洞靶场，后面会持续更新，如果您觉得对你有些许帮助，请加个⭐，您的支持将是[`Vulnerabilities_Server`](https://github.com/A7cc/Vulnerabilities_Server)对你有些许帮助，请加个⭐，您的支持将是[`Vulnerabilities_Server`](https://github.com/A7cc/Vulnerabilities_Server)前进路上的最好的见证！

# 贰 Vulnerability

目前有这些漏洞，如果有好的`idea`漏洞，可以提个`issues`给我，我来加：

```bash
home模块：未授权访问

home的ping功能：命令执行

home获取金句处：SSRF

登录：用户名枚举

登录：万能验证码

登录：暴力破解

多处存在：越权、未授权

菜品价格：正负值反冲

food删除处：任意文件删除

视频图片上传处：任意文件上传（getshell不了）

错误处理：未做统一错误处理，导致源码泄露

订单查询功能：SQL注入

日志功能：日志信息泄露

JWT：密钥为空

密码修改：任意密码修改

数据库文件下载和删除功能：文件下载、删除和读取

返回用户信息：密码泄露

测试性功能处：ZIP的漏洞
```

> 注意：可能会有其他漏洞，在写的时候由于突然的想法加但是没提出来，如果发现的话，帮忙提个`issues `（不是交`CVE`，用这个系统交`CVE`的是`SB`）。。。
> `Web`的用户名密码：`admin`/`hJ58F_Qr96LW`

# 叁 部署

- `Python`后端

用到的技术：后端是用的`python3`、`Django`框架和`mysql`数据库。

创建一个`vul_server_py`的`mysql`数据库，然后导入`dbdata`文件夹下的`vul_server_py.sql`数据即可完成数据库部署！

接着到`Python_Server\src`目录下下载库：

```bash
pip install -r requirements.txt
```

最后运行：

```bash
python .\manage.py runserver 8081
```

- `Vue`前端

如果有`node`环境的话直接，运行`npm install`下载组件即可。可能会出现下面这种情况，可以忽略：

![image-20240909180126928](../README/image-20240909180126928.png)

如果没有`node`环境的话，直接用后端的`swagger`即可运行。

```bash
http://localhost:8081/swagger/index.html
```

- `Docker`部署

后端直接在`Python_Server`目录下执行：

```bash
docker-compose up -d
```

前端的话，直接在`Vue_Server`目录下执行：

```bash
docker-compose up -d
```
