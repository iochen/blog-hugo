---
title: hexo评论系统——Gitment
date: 2018-01-05 19:22:56
categories:
- Tutorials
tags:
- hexo
- gitment
---

由于**hexo**生成的是静态博客，所以有些功能是无法附加上去的，就比如说评论系统。  
我用过的评论系统不多，也只有**Disqus**和**Gitment**，其他的也只是了解一些，我就来谈谈为什么要用**Gitment**吧！  
国内之前是有**多说**的，但由于一些“原因”被迫关闭了。现在国内比较流行的是**畅言**，但毕竟这还是要经过审查的，所以我没有用。  
之所以之前使用**Disqus**，是因为它在国外的应用很广，自己也试用了用，感觉不错，所以之前一直用的是**Disqus**。  
**Gitment**用下来我个人感觉整体还不错，下面，我就来说说应该怎样配置**Gitment**吧。  
<!--more-->

<h1><font color="red"><b>！！由于Gitment不再维护，建议使用<a href="/2018/01/06/use-gitalk-in-hexo/">Gitalk</a>！！</b></font></h1>



## 配置**Gitment**

首先，你得有一个**Github**账号（这个就不多说了吧），下面，打开[***这个链接***](https://github.com/settings/applications/new)，具体填法如下：
```cpp
Application name //应用名称，随便填
Homepage URL //没有过多要求，可以填自己的博客地址
Application description //应用描述，描述一下，无要求
Authorization callback URL //这个就有要求了，填自己要使用Gitment的博客地址，不可乱填
```
接着，你就可以得到`Client ID`和`Client Secret`，之后会用到的。接下来，我们回到hexo的主题配置里继续！


### 如何配置

如果在主题的配置文件里有`gitment` ，就继续往下看吧！如果没有，请翻到底部，或点击左边的`附：如何添加Gitment功能`。
（我只用过**Next**主题，只知道**NexT**主题是有的）  
怎么填呢？见下面

```cpp
gitment:
  enable: true //是否打开
  count: true  //是否计数
  cleanly: false //是否隐藏'Powered by Gitment'字样
  language: zh-Hans //语言
  github_user: geedme //你的GitHub用户名或ID
  github_repo: gitment //你要存放的库，下文会详细写
  client_id: xxxxx //上文的的Client Id

  client_secret: xxxxxx //上文的Client Secret
```
*如需复制，请删除 // 及后方内容，或将 // 改为 # *  
*我之前踩过一个坑，就是 冒号 和后面的内容之间一定要加空格！否则会出错的！！*


至于还有几个选项没有写，一是因为不常用，二是因为……我也没用过……不能误导大家。  
可能`repo`说的不太明白（其实是我第一次没搞懂），我这里详细说一下。其实**Gitment**是基于**GitHub**的**issues**功能的，所以，你要为他建一个库或用现成的库，我个人建议新建一个，而`repo`就是你要用的库的名称。 比如，我就为**Gitment**专门建了一个叫`gitment`的库，所以在`github_repo: `处填`gitment`。


### 附：如何添加Gitment功能
*附加内容未亲自实验过，内容来源于网络*

见[此网站](https://dj9399.github.io/post/Gitment%E5%AE%89%E8%A3%85%E9%85%8D%E7%BD%AE/)(懒得写了……)

## 其它

这是我写的第一篇"有意义"的博文，欢迎大家在下方评论！
