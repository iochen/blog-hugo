---
title: "ICMP首部与 IP首部 校验和的计算"
tags:
  - network
  - icmp
  - ip
  - Go(lang)
categories:
  - Algorithm
keywords:
  - checksum
  - icmp
  - header
  - ip
  - 详细
  - 校验和
  - 算法
  - 计算
  - 方法
  - Go
  - golang
description: "ICMP首部 与 IP首部 校验和的计算"

date: 2020-08-03T17:55:59+08:00
imgs:
cover:
comments: true
single: false
code: true
math: true
typora-root-url: ../
---

本篇文章详细讲解了计算 **ICMP 首部** 与 **IP 首部** 的计算方法，以及其校验方法

<!--more-->

## 开头

以下内容是一段 **IPv4** 的示例首部数据包

```hex
4500 0073 0000 4000 4011 *B861* C0A8 0001
C0A8 00C7
```
_注：被`*`引起的部分为需要计算的校验和_

既然是要计算校验和，那开始肯定是不知道的，所以我们默认为`0000`，所以，我么就是要对下面这一串数据进行校验和操作。

```hex
4500 0073 0000 4000 4011 0000 C0A8 0001
C0A8 00C7
```

## 第一步

每**16**字节相加求和，于是我们就有了   

$$\mathtt{0x4500}+\mathtt{0x0073}+\mathtt{0x0000}+...+\mathtt{0x00C7}=\mathtt{0x2479C}$$  

### **Go**代码实现

注：传入`data []byte`，大端编码

```go
var	sum    uint32
var	length = len(data)
var	index  int
for length > 1 {
	sum += uint32(data[index])<<8 + uint32(data[index+1])
	index += 2
	length -= 2
}
// 如果len(data)为奇数
if length > 0 {
	sum += uint32(data[index])
}
```

## 第二步

由于我们的校验码只有**16**bit，所以我们就要把上面的$\mathtt{0x2479C}$转为**16**bit  
`2479C`可以转化为

```
0002 479C
```

那我们可以将两个相加，即   

$$\mathtt{0x0002}+\mathtt{0x479C}=\mathtt{0x479E}$$  

如果这样算下来还是超过**16**bit，则需再次进行此运算。
那么最多需要几次运算呢，我们可以算一下。

### 极端情况

由谷歌可知，**TCP**包最大为`65535 bytes`，以**16**bit 为一个小节的话，也就是`32767.5`个小节。
假设每个小节都是$\mathtt{0xFFFF}$（剩余的半个小节以$\mathtt{0x00FF}$计算）。  
已知`32767`的十六进制为$\mathtt{0x7FFF}$，那求和最大是
  
$$\mathtt{0x7FFF}\times{}\mathtt{0xFFFF}+\mathtt{0x00FF}=\mathtt{0x7FFE8100}
$$

第一次进行第二步运算后结果为$\mathtt{0x100FE
}$，第二次为$\mathtt{0x00FF}$  
所以不难看出，最多需要<mark>**两次**</mark>第二步的运算即可。

### **Go**代码实现

```go
sum = (sum >> 16) + (sum & 0xFFFF)
// 与上一行代码作用相同，
// 超过16bit的部分会在最后被舍弃掉
sum += sum >> 16
```

## 第三步

取反运算，这一步的目的是为了验证的方便。  

$$\overline{\mathtt{0x479E}} = \mathtt{0xB861}$$

### **Go**代码实现

```go
return uint16(^sum)
```

## 总结

### **Go**代码实现

```go
func Checksum(data []byte) uint16 {
    var	sum    uint32
    var	length = len(data)
    var	index  int
    for length > 1 {
    	sum += uint32(data[index])<<8 + uint32(data[index+1])
    	index += 2
    	length -= 2
    }
    // 如果len(data)为奇数
    if length > 0 {
    	sum += uint32(data[index])
    }

    sum = (sum >> 16) + (sum & 0xFFFF)
    // 与上一行代码作用相同，
    // 超过16bit的部分会在最后被舍弃掉
    sum += sum >> 16
    return uint16(^sum)
}
```