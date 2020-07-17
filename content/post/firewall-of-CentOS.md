---
title: CentOS7防火墙的使用
date: 2018-06-02 18:28:52
tags:
- CentOS
- Linux
- firewall
categories:
- Linux
- CentOS
---
记录**CentOS7**下**firewall(firewalld)**的用法，包括*开启*、*关闭*、*查看*等操作。
<!--more-->

# 命令
直接上命令  
## 开启
```bash
$ sudo firewall-cmd --zone=public  --add-port=80/tcp --permanent
```
[*注：执行此条命令后不会立即生效，请参阅此处*](#重启)  
这里是开启**TCP**的**80**端口，下方是详解  
```
--zone=public 操作的域  
--add-port=80/tcp 要开启的端口，格式为`端口/协议`  
--permanent 指的是永久生效，去除即临时生效  
```

## 关闭
```bash
$ sudo firewall-cmd --remove-port=80/tcp --permanent
```
[*注：执行此条命令后不会立即生效，请参阅此处*](#重启)  
这里是关闭**TCP**的**80**端口，下方是详解  
```
--zone=public 操作的域  
--remove-port=80/tcp 要开启的端口，格式为`端口/协议`  
--permanent 指的是永久生效，去除即临时生效
```

## 查看
```bash
$ sudo firewall-cmd --list-ports
```

## 重启
```bash
$ sudo firewall-cmd --reload
```


# 启用or关闭**firewall(firewalld)**服务
## 关闭
```bash
$ sudo systemctl stop firewalld.service
```
## 开启
```bash
$ sudo systemctl start firewalld.service
```
## 重启
```bash
$ sudo systemctl restart firewalld.service
```
## 状态
```bash
$ sudo systemctl status firewalld.service
```
我自己的示例(开启状态):
![](https://img.iochen.com/AJ5Ssm0L.png)
