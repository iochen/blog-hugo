---
title: 从订阅链接生成V2Ray配置 - V2Gen的使用
typora-root-url: ../
date: 2020-02-02 18:26:20
tags:
- V2Ray
- V2Gen
- Vmess
categories:
- Tutorials
description: V2Gen的使用说明
comments: true
---



# V2Gen

从订阅以及 `vmess://XXXXXXXXXXXX` 链接中生成 V2Ray json 文件



## 项目地址

[***github.com/iochen/v2gen***](https://github.com/iochen/v2gen/)



## 使用预览

```bash
richard@openSUSE~> v2gen -u "https://example.ltd/xxxxxx"
[0]     中继香港G1 Media (HK)(1)        [21.00ms (cu2.example.ltd)]
[1]     中继香港G2 Media (HK)(1)        [21.02ms (cu1.example.ltd)]
[2]     中继香港C1 Media (HK)(1)        [34.36ms (hkc1.example.ltd)]
[3]     中继香港C2 Media (HK)(1)        [39.67ms (hkc2.example.ltd)]
[4]     中继香港C3 Media (HK)(1)        [43.59ms (hkc3.example.ltd)]
[5]     中继香港C4 Media (HK)(1)        [35.71ms (hkc4.example.ltd)]
...(节点过多，不一一展示)
[47]    美国GIA 3 Media (US)(0.7)       [140.48ms (us3.example.ltd)]
[48]    香港 8 (HK)(1)                  [70.38ms (hk9.example.ltd)]
[49]    香港 9 (HK)(1)                  [72.74ms (hk10.example.ltd)]
[50]    香港负载均衡 1 Test (HK)(1)      [16.41ms (1.1.1.1)]
=====================
Please Select:
```

你可以到[***这里***](/donate/)使用博主的AFF



## 如何安装

编译它（请确保您的`$GOPATH/bin/`已添加至`$PATH`中）

```sh
go get -u github.com/iochen/v2gen/cmd
```

或者到 [***GitHub Release***](https://github.com/iochen/v2gen/releases) 下载



## 常用情景

```bash
v2gen -u "https://example.ltd/xxxxxx" -o <V2Ray配置文件路径> # -o 参数不加，为默认路径
v2gen -u "https://example.ltd/xxxxxx" -r  # 随机节点
v2gen -u "https://example.ltd/xxxxxx" -n <节点编号> -o "-" | jq .  # 使用管道进行进阶操作（此处为举例）
```



## 参数

```Usage
Usage of v2gen:
  -c string
        v2gen配置文件路径 (default "/etc/v2ray/v2gen.ini")
  -init
        初始化v2gen配置文件
  -n int
        节点引索 (default -1)
  -np
        不ping
  -o string
        输出路径 (default "/etc/v2ray/config.json")
  -r    随机节点引索
  -tpl string
        V2Ray模板路径
  -u string
        订阅链接
  -v    展示版本
  -vmess string
        vmess 链接
```

## V2Gen 用户配置

你可以使用 `v2gen --init` 来生成一个新的

```yaml
# V2Ray 日志等级
# ( debug | info | warning | error | none )
loglevel: warning

# Socks 端口
socksPort: 1080

# Http 端口
httpPort: 1081

# 是否允许UDP流量
# ( true | false )
udp: true

# 安全
# ( aes-128-gcm | aes-256-gcm | chacha20-poly1305 | auto | none )
security: aes-256-gcm

# 是否开启 mux
# ( true | false )
mux: true

# Mux 并发数
concurrency: 8

# DNS 服务器
dns1: 9.9.9.9
dns2: 1.1.1.1

# 中国IP与网站是否直连
# ( true | false )
china: true

```

下面的配置可能不会在所有节点上生效

```yaml
# 是否允许不安全连接 ( true | false )
allowInsecure: false

# KCP mtu 值
mtu: 1350

# KCP tti 值
tti: 20

# KCP 最大上行速度
# 单位: MB/s
up: 5

# KCP 最大下行速度
# 单位: MB/s
down: 20

# 是否开启 UDP 拥堵控制 ( true | false )
congestion: false

# 读缓冲区大小
# 单位: MB
readBufferSize: 1

# 写缓冲区大小
# 单位: MB
writeBufferSize: 1
```

## 版本

*V1.1.2*



## TODO
* *V1.2.1*:    **VmessPing**




## 协议

MIT LICENSE

## 注意

不支持第一版本格式