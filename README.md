# UDNS

一个更新阿里云域名解析的工具。

## 应用场景

因为对于家庭环境，网络的公网 ip 往往是会变动的。假设我们的电脑上有一些需要对外提供的服务，但是如果用 ip 的话，因为 ip 会变，所以在 ip 变动之后，通过原来的 ip 就访问不了了。

因此，我们可以通过将域名解析到我们的这个公网 ip 上，然后通过定时任务来检查当前的公网 ip 是否已经变了，如果已经发生变化，则更新域名解析指向的 ip。这样一来，
外部就可以通过域名来相对稳定地访问我们电脑对外提供的服务了。

比如配置如下，假设当前电脑的公网 ip 为 `1.2.3.4`：

```yaml
domain: example.com
replace: vpn.example.com
```

则执行脚本会将自己的阿里云域名 `vpn.example.com` 解析为 `1.2.3.4`，为 A 记录。


## 安装

```shell
go install github.com/eleven26/udns@latest
```

> 通过这种方式安装，需要将 $GOPATH 加到环境变量。

## 配置

将以下内容复制，然后保存到 `/etc/udns.yml` 或者 `~/udns.yml`

```yaml
endpoint: "alidns.cn-shenzhen.aliyuncs.com"
access_key_id: ""
access_key_secret: ""
domain: example.com
replace: vpn.example.com
```

## 执行

```shell
udns
```

## 添加为定时任务

在 Linux 或者 macOS 中可以通过 `crontab` 来添加定时任务，进行定时检查更新。

如，每 10 分钟检查一次：

> `$GOPATH` 需要替换为本地 `GOPATH` 的具体路径。

```shell
*/10 * * * * $GOPATH/udns
```