---
title: openssl RSA简单使用
date: 2018-02-10 11:44:49
categories:
 - Tutorials
tags:
 - openssl
 - 加解密
 - RSA
 - encrypt
 - decrypt
comments: true
---

最近看到的关于TLS的文章比较多啊，我也来凑凑热闹，来一篇关于非对称加密用法的博客。当然了，这篇文章对不懂的人也没什么用，懂的人也不会看这篇，还是写给自己看吧
<!--more-->
这里的非对称加密用的是**RSA**
# 基本
非对称加密密钥的用法主要是如下几种：
```
加密 → 公钥
解密 → 私钥
签名 → 私钥
解密 → 私钥
```


# 用法
## 生成密钥
### 生成私钥
比如我要生成一个叫`pr.pem`的私钥文件，直接
```
openssl genrsa -out pr.pem
```
就可以了，就像这样：
![生成私钥](/img/openssl-RSA-1/f05d8842-7080-4b4c-bcc2-bbc07ceaefc5.png)

### 提取公钥
比如我要从`pr.pem`里提取`pu.pem`的公钥，直接
```
openssl rsa -pubout -in pr.pem -out pu.pem
```
即可，可以看到：
![提取公钥](/img/openssl-RSA-1/6ed6724b-1fac-415e-bd11-8403680fe04c.png)
### 小结
至此，公钥和私钥就生成好了，来看看吧！  
私钥：  
![私钥](/img/openssl-RSA-1/b567d11b-55c7-4e2e-9629-32a406427301.png)
公钥：  
![公钥](/img/openssl-RSA-1/4f6e2b69-034f-4bdd-a02a-52ef9310e0bd.png)

## 加解密
### 加密
我们先创建一个叫`test`的文件，写点什么
![测试文件](/img/openssl-RSA-1/35036b25-44cf-4565-a645-ff2719ce065e.png)
接下来加密
```
openssl rsautl -pubin -encrypt -in test -inkey pu.pem -out test_en
```
来看看加密后的文件：  
*注：现在用`cat`是看不了的，只是一堆乱码，所以我用`hexdump`查看二进制文件*
![加密后](/img/openssl-RSA-1/a0caf393-a88d-4c94-83d5-4493d6278fa9.png)

### 解密
简单！输入
```
openssl rsautl -decrypt -in test_en -inkey pr.pem -out test_dn
```
又回来了不是吗！  
来看看解密后的文件：
![解密后](/img/openssl-RSA-1/ed16ca35-cee3-44ff-a71d-fbcdfd57b0eb.png)
完全一样！

## 签名与验证
### 签名
我想为`test`文件签名，直接
```
openssl dgst -sha256 -sign pr.pem -out test_sign test
```
即可生成签名文件`test_sign`  


### 验证签名
一样，一句命令！
```
openssl dgst -sha256 -verify pu.pem -signature test_sign test
```
好的，得到一句
![验证通过](/img/openssl-RSA-1/e3b95074-4eee-4d75-907e-682898c8d55f.png)
这就是验证通过了  

如果这个原来数据被一些人修改后，比如：  
*这里我为`test`文件加了一行*
![修改后](/img/openssl-RSA-1/a9569407-189a-411c-9ebe-a9a19031e2c7.png)
就会出现：
![验证失败](/img/openssl-RSA-1/d0ce75e7-7750-4b53-b574-069953fbc965.png)

## 总结
这里我有一些没有说，因为写着写着就写的很详细很长了，具体可以参考[这位](https://www.cnblogs.com/gordon0918/p/5382541.html)的博客。

# 个人总结
其实，我只要记6段代码就好了
```
openssl genrsa -out pr.pem                                            //生成
openssl rsa -pubout -in pr.pem -out pu.pem                            //提取
openssl rsautl -pubin -encrypt -in test -inkey pu.pem -out test_en    //加密
openssl rsautl -decrypt -in test_en -inkey pr.pem -out test_dn        //解密
openssl dgst -sha256 -sign pr.pem -out test_sign test                 //签名
openssl dgst -sha256 -verify pu.pem -signature test_sign test         //验证
```

# 其它
- 这里的密钥对不会在任何地方使用！
- ~~博主的正式公钥：地址：[https://mega.nz/#!sM4GDIxY](https://mega.nz/#!sM4GDIxY)  密钥：`!0bW3hOti7GQHtkXYgXY52o8IK9RZVos7hzdZV8fA4uc`~~ 已弃用


有问题请在下面问！（至于我会不会……😭）
