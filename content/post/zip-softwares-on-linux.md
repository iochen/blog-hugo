---
title: Linux 下的压缩，打包及解压操作
date: 2019-08-18 11:27:43
tags:
- Linux
- zip
- 7zip
- gzip
- tar
categories:
- Linux
---
此文章将介绍 如何在Linux命令行下，压缩，解压，打包一个或多个文件
<!--more-->
# 打包
## .tar
假设有一个文件（夹）`foo`,要打包为`archive.tar`
```shell
# tar cf archive.tar foo
```
假设有多个文件（夹）`foo`和`bar` ,要打包为`archive.tar`  
```bash
tar cf archive.tar foo bar
```

# 压缩
## .tar.gz
假设有文件（夹）`foo`和`bar`,要压缩为`archive.tar.gz`
```bash
tar cfz archive.tar.gz foo bar
```
## .tar.xz
假设有文件（夹）`foo`和`bar`,要压缩为`archive.tar.xz`
```bash
tar cfz archive.tar.xz foo bar
```
## .zip 
假设有**文件**`foo`和`bar`,要压缩为`archive.zip` (！！注意，不适用于文件夹！！)
```bash
zip archive.zip foo bar
```
假设有**文件夹**`foo`和`bar`,要压缩为`archive.zip`
```bash
zip -r archive.zip foo bar
```
## .7z 
假设有文件（夹）`foo`和`bar`,要压缩为`archive.7z`
```bash
7z a archive.7z foo bar
```

# 解压
## .tar.*
解压/解包`.tar.*`文件，假设为`archive.tar.gz`
```bash
tar xf archive.tar.gz
```
## .7z 
解压`.7z`文件，假设为`archive.7z`
```bash
7z e archive.7z
```
## .zip
解压`.zip`文件，假设为`archive.zip`
```bash
unzip archive.zip
```
