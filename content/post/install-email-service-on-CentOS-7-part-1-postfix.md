---
title: 基于CentOS 7,使用Postfix & Dovecot搭建邮件服务器 -Part 1
date: 2018-07-21 19:52:16
tags:
- Postfix
- Dovecot
- 邮件服务器
- 自建邮件服务器
- CentOS
- Linux
- DKIM
- SPF记录
- PTR记录
- DMARC记录
- MX记录
categories:
- Tutorials
comments: true
---

本文将详细讲解 如何在CentOS 7下，使用Postfix & Dovecot 搭建自己的邮件服务器。  
这是第一部分: Postfix 与 DNS各种记录。
<!--more-->

# 前言
在我自己搭建的时候，看了许多网络教程，但博主自己在实现过程中遇到了许多坑，比如教程本身有问题，讲解不明等，所以在此自己写一篇。  

# 实践
## **变量设定**
<font color=red>
重要!!!   
**假如你的IP为`8.8.8.8`, 域名为`example.com`, 你想要创建的邮箱为` xxx@example.com`**
</font>

## DNS记录
### 必要记录
|名称(Name)|记录类型(Type)|记录内容(Value)|注|
|:---:|:---:|:---:|:---:|
|`mail` |A   |`8.8.8.8`|/|
|`@`   |MX   |`mail.example.com`|优先级(Priority)为 `10`|
### 选填记录
变量假设: 你的邮箱为` i@example.com`  

|名称(Name)|记录类型(Type)|记录内容(Value)|注|
|:---:|:---:|:---:|:---:|
|`@` |TXT|`v=spf1 mx ~all`|[见下](#SPF记录)|
|`_dmarc`|TXT|`v=DMARC1;p=reject;rua=i@example.com`|[见下](#DMARC记录)|

#### SPF记录
**SPF记录**可确定允许哪些IP 使用你的域名 来发送电子邮件。  
SPF的设置选项可以参考：[SPF Record Syntax](http://www.openspf.org/SPFRecordSyntax)  *注:此链接未使用HTTPS*   
部分如下：  
`a`：所有该域名的**A记录**都为通过，a不指定的情况下为当前域名  
`ip4`：指定通过的IPv4地址  
`mx`：mx记录域名的A记录IP可以发邮件  
`all`：指其它IP该作何处理，`-`表示**拒绝**，`~`表示**软拒绝**，`+`表示**通过**  
这里的`v=spf1 mx ~all`，则表示允许MX记录的IP收发邮件,软拒绝其它IP发来的邮件。  
#### DMARC记录
**DMARC协议**是有效解决信头From伪造而诞生的一种新的邮件来源验证手段，为邮件发件人地址提供强大保护，并在邮件收发双方之间建立起一个数据反馈机制。  
具体信息可以看这里：[DMARC Overview](https://dmarc.org/overview/)。  
**DMARC记录**中常用的参数解释:  
`p`：用于告知收件方，当检测到某邮件存在伪造我（发件人）的情况，收件方要做出什么处理，处理方式从轻到重依次为：`none`为不作任何处理；`quarantine`为将邮件标记为垃圾邮件；`reject`为拒绝该邮件。  
`rua`：用于在收件方检测后，将一段时间的汇总报告，发送到哪个邮箱地址。  
`ruf`：用于当检测到伪造邮件时，收件方须将该伪造信息的报告发送到哪个邮箱地址。  
例如我设置的是`v=DMARC1;p=reject;rua=i@example.com`，意思是拒绝伪造邮件，并且将一段时间的汇总报告发送给`i@example.com`。

## 证书签发
这里使用**Certbot**
### 安装Certbot
```bash
$ sudo yum -y install yum-utils
$ sudo yum-config-manager --enable rhui-REGION-rhel-server-extras rhui-REGION-rhel-server-optional
$ sudo yum install certbot
```
### 签发证书
```bash
$ sudo certbot certonly -d mail.example.com
```
看提示, 使用**root用户**登录, 在提示的目录下会有`cert.pem` `chain.pem` `fullchain.pem` `privkey.pem`   
博主这些文件的所在目录是`/etc/letsencrypt/live/mail.example.com/`  
那么**证书文件**就是`/etc/letsencrypt/live/mail.example.com/cert.pem`  
**密钥文件**是`/etc/letsencrypt/live/mail.example.com/privkey.pem`  
下文请注意替换!!!

## Postfix
### 安装Postfix
```bash
$ sudo yum -y install postfix
```
### 配置Postfix
编辑`/etc/postfix/main.cf`文件
```conf
# 第75行: 取消注释并修改
myhostname = mail.example.com

# 第83行: 取消注释并修改
mydomain = example.com

# 第99行: 取消注释并修改
myorigin = $mydomain

# 第116行: 修改
inet_interfaces = all

# 第164行: 添加
mydestination = $myhostname, localhost.$mydomain, localhost, $mydomain

# 第264行: 取消注释并修改为你的内网(就填下面的一般也行)
mynetworks = 127.0.0.0/8, 10.0.0.0/24

# 第419行: 修改
home_mailbox = Mailbox/

# 第574行: 添加
smtpd_banner = $myhostname ESMTP

# 把下面的内容添加到最后
# 限制邮件大小为10M
message_size_limit = 10485760
# 限制邮箱大小为1G
mailbox_size_limit = 1073741824
# SMTP验证用内容 (为接下来的Dovecot做准备)
smtpd_sasl_type = dovecot
smtpd_sasl_path = private/auth
smtpd_sasl_auth_enable = yes
smtpd_sasl_security_options = noanonymous
smtpd_sasl_local_domain = $myhostname
smtpd_recipient_restrictions = permit_mynetworks,permit_auth_destination,permit_sasl_authenticated,reject
# TLS/SSL内容, 注意替换路径!!!
smtpd_use_tls = yes
smtpd_tls_cert_file = /etc/letsencrypt/live/mail.example.com/cert.pem
smtpd_tls_key_file = /etc/letsencrypt/live/mail.example.com/privkey.pem
smtpd_tls_session_cache_database = btree:/etc/postfix/smtpd_scache
```
编辑`/etc/postfix/master.cf`
```
# 第26~28行: 取消注释
smtps       inet   n       -       n       -       -       smtpd
  -o syslog_name=postfix/smtps
  -o smtpd_tls_wrappermode=yes
```

### 重启Postfix并开机自启
```bash
$ sudo systemctl restart postfix && sudo systemctl enable postfix
```
### firewall相应配置
```bash
$ sudo firewall-cmd --add-port={25/tcp,465/tcp} --permanent && sudo firewall-cmd --reload
```
**25/tcp**为**SMTP**端口, **465/tcp**为**SMTPS**端口  
注: 可能部分VPS未开放25端口, 请与VPS提供商联系!

## DKIM
**DomainKeys Identified Mail**的缩写，域名密钥识别邮件标准。
### 安装
#### 下载安装EPEL  
64 bit:  
```bash
$ sudo rpm -Uvh http://download.fedoraproject.org/pub/epel/6/x86_64/epel-release-6-8.noarch.rpm
```
32 bit:
```bash
$ sudo rpm -Uvh http://download.fedoraproject.org/pub/epel/6/i386/epel-release-6-8.noarch.rpm
```
#### 安装opendkim
```
$ sudo yum install opendkim
```

### 生成DKIM key (注意修改域名)
```bash
$ export domain=example.com
```

接下来请以root用户运行!!!  
```bash
$ mkdir /etc/opendkim/keys/$domain
$ cd /etc/opendkim/keys/$domain
$ opendkim-genkey -d $domain -s default
$ chown -R opendkim:opendkim /etc/opendkim/keys/$domain
$ echo "default._domainkey.$domain $domain:default:/etc/opendkim/keys/$domain/default.private" >> /etc/opendkim/KeyTable
$ echo "*@$domain default._domainkey.$domain" >> /etc/opendkim/SigningTable
```
生成之后打开``/etc/opendkim/keys/example.com/default.txt`，里面就是DKIM key，需要添加到DNS，主机记录为`default._domainkey`，记录值为括号里面的（去掉引号,换行）,  
例如:

|名称(Name)|记录类型(Type)|记录内容(Value)|注|
|:---:|:---:|:---:|:---:|
|`default._domainkey` |TXT   |`v=DKIM1; k=rsa; p=...`|/|


### 修改openDKIM配置

编辑`/etc/opendkim.conf`  
1. 将`Mode` 改为 `Mode sv`   
2. 将`Domain` 改为 `Domain example.com` (注意替换)   
3. 将所有变量前面的#去掉，但是`KeyFile`,`Statistics`,`ReportAddress`加上`#`   
4. 再把`SigningTable /etc/opendkim/SigningTable`改成`SigningTable refile:/etc/opendkim/SigningTable`  

### 设置Postfix
编辑`/etc/postfix/main.cf`  
加上下面几行  
```conf
# opendkim 配置
smtpd_milters = inet:127.0.0.1:8891
non_smtpd_milters = inet:127.0.0.1:8891
milter_default_action = accept
```
### 重启服务

```bash
$ sudo systemctl restart postfix opendkim && sudo systemctl enable opendkim
```
P.S.: 第一次启动如果出现` Generating default DKIM keys: hostname: Unknown host` 可以在 `/etc/hosts` 上面加上域名，例如：
```hosts
127.0.0.1 example.com localhost localhost.localdomain localhost4 localhost4.localdomain4
```

# 小结
这里将**SMTP(S)**, 即发件服务配置好了, 关于**Dovecot**及后续内容(Part 2),   
**链接在此 -> [基于CentOS 7,使用Postfix & Dovecot搭建邮件服务器 -Part 2](/2018/07/27/install-email-service-on-CentOS-7-part-2-dovecot/)**




# 参考
[CentOS 7.2 部署邮件服务器（Postfix）](https://blog.csdn.net/wh211212/article/details/53040620)*注:作者未注明许可协议*  
[邮件服务器添加SPF、DKIM、DMARC、PTR提高送达率 ](http://lomu.me/post/SPF-DKIM-DMARC-PTR)*注:作者未注明许可协议* *注:此链接未使用HTTPS*  
