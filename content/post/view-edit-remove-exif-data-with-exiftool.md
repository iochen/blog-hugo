---
title: 'Linux下使用ExifTool对EXIF的操作'
typora-root-url: ../
date: 2020-02-04 17:48:18
tags:
- ExifTool
- EXIF
- Linux
categories:
- Apps
- ExifTool
description: Linux下使用ExifTool查看、修改与删除图片的EXIF数据
---

# 前言

**EXIF**数据可以包括你的相机信息，快门，光圈，ISO，创建时间，编辑信息，GPS信息等信息，他们会在无心之中泄漏你的隐私，特别是上传照片到互联网上时，那更是有极大的泄漏风险。因此，在上传照片到网路上时，请务必删除**EXIF**信息。  

本篇文章就是要教你在**Linux**下，如何使用**ExifTool**查看、修改与删除图片的**EXIF**数据



# 安装**ExifTool** 

## openSUSE

```bash
zypper in exiftool
```

## Ubuntu

```bash
apt-get install libimage-exiftool-perl
```

## CentOS

```bash
yum install perl-Image-ExifTool
```

## Others

到[**ExifTool**官网](https://exiftool.org/)下载二进制包，解压并进入。如有需要，请链接到`/usr/bin/`



# 使用ExifTool

## 查看**EXIF**

### 直接查看

```bash
exiftool {{InputFile}}
```

可选参数：

* `-a` 查看所有metadata，但博主实测并没有增加多少新的数据，还重复了很多



### 导出为**HTML**

```bash
exiftool -h {{InputFile}} > {{OutputFile}}
```

可选参数：

* `-a` 查看所有metadata



### 导出为json

```bash
exiftool -j {{InputFile}} > {{OutputFile}}
```

可选参数：

* `-a` 查看所有metadata



## 编辑EXIF

### 写入到单文件

```bash
exiftool -{{TAG}}={{Content}}  {{InputFile}}
```

*example:*

```bash
exiftool -comment="Good Quality" -artist="Richard Chen" input.png
```

### 写入到多文件或文件夹

```bash
exiftool -{{TAG}}={{Content}} {{InputFile1}} {{InputFile2}} {{Folder}}
```

*example:*

```bash
exiftool -comment="Good Quality" -artist="Richard Chen" input_1.png input_2.jpg ~/Pictures/
```



## 删除EXIF

```
exiftool -all= {{InputFile}}
```

*example:*

```
exiftool -all= input_1.png input_2.jpg ~/Pictures/
```

### 文件夹下所有图片

此方法适用于多级文件夹

```bash
find {{Folder}} -iname "*.{{Ext}}" -exec exiftool -all= {} +
```

*example:*

```bash
find ~/Pictures/ -iname "*.jpg" -exec exiftool -all= {} +
```



# *References*

[*ExifTool Command-Line Examples*](https://exiftool.org/examples.html)

[*User Contributed Perl Documentation*](https://exiftool.org/exiftool_pod.pdf)

*Other resources on the Internet*





