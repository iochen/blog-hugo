---
title: "做一个简易的马克思发生器"  
tags:  
categories:  
keywords:
description: "Build a Simple Marx Generator"  

date: 2021-04-18T10:00:21+08:00  
imgs:
cover:
comments: true  
typora-root-url: ../
single: false
code: true  
draft: false
math: false  
---

你喜欢一些能产生火花的东西，比如说特斯拉线圈吗？但你又没有时间，金钱或耐心来打造一款自己的高压发生设备。

<!--more-->

那么，这里就有一个简单的项目，使你可以创造自己想要的又大，又长，能让邻居找上门来的项目，这甚至比特斯拉线圈要来得有趣！

## 工作原理

马克思发生器由多个单元组成，每单元由 电阻，电容和火花隙 组成，它们长得像这样：

![](/img/build-a-simple-marx-generator/FG8ASBFFK75D8CI.jpg)

（译者注：上图中，输出端从 左下 而非 左上 引线可以得到更好效果）

![](/img/build-a-simple-marx-generator/FNP7CAFFKB3DRCR.jpg)

![](/../../../../../Desktop/FIMRNEOFKB3DRCT.jpg)

![](/img/build-a-simple-marx-generator/FWBV587FKB3DRCS.jpg)

电容通过电阻并联充电，因此每个电容两端电压都为输入电压。 当所有火花隙都接通（有火花产生）时，电容就像是串联了起来一样。所以说，输出电压就等于输入电压乘以电容的数目，从而在装置末端产生长长的火花。

电阻b（Rb）具有镇流作用，用于防止在第一个间隙上形成连续的电弧，以防马克思发生器因此着火。 电阻的阻值主要取决于输入电源。在这里，我们将使用 10MΩ 电阻。您可以通过降低 电阻b（Rb）的阻值以尝试获得频率更高的放电现象，避免在火花隙上形成电弧。 

现在，既然知道它的原来，那么就来试一试吧！

## 材料

我使用了手头的材料。

最佳选项是：陶瓷电容 和 1MΩ 1W 电阻

陶瓷电容最适合该项目，因为它们价格便宜，容易获得，并且适合用于脉冲放电电路。聚丙烯电容也可以，但价格昂贵，高额定电压的聚丙烯电容很难找到。

电阻应为 1MΩ 和 1W 左右的类型，普通的1/4W 电阻非常不合适，因为它们容易引起电弧，从而烧坏马克思发生器，因此请使用 1W 玻璃釉电阻或碳膜电阻。

如有困惑，请尝试先进行小规模实验，然后再购买更多零件，以测试性能如何。

我使用的是：

- 1nF, 4kV 陶瓷电容 * 10
- 1MΩ, 2W, 500V 碳膜电阻 * 20
- 电阻b（Rb）使用了 4.7MΩ, 1W, 3.5kV 玻璃釉电阻 * 2

![](/img/build-a-simple-marx-generator/FR3XYRKFK917304.jpg)

你也可以为马克思发生器制作底座，最佳底座是玻璃（但易碎），陶瓷（易碎），塑料（理想，但有些可能会导电）或任何您认为不错的东西！我采用的是木头，这是一个非常糟糕的主意（因为木头是导电的），但是我别无选择，要么是木制底座，要么没有底座。

你可以使用较大的电容来获得更大，更响亮的火花，但充电时间也会更长！

你电容两端电压常常可以超出其额定电压，但这显然会增加发生事故的机率！

别忘了，输出端的火花大小取决于输入电压和级数...

## 工具

![Also You Need the Tools!](/img/build-a-simple-marx-generator/FSA2723FITUIXBN.jpg)

你需要一把顺手的电烙铁，否则构建一个马克思发生器将十分困难，并且由于电晕放电，得到的结果会远远不如预期。别忘了，助焊剂也是你的好帮手。

热胶枪也非常方便，它可以固定几乎所有东西。

当然，除此之外，常用工具也很必要，这里不再赘述。

## 电源

![Power Supply](/img/build-a-simple-marx-generator/FG9UBFOFKB3DRCU.jpg)

<details><summary>more pictures</summary>

![Power Supply](/img/build-a-simple-marx-generator/FV09JCIFKB3DRCV.jpg)

![Power Supply](/img/build-a-simple-marx-generator/F67VW5QFLJJUT9K.jpg)

</details>

马克思发生器需要 4-8kV 的电压，但只需很小的电流。这种电源可在空气净化器，复印机和激光打印机中找到。

我自己制作了一个的 6kV 的电源。它采用了一个小型自制 450V 逆变器，后方连接 18级电压倍增器，可得到约 7kV 的电压。

马克思发生器也没有正负极之分，所以直接连接上去即可。

## 组装

![The Construction of the Marx Generator.](/img/build-a-simple-marx-generator/F3G1YQAFK917303.jpg)

<details><summary>more pictures</summary>

![The Construction of the Marx Generator.](/img/build-a-simple-marx-generator/FPDHNBGFK917302.jpg)

![The Construction of the Marx Generator.](/img/build-a-simple-marx-generator/FQTXHZRFK917301.jpg)

![The Construction of the Marx Generator.](/img/build-a-simple-marx-generator/F9D56RAFK917300.jpg)

![The Construction of the Marx Generator.](/img/build-a-simple-marx-generator/F78CEL5FK9172ZX.jpg)

![The Construction of the Marx Generator.](/img/build-a-simple-marx-generator/F7Q5U2JFK9172ZZ.jpg)

</details>

每个人的方式都会有所不同，而我将把自己的方式告诉你：

我先将电容用热熔胶粘在基座上。

然后，我电烙铁将 1M 的电阻焊到电容上，并剪断多余的脚。但是我没有剪断电容的引脚，待会你就知道为什么了

之后，我将 4.7M 电阻焊了上去。

最后，请确保焊点完美，并确保除放电处外，没有突出的尖端。 否则，会发生尖端放电，削弱电弧强度。

## 火花隙

![Spark Gaps](/img/build-a-simple-marx-generator/FOU392NFK9172ZY.jpg)

这就是为什么之前我说不要剪掉电容的引脚，因为我们要用它构成火花隙

将引脚弯曲成**V**形，并使**V**形引脚的末端靠在一起，大约2毫米的距离。

但是不要将末端靠在一起，否则会导致大量放电，使得马克思发生器无法正常工作

## 成果

![Testing... Testing...](/img/build-a-simple-marx-generator/FSYXZVIFKB3DRD8.jpg)

![Testing... Testing...](/img/build-a-simple-marx-generator/FAHXS7LFKB3DRCQ.jpg)

![img](/img/build-a-simple-marx-generator/FG4G5C1FKB3DRCL.jpg)



## 译者注

本文原文、图片来自 [Build a Simple Marx Generator](https://www.instructables.com/Build-a-simple-Marx-Generator/)，由 [Richard Chen's Blog](https://iochen.com/) 翻译

本文以 **CC BY-NC-SA 4.0** 许可进行分发