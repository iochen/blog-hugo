---
title: "在 Flash 退出舞台后为 Chrome 启用 Adobe Flash"  
tags:  
- Flash
- Chrome
- Translation
categories:  
- Tutorials
keywords:
- 启用 Adobe Flash
- 2021年
description: "在 Flash 退出舞台后为 Chrome 启用 Adobe Flash"  

date: 2021-02-04T19:35:24+08:00  
imgs:
cover:  
comments: true  
single: false
typora-root-url: ../
  
code: true  
math: false  
---



就像你知道的那样（或不知道），Adobe Flash 于 2020年12月31日 退出历史舞台，而自2021年1月12日起完全无法使用！

<!--more-->

但是，对于部分用户以及 IT 人员而言，使用仍只支持 Adobe Flash 的旧系统时会引起灾难。

就算在此之前，出于安全考虑，也仍需要采取特殊步骤来允许 Adobe Flash。

而自 2021年1月12日 起，这些步骤将不再有效。即使选择“允许” Flash，Adobe 也会故意阻止它。你会看到类似下面的内容。

![Adobe Flash 已禁用](/img/enable-adobe-flash-chrome-after-end-of-life/Disabled.png)

## 解决方案

该解决方法仅应用于紧急情况，因为是时候该减少对 Flash 的依赖了。

为了解决此问题，您将需要 Chrome 版本 87 或更早版本，因为版本88不包括 Adobe Flash。

### STEP 1

#### Windows

```
C:\Users\{用户名}\AppData\Local\Google\Chrome\User Data\Default\Pepper Data\Shockwave Flash\System\
```

需要在如上路径创建一个 `mms.cfg` 文件

（请自行替换 `{用户名}`）

#### Mac

```
/Users/{用户名}/Library/Application Support/Google/Chrome/Default/Pepper Data/Shockwave Flash/System/
```

需要在如上路径创建一个 `mms.cfg` 文件

（请自行替换 `{用户名}`）

#### Linux

```
~/.config/google-chrome/Default/Pepper Data/Shockwave Flash/System/
```

需要在如上路径创建一个 `mms.cfg` 文件

### STEP 2

`mms.cfg` 文件内容如下：

```cfg
EnableAllowList = 1
AllowListUrlPattern = http://COMPUTERNAME/
AllowListUrlPattern = https://IP.ADDRESS/
AllowListUrlPattern = *://server.domain.com/
AllowListUrlPattern = ...
```

按需添加 `AllowListUrlPattern = ...`

例如：

- `AllowListUrlPattern = *://google.com/`
- `AllowListUrlPattern = *://192.168.1.1/`

### STEP 3

要访问网站：

1. 创建此文件并编辑保存后，打开 **Chrome**
2. 打开到您要查看的页面
3. 单击地址栏左方的 **小锁** 或 **不安全图标**
4. 将 **Adobe Flash** 从 **询问** 更改为 **允许**
5. **Chrome** 提示时，点击 **刷新**
6. 不要更新，单击 **运行一次**。

（译者注：译者使用的是 英文版，选项名称可能有所出入）

### 其它

（以下信息与原文章无关）

注：博主在 Linux 上实测有效，在其它平台未测试

本文翻译自：[Enable Adobe Flash on Chrome after End of Life](https://www.stephenwagner.com/2021/01/13/enable-adobe-flash-chrome-after-end-of-life/)



---

内容著作权归 [原作者](https://www.stephenwagner.com/) 所有

译文著作权归 [Richard Chen's Blog](https://iochen.com/) 所有

二次分发同时标明 [原文]((https://www.stephenwagner.com/2021/01/13/enable-adobe-flash-chrome-after-end-of-life/)) 与 [译文](https://iochen.com/post/enable-adobe-flash-chrome-after-end-of-life/) 出处