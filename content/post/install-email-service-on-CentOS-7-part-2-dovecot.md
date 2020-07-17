---
title: 基于CentOS 7,使用Postfix & Dovecot搭建邮件服务器 -Part 2
date: 2018-07-27 20:54:21
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
- Linux
- CentOS
---
本文将详细讲解 如何在CentOS 7下，使用Postfix & Dovecot 搭建自己的邮件服务器。  
这是第二部分: Dovecot 与 其它优化。  

<!--more-->

# 前言
在我自己搭建的时候，看了许多网络教程，但博主自己在实现过程中遇到了许多坑，比如教程本身有问题，讲解不明等，所以在此自己写一篇。      

关于**Postfix**及前面内容(Part 1),   
**链接在此 -> [基于CentOS 7,使用Postfix & Dovecot搭建邮件服务器 -Part 1](/2018/07/21/install-email-service-on-CentOS-7-part-1-postfix/)**

# 实践
## **变量设定**
<font color=red>
重要!!!   
**假如你的IP为`8.8.8.8`, 域名为`example.com`, 你想要创建的邮箱为` xxx@example.com`**  
**证书, 密钥请点此 ->** [**证书路径**](/2018/07/21/install-email-service-on-CentOS-7-part-1-postfix/#签发证书)
</font>

## Dovecot
### 安装
```bash
$ sudo yum -y install dovecot
```
### 配置
编辑`/etc/dovecot/dovecot.conf`
```
# 第24行 : 取消注释并修改
protocols = imap pop3 lmtp
# 第30行: 取消注释并修改
listen = *, ::
```
编辑`/etc/dovecot/conf.d/10-auth.conf`
```
# 第10行: 取消注释并修改
disable_plaintext_auth = yes
# 第100行: 修改
auth_mechanisms = plain login
```
编辑`/etc/dovecot/conf.d/10-mail.conf`
```
# 第30行: 取消注释并修改
mail_location = maildir:~/Mailbox
```
编辑`/etc/dovecot/conf.d/10-master.conf`
```
# 第89~93行: 全部注释!!!
# unix_listener auth-userdb {
   #mode = 0666
   #user =
   #group =
# }

# 第96~98行: 取消注释并修改
unix_listener /var/spool/postfix/private/auth {
    mode = 0666
    user = postfix
    group = postfix
}
```
编辑`/etc/dovecot/conf.d/10-ssl.conf`
```
# 第8行: 修改
ssl = required

# 第14~15行: 取消注释并修改,注意替换路径!!! 路径见本文前面 变量设定 部分
ssl_cert = </etc/letsencrypt/live/mail.example.com/cert.pem
ssl_key = </etc/letsencrypt/live/mail.example.com/privkey.pem
```
### 重启Dovecot
```bash
$ sudo systemctl restart dovecot && sudo systemctl enable dovecot
```
## firewall相应配置
```bash
$ sudo firewall-cmd --add-port={110/tcp,143/tcp,995/tcp,993/tcp} --permanent && sudo firewall-cmd --reload
```
**110/tcp**为**POP3**端口, **143/tcp**为**IMAP**端口, **995/tcp**为**POP3S**端口, **993/tcp**为**IMAPS**端口

# 检验
假设你的Linux服务器中, 你的用户名为`user`, 则可按下方配置操作:
![email for user](/img/install-email-service-on-CentOS-7-part-2-dovecot/68eee113-65c4-4cd9-9347-1b1cf2c608e8.png)

# 参考
[CentOS 7.2 部署邮件服务器（Postfix）](https://blog.csdn.net/wh211212/article/details/53040620)*注:作者未注明许可协议*  
[邮件服务器添加SPF、DKIM、DMARC、PTR提高送达率 ](http://lomu.me/post/SPF-DKIM-DMARC-PTR)*注:作者未注明许可协议* *注:此链接未使用HTTPS*  
