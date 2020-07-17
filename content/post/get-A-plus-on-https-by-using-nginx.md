---
title: nginx配置HTTPS并取得A+
date: 2018-07-11 19:25:20
tags:
- nginx
- server
- Linux
- HTTPS
- safe
categories:
- Apps
- nginx
typora-root-url: ../
---
主要是我自己的nginx配置, 并且HTTPS取得**A+**.  

<!--more-->
先来展示一下:

![58b2543b-5fc7-43a3-9473-f30cf6218ced](/img/get-A-plus-on-https-by-using-nginx/58b2543b-5fc7-43a3-9473-f30cf6218ced.png)



# 直接上配置!

```nginx
server {
    listen 80;
    listen [::]:80;
    server_name domain.ltd; # 你的域名
    return 301 https://domain.ltd$request_uri; # 你的域名
    error_page 497 =301 https://domain.ltd$request_uri; # 你的域名
}

server {
    listen 443  ssl http2;
    listen [::]:443  ssl http2;
    server_name domain.ltd; # 你的域名
    add_header Strict-Transport-Security "max-age=63072000; includeSubDomains; preload";
    ssl        on;
    ssl_certificate /etc/certificate_path/file.pem ; # 你的证书路径
    ssl_certificate_key /etc/certificate_key_path/file.pem ; # 你的证书key路径
    ssl_protocols TLSv1.3 TLSv1.2;
    ssl_stapling on;
    ssl_ciphers "ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-SHA256:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA:ECDHE-ECDSA-AES256-SHA:AES128-GCM-SHA256:AES256-GCM-SHA384:AES128-SHA256:AES256-SHA256:AES128-SHA:AES256-SHA:AES:CAMELLIA:DES-CBC3-SHA:!aNULL:!eNULL:!EXPORT:!DES:!RC4:!MD5:!PSK:!aECDH:!EDH-DSS-DES-CBC3-SHA:!EDH-RSA-DES-CBC3-SHA:!KRB5-DES-CBC3-SHA";
    ssl_prefer_server_ciphers on;
    ssl_session_cache shared:SSL:10m;
    ssl_session_timeout 10m;
    ssl_dhparam /etc/myssl/dhparam.pem; # 你的DH证书路径,下面会说
    charset utf-8;
    access_log  /var/log/nginx/access.log; # 可自行修改
    error_log   /var/log/nginx/error.log; # 可自行修改

    location / {
         root   /usr/share/nginx/path; # 你的网站根目录
         try_files $uri $uri/ /index.php?$uri;
         index  index.php index.html;
    }

    # ERROR pages
    error_page  404        /404.html;
    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /usr/share/nginx/path; # 你的网站根目录
    }

    # GZIP
    gzip on;
    gzip_min_length 1k;
    gzip_comp_level 2;
    gzip_types text/plain application/javascript application/x-javascript text/css application/xml text/javascript application/x-httpd-php image/jpeg image/gif image/png image/jpg;
    gzip_vary on;
    gzip_disable "MSIE [1-6]\.";

    # PHP
    location ~ \.php$ {
        root           html;
        fastcgi_pass   127.0.0.1:9000;
        fastcgi_index  index.php;
        fastcgi_param  SCRIPT_FILENAME /usr/share/nginx/path/$fastcgi_script_name; # 你的网站根目录
        include        fastcgi_params;
    }
}
```

# 必要修改&&说明
## 必要替换:
```
domain.ltd    => 你的域名
/etc/certificate_path/file.pem    => 你的证书路径
/etc/certificate_key_path/file.pem    => 你的证书key路径
/etc/myssl/dhparam.pem    => 你的DH证书路径,下面会说
/usr/share/nginx/path    => 你的网站文件根目录
```
## 个别说明
DH证书: 请先安装`openssl`
```
openssl dhparam -dsaparam -out dhparam.pem 4096
```
建议使用**nohup**,生成会花费一定时间.  
[*nohup ?*](#nohup)

# 附
## nohup
nohup使用:
```bash
$ nohup [应该的命令] &
```
会将屏幕输出内容存放在本地目录下`nohup.out`文件内
