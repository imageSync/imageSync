# 命令行工具：imageSync

### 作用：
- 拉取海外的镜像（比如k8s需要的一些镜像）到本地
- 再把本地的镜像推动到自己指定的仓库中

### 工作机制：
借用了本地`docker daemon`中配置的`registry-mirrors`的加速功能，当我们利用`registry-mirrors`列表中的加速源，加速的把一些很难下载的镜像下载到本地后；该工具再自动化的把这个镜像推送到您云端（自己的私有镜像仓库中）保存起来。 该命令的作用就是自动化处理了这个过程中的琐事，实现一键全自动化的`拉取->推送`

> 这里列举一些加速比较好的镜像仓库：
```bash
https://hub-mirror.c.163.com
https://mirror.baidubce.com
http://f1361db2.m.daocloud.io
```

<br><br>

### 使用教程：

##### 安装
(明天在写...)

<br>

##### 使用

执行如下命令，会在用户家目录中生成一个名为`.imageSync`的配置文件，然后大家取修改一下该文件中的账号密码信息即可进行后续的使用。
```bash
imageSync init
```

> 这个命令不用多次执行（每一次执行都会重新覆盖为初始状态）

`~/.imageSync`文件内容如下：

```json
{
  "username": "admin",
  "password": "123456",
  "server_address": "registry.cn-shanghai.aliyuncs.com",
  "image_tag": "registry.cn-shanghai.aliyuncs.com/tay3223/images"
}
```

- username: 您（私有镜像仓库）的账号
- password: 您（私有镜像仓库）的密码
- server_address: 您（私有镜像仓库）的登录域名
- image_tag: 当您把镜像向私有仓库推送时，需要打上的Tag前缀

<br><br>

上面这几个参数信息，与我们平时使用命令行的`docker login`命令时所用到的概念是一样的。过去我们使用`docker login`或者`docker push`时，需要用到`账号、密码、域名、tag`等等，此处也是使用的这些内容。

> 不过`imageSync`命令比较方便一些，只需要在家目录配置文件（~/.imageSync）中配置一次。 随后就可以随意使用命令行工具imageSync进行对镜像的快速操作了。

<br>



下文借用阿里云的一部分文档来演示说明一下：


```bash
（原文）将镜像推送到阿里云Registry中时：

$ docker login --username=**** registry.cn-shanghai.aliyuncs.com
$ docker tag [ImageId] registry.cn-shanghai.aliyuncs.com/tay3223/images:[镜像版本号]
$ docker push registry.cn-shanghai.aliyuncs.com/tay3223/images:[镜像版本号]
```

这里可以看到使用docker原生命令来进行镜像操作时，也是需要这四类信息，只不过我这里时把它变成落地到配置文件中了。

<br><br>

> 疫情封闭时随手写的项目，有很多地方不是很完善，如果对这个小工具感觉有意思，想一起延展性开发点小功能的，欢迎随时联系作者。



(全文完)
