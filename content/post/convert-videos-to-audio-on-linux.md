---
title: Linux下使用FFmpeg,将mp4等格式转mp3格式
date: 2019-08-18 11:09:08
tags:
- FFmpeg
- App Tutorials
categories:
- Tutorials
description: Linux下mp4等格式转mp3格式
keywords:
- Linux
- FFmpeg
- 格式转换
- mp4
- mp3
- 教程
description: "Linux下使用FFmpeg,将mp4等格式转换为mp3格式"  
comments: true  
---


本文章将介绍 Linux下使用FFmpeg,将各类格式互相转换，以及如何批量转换
<!--more-->
# 单个转换
## 命令
```bash
ffmpeg -i {输入文件} {输出文件}
```
## 示例
### mp4 转 mp3
```bash
ffmpeg -i foo.mp4 foobar.mp3
```
### flv 转 mp3
```bash
ffmpeg -i foo.flv foobar.mp3
```
### mp4 转 wav
```bash
ffmpeg -i foo.mp4 foobar.wav
```
### wav 转 mp3
```bash
ffmpeg -i foo.wav foobar.mp3
```

# 批量转换
比如，文件夹下有`*.mp4`文件，要批量转换为`XX.mp3`文件，可以用如下方法
```bash
for i in ./*.mp4
do
ffmpeg -i $i ${i}.mp3
done
```

再比如，文件夹下所有文件文件，都要批量转换为`XX.mp3`文件，可以用如下方法
```bash
for i in ./*
do
ffmpeg -i $i ${i}.mp3
done
```
