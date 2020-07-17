---
title: "[Unsolved] Wrong file size"

typora-root-url: ../
date: 2020-02-03 14:30:22
tags:
- Issue
- Linux
categories:
- Linux
- openSUSE
description: wrong file size shown on kde
---

# Environment

* System: **openSUSE** *15.1*
* Desktop: **KDE** *Plasma*



# Details

![image-20200203174433732](/img/filesize-puzzle/image-20200203174433732.png)



Other files under the same folder:

![image-20200203174529757](/img/filesize-puzzle/image-20200203174529757.png)



Copyed to other folder:

![image-20200203174744479](/img/filesize-puzzle/image-20200203174744479.png)



After rebooted:

![image-20200203175030071](/img/filesize-puzzle/image-20200203175030071.png)



File type:

![image-20200203175053415](/img/filesize-puzzle/image-20200203175053415.png)



# My Guess

There is a place, which caches the file size of every file. When the file is generating(not finished), it caches the file size. After moved to other folder, or computer is rebooted, it will refresh the cache.   

But... it's a bug though :)