---
title: 从订阅链接生成V2Ray配置 - V2Gen的使用
typora-root-url: ../
date: 2020-02-02 18:26:20
tags:
- V2Ray
- V2Gen
- Vmess
categories:
- Tutorials
description: V2Gen的使用说明
comments: true  
code: true
---

本文章讲述了 v2gen 的实际使用用例，从订阅以及 `vmess://XXXXXXXXXXXX` 链接
中生成 V2Ray json 文件。

<!-- more -->

## 项目地址
详细说明：
[***github.com/iochen/v2gen***](https://github.com/iochen/v2gen/)

## 预览
```data
[ 0] 中继香港C1 Media (HK)(1)      [518ms  (0 errors)]
[ 1] 中继香港C3 Media (HK)(1)      [527ms  (0 errors)]
[ 2] 中继香港C2 Media (HK)(1)      [536ms  (0 errors)]
[ 3] 中继香港C5 Media (HK)(1)      [451ms  (0 errors)]
[ 4] 中继香港C6 Media (HK)(1)      [452ms  (0 errors)]
[ 5] 中继香港G2 Media (HK)(1)      [904ms  (0 errors)]
[ 6] BGP中继香港 2 Media (HK)(1)   [468ms  (0 errors)]
[ 7] BGP中继香港 3 Media (HK)(1)   [778ms  (0 errors)]
[ 8] BGP中继香港 1 Media (HK)(1)   [881ms  (0 errors)]
[ 9] 中继香港G1 Media (HK)(1)      [1.35s  (1 errors)]
...
[50] 日本中继 3 Media (JP)(1)      [641ms  (0 errors)]
=====================
Please Select:
```

## 安装
请注意，本程序并没有**GUI**（图形）界面
### 方法一：从Release中下载
- [Linux AMD64](https://github.com/iochen/v2gen/releases/latest/download/v2gen_amd64_linux)
- [Linux 386](https://github.com/iochen/v2gen/releases/latest/download/v2gen_386_linux)
- [Linux ARM64](https://github.com/iochen/v2gen/releases/latest/download/v2gen_arm64_linux)
- [Linux ARM](https://github.com/iochen/v2gen/releases/latest/download/v2gen_arm_linux)


- [Windows AMD64](https://github.com/iochen/v2gen/releases/latest/download/v2gen_amd64_windows.exe)
- [Windows 386](https://github.com/iochen/v2gen/releases/latest/download/v2gen_386_windows.exe)  
然后请将相应文件加以执行权限并放入到**PATH**中，命名为`v2gen`
### 方法二：从源码自行编译
```shell
$ go get -u -v iochen.com/v2gen/cmd/v2gen
```

## 常用命令
### 测速并选择
```shell
$ v2gen -u 订阅链接 -o V2Ray配置文件路径 
```

### 测速并排序与选择
```shell
$ v2gen -u 订阅链接 -o V2Ray配置文件路径 -sort
```

### 测速并采用最优节点
```shell
$ v2gen -u 订阅链接 -o V2Ray配置文件路径 -best
```

### 随机节点
```shell
$ v2gen -u 订阅链接 -o V2Ray配置文件路径 -r
```

### 设置测速（或延时）目标链接
```shell
$ v2gen -u 订阅链接 -o V2Ray配置文件路径 -dest 目标链接
```

### 打印到控制台以进行管道操作
```shell
$ v2gen -u 订阅链接 -o "-" -n 序号 | jq .
```

## 进阶操作
请到 [GitHub README](https://github.com/iochen/v2gen/blob/master/README_zh_cn.md) 
文件中继续探索！

## TODO
- 重构 v2gen
- 优化多参数时的逻辑判断