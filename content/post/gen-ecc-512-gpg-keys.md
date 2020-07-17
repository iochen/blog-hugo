---
title: GPG中创建ECC-512公私钥
date: 2019-08-18 12:09:56
tags:
- ECC
- 非对称加密
- GPG
- PGP 
- ECC-512
categories:
- Crypto
- ECC
---
此文章主要讲解
- 如何在GPG中创建ECC-512公私钥
- 解决创建过程中的错误

错误：

```error
gpg: signing failed: Invalid length
gpg: make_keysig_packet failed: Invalid length
Key generation failed: Invalid length
```
# 创建
```
gpg --expert --full-gen-key
```
选择`ECC and ECC`以及加密方式
按提示输入即可
# 错误
```error
gpg: signing failed: Invalid length
gpg: make_keysig_packet failed: Invalid length
Key generation failed: Invalid length
```
## 解决方案
编辑`~/.gpg/gpg.conf`文件
修改（或添加）`cert-digest-algo`为：
```conf
cert-digest-algo SHA256
```
即可

