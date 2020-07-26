---
title: Hexo中Gitalk配置使用教程-可能是目前最详细的教程
date: 2018-01-06 19:22:20
categories:
- Tutorials
tags:
- hexo
- gitalk
comments: true
---

**Gitalk**的详细配置与示例

<!--more-->

# GitHub部分

首先，你得有一个**Github**账号（这个就不多说了吧），下面，打开[***github.com/settings/applications/new***](https://github.com/settings/applications/new)，具体填法如下：
```cpp
Application name //应用名称，随便填
Homepage URL //没有过多要求，可以填自己的博客地址
Application description //应用描述，描述一下，无要求
Authorization callback URL //这个就有要求了，填自己要使用Gitalk的博客地址，不可乱填
```
接着，你就可以得到`Client ID`和`Client Secret`，之后会用到的。接下来，我们回到hexo的主题配置里。

# 文件修改

## swig文件修改法

在这里，我以**NexT**主题做示范，如果发现你的主题大部分以`.ejs` 结尾 请见 [ejs文件修改法](#ejs文件修改法)

### 添加swig代码

在主题的`\layout\_third-party\comments`目录中，新建一个`gitalk.swig`文件，文件内容如下：

```java
{% if page.comments && theme.gitalk.enable %}

  <link rel="stylesheet" href="https://unpkg.com/gitalk/dist/gitalk.css">

  <script src="https://unpkg.com/gitalk/dist/gitalk.min.js"></script>

   <script type="text/javascript">
        var gitalk = new Gitalk({
          clientID: '{{ theme.gitalk.ClientID }}',
          clientSecret: '{{ theme.gitalk.ClientSecret }}',
          repo: '{{ theme.gitalk.repo }}',
          owner: '{{ theme.gitalk.owner }}',
          admin: {{ theme.gitalk.adminUser }},
          id: {{ theme.gitalk.ID }},
          labels: {{ theme.gitalk.labels }},
          perPage: {{ theme.gitalk.perPage }},
          pagerDirection: '{{ theme.gitalk.pagerDirection }}',
          createIssueManually: {{ theme.gitalk.createIssueManually }},
          distractionFreeMode: {{ theme.gitalk.distractionFreeMode }}
        })
        gitalk.render('gitalk-container')           
       </script>
{% endif %}

```

还是，在主题的`\layout\_third-party\comments`目录中，找到一个`index.swig`的文件，打开，添加这一行代码：
```java
{% include 'gitalk.swig' %}
```
接着，在主题的`\layout\_partials`目录中，找到`comments.swig`文件，打开，找到
```java
{% elseif theme.valine.appid and theme.valine.appkey %}
  <div class="comments" id="comments">
  </div>  
{% endif %}
```
在空了一行的地方加上以下代码：
```java
{% elseif theme.gitalk.enable %}
<div id="gitalk-container"></div>
```
### 进行配置

接下来，转到主题的配置文件里，加上这一段：

```js
gitalk:
  enable: true
  ClientID: xxxxxx
  ClientSecret: xxxxxxxxxxxx
  repo: gitalk
  owner: iosite
  adminUser: "['iosite']"
  ID: location.pathname
  labels: "['Gitalk']"
  perPage: 15
  pagerDirection: last
  createIssueManually: true
  distractionFreeMode: false
```
*部分功能未添加，因为大部分人应该也用不到，这里就不多说了，如果需要更多，请见[官方文档](https://github.com/gitalk/gitalk)

各个功能：
```js
  enable: true #指的是是否开启Gitalk
  ClientID: xxxxxx #之前的Client ID
  ClientSecret: xxxxxxxxxxxx #之前的Client Secret
  repo: gitalk #你要存放的项目名，下文会详细再说
  owner: iosite #这个项目名的拥有者（GitHub账号或组织）
  adminUser: "['iosite']" #管理员用户，下文也会详细讲
  ID: location.pathname #页面ID，不知道就默认的就好了
  labels: "['Gitalk']" #GitHub issues的标签，下面会详细说
  perPage: 15 #每页多少个评论
  pagerDirection: last #排序方式是从旧到新（first）还是从新到旧（last）
  createIssueManually: true #如果当前页面没有相应的 isssue ，且登录的用户属于 admin，则会自动创建 issue。如果设置为 true，则显示一个初始化页面，创建 issue 需要点击 init 按钮。
  distractionFreeMode: false #是否启用快捷键(cmd|ctrl + enter) 提交评论.
```
详细说明：  
- repo  
  **Gitalk**是基于**GitHub**的**issues**功能的，所以，你要为他建一个库或用现成的库，我个人建议新建一个，而`repo`就是你要用的库的名称。 比如，我就为**Gitalk**专门建了一个叫`gitalk`的库，所以在`repo: `处填`gitalk`。  

- adminUser  
  即管理员帐号。如果你是个人账号，那么这=里就填你的账户名和协作者的账户名。

  以数组形式，外面再加一层引号

- labels  
  主要是说一下，如果你之前用的是**Gitment**，又有评论，把这里改成`gitment`即可，如果没有，默认的就可以啦！以数组形式，外面再加一层引号

### 样式修改

最后，在`layout/_third-party/comments/index.swig`中添加这样一行：

```css
.gt-container a{border-bottom: none;}
```
完成！

## ejs文件修改法

### 修改源码

添加`layout/_partial/post/gitalk.ejs`

```html
  <link rel="stylesheet" href="https://unpkg.com/gitalk/dist/gitalk.css">
  
  <script src="https://unpkg.com/gitalk/dist/gitalk.min.js"></script>

   <script type="text/javascript">
        var gitalk = new Gitalk({
          clientID: '<%=theme.gitalk.ClientID%>',
          clientSecret: '<%=theme.gitalk.ClientSecret%>',
          repo: '<%=theme.gitalk.repo%>',
          owner: '<%=theme.gitalk.owner%>',
          admin: <%=theme.gitalk.adminUser%>,
          id: <%=theme.gitalk.ID%>,
          labels: <%=theme.gitalk.labels%>,
          perPage: <%=theme.gitalk.perPage%>,
          pagerDirection: '<%=theme.gitalk.pagerDirection%>',
          createIssueManually: <%=theme.gitalk.createIssueManually%>,
          distractionFreeMode: <%=theme.gitalk.distractionFreeMode%>
        })

        gitalk.render('gitalk-container')
       </script>
```

在 `layout/_partial/article.ejs` 倒数第二行加入

```html
  <% if (theme.gitalk.enable){ %>
       <div id="gitalk-container"></div>
       <%- include post/gitalk.ejs %>
  <% } %>

```

配置修改法见[上文](#进行配置)，与swig格式大同小异


Demo就在下面，欢迎留言评论！
