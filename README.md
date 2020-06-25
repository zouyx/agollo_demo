Agollo_Demo - Demo For Apollo
================

[Apollo](https://github.com/zouyx/agollo)的使用Demo

主要用于验证Agollo客户端自动刷新功能

Dependence
------------

## 使用go get方式

```
go get -u github.com/cihub/seelog
go get -u github.com/zouyx/agollo
```

## 使用gopm方式(推荐)

```
gopm get github.com/cihub/seelog -v -g
gopm get github.com/zouyx/agollo -v -g
```

推荐[gopm](https://github.com/gpmgo/gopm)，不需要翻墙

Installation
------------

如果还没有安装Go开发环境，请参考以下文档[Getting Started](http://golang.org/doc/install.html) ，安装完成后，请执行以下命令：

## Mac/Linux

``` shell
./build.sh
```

## Windows

``` shell
```

*请注意*: 最好使用Go 1.8进行编译

Run
------------
完成编译后

- 进入build文件夹，配置app.properties（参考：[使用配置](https://github.com/zouyx/agollo/wiki/使用指南)）
- 执行
  - Mac/Linux : ./agollo-demo
  - Windows : agollo-demo.exe
  
