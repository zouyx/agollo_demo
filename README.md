agollo_demo - Demo For Apollo
================

[agollo](https://github.com/apolloconfig/agollo) 的使用Demo

主要用于验证 Agollo 客户端自动刷新功能

Dependence
------------

## 使用go get方式

```
go mod
```

## 使用gopm方式(推荐)

```
gopm get github.com/cihub/seelog -v -g
gopm get github.com/apolloconfig/agollo -v -g
```

推荐[gopm](https://github.com/gpmgo/gopm)，不需要翻墙

Installation
------------

如果还没有安装Go开发环境，请参考以下文档[Getting Started](http://golang.org/doc/install.html) ，安装完成后，请执行以下命令：

## 用法

* [网页版](web)
* [hello world](helloworld)
* [自定义用法](custom)
    * [自定义日志组件](custom/log)
    * [自定义缓存组件](custom/cache)
    * [监听变更事件](custom/listener)

## 初始化方法
### 创建应用 
appid：agollo-test
应用名称：agollo-test 
#### 创建私有 Namespace

##### testjson.json
```json
{
  "a":"b",
  "c":"d"
}
```
##### yml
```yml
application:
    name: agollo-test
    port: 8000
```

### 创建加密应用
appid：agollo-secret-test
应用名称：agollo-secret-test

##### application
```
connect: 10
quatity: 5
```

#####　开启密钥