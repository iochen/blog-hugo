---
title: "在 Go 中获取 GOPATH 的最佳方案"  
tags:  
- Go(lang)
categories:  
- Tutorials
keywords:
- Go
- golang
- GOPATH
- 获取GOPATH
- 教程
description: "在 Go 中获取 GOPATH 的最佳方案"  

date: 2020-07-17T20:47:28+08:00  
imgs:
cover:  
comments: true  
single: false
code: true  
math: false  
typora-root-url: ../  
---

此文章主要讲述了在 golang 程序中完美获取 GOPATH 的方法

<!--more-->

## 初级
golang 的最初几个版本都是依赖于环境变量中的 **GOPATH** 来进行判断的，
所以在这种情况下，只需要
```go
os.Getenv("GOPATH") // import "os"
```
即可，最多再解析一下
```go
sep := string(filepath.ListSeparator)
pathList := strings.Split(path, sep) // import "strings"
```

## 中级
我们有时发现，即使是没有 **GOPATH** 这个 env，也是可以正常运行。
经阅读 golang 源代码后，不难发现，其有默认值。  
[golang 源代码](https://golang.org/src/go/build/build.go#L273):
```go
c.GOPATH = envOr("GOPATH", defaultGOPATH())
````
不难发现，其实我们 **build** 包中就已有 envOr 的实现，于是，
现在的我们只需要
```go
build.Default.GOPATH // import "go/build"
```
即可

## 高级
我们发现
```shell
$ go env -w GOPATH=/foo/bar
```
时，并未将 `/foo/bar` 写入环境变量，而是写入了 env 文件，经
查看源代码后，不难发现 [EnvFile 的实现](https://github.com/golang/go/blob/c4f2a9788a7be04daf931ac54382fbe2cb754938/src/cmd/go/internal/cfg/cfg.go#L151)
```go
// EnvFile returns the name of the Go environment configuration file.
func EnvFile() (string, error) {
	if file := os.Getenv("GOENV"); file != "" {
		if file == "off" {
			return "", fmt.Errorf("GOENV=off")
		}
		return file, nil
	}
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	if dir == "" {
		return "", fmt.Errorf("missing user-config dir")
	}
	return filepath.Join(dir, "go/env"), nil
}
``` 

于是，最终我们有了
```go
func envFile() (string, error) {
	if file := os.Getenv("GOENV"); file != "" {
		if file == "off" {
			return "", fmt.Errorf("GOENV=off")
		}
		return file, nil
	}
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	if dir == "" {
		return "", fmt.Errorf("missing user-config dir")
	}
	return filepath.Join(dir, "go", "env"), nil
}

func getRuntimeEnv(key string) (string, error) {
	file, err := envFile()
	if err != nil {
		return "", err
	}
	if file == "" {
		return "", fmt.Errorf("missing runtime env file")
	}
	var data []byte
	var runtimeEnv string
	data, err = ioutil.ReadFile(file)
	envStrings := strings.Split(string(data), "\n")
	for _, envItem := range envStrings {
		envItem = strings.TrimSuffix(envItem, "\r")
		envKeyValue := strings.Split(envItem, "=")
		if strings.ToLower(envKeyValue[0]) == strings.ToLower(key) {
			runtimeEnv = envKeyValue[1]
		}
	}
	return runtimeEnv, nil
}

goPath, envErr := getRuntimeEnv("GOPATH")
if envErr != nil {
	fmt.Println("Failed: please set '$GOPATH' manually, or use 'datapath' option to specify the path to your custom 'data' directory")
	return
}
if goPath == "" {
	goPath = build.Default.GOPATH
}

sep := string(filepath.ListSeparator)
pathList := strings.Split(path, sep)
```

## 相关
Related Pull Request: [https://github.com/v2ray/domain-list-community/pull/637](https://github.com/v2ray/domain-list-community/pull/637)  