---
title: "dovecot permission denied 的解决办法"
date: 2019-07-01 11:15:45
description: "出现 doveconf: Fatal: Error in ... ssl_cert: ... Can not open file: Permission denied 的解决办法" 
tags:
- openSUSE
- dovecot
- issue
categories:
- Issues
keywords:
- dovecot
- permission denied
- 错误
- 解决方案
comments: true
---







今天，在**openSUSE Leap 15.0**上配置**dovecot**的时候，出现了这个错误：

```log
doveconf: Fatal: Error in configuration file /etc/dovecot/dovecot.conf  line 7: ssl_cert: /etc/{some path} Can not open file: Permission  denied
```

以下是我的解决办法

<!--more-->

# 环境

我的SSL证书是由**certbot**颁发的，在`/etc/certbot/...`下

# 错误现象

**systemed**开启**dovecot**时，报错

```log
doveconf: Fatal: Error in configuration file /etc/dovecot/dovecot.conf  line 7: ssl_cert: /etc/{some path} Can not open file: Permission  denied
```

# 解决方案

在`profiles/apparmor.d/abstractions/ssl_certs`中对应位置添加

```c
  /etc/certbot/archive/*/cert*.pem r,
  /etc/certbot/archive/*/chain*.pem r,
  /etc/certbot/archive/*/fullchain*.pem r,
```

在`profiles/apparmor.d/abstractions/ssl_keys`对应位置添加

```
  /etc/certbot/archive/*/privkey*.pem r,
```

详细可参考

[https://gitlab.com/iochen/apparmor/commit/3016ffb3367e03ee2129401472d44d5eea4c1fb2](https://gitlab.com/iochen/apparmor/commit/3016ffb3367e03ee2129401472d44d5eea4c1fb2)

[https://gitlab.com/iochen/apparmor/commit/4d275bab696f58e1431d26da642e82adbe092526](https://gitlab.com/iochen/apparmor/commit/4d275bab696f58e1431d26da642e82adbe092526)



# 后续

已在[apparmor官方仓库](https://gitlab.com/apparmor/apparmor)中提出[PR](https://gitlab.com/apparmor/apparmor/merge_requests/397)

# 参考

[LEAP 42.3 Unexpected permissions issue with Dovecot](https://forums.opensuse.org/showthread.php/531740-Unexpected-permissions-issue-with-Dovecot)

