package main

import (
	"errors"
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/agcache"
	"github.com/apolloconfig/agollo/v4/env/config"
	"sync"
)

func main() {
	c := &config.AppConfig{
		AppID:          "testApplication_yang",
		Cluster:        "dev",
		IP:             "http://106.54.227.205:8080",
		NamespaceName:  "testyml.yml",
		IsBackupConfig: false,
		Secret:         "6ce3ff7e96a24335a9634fe9abca6d51",
	}

	agollo.SetCache(&DefaultCacheFactory{})
	agollo.SetLogger(&DefaultLogger{})

	client,err:=agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})

	if err!=nil{
		fmt.Println("err:", err)
		panic(err)
	}

	checkKey(c.NamespaceName,client)
}

func checkKey(namespace string,client agollo.Client) {
	cache := client.GetConfigCache(namespace)
	count:=0
	cache.Range(func(key, value interface{}) bool {
		fmt.Println("key : ", key, ", value :", value)
		count++
		return true
	})
	if count<1{
		panic("config key can not be null")
	}
}

//DefaultCache 默认缓存
type DefaultCache struct {
	defaultCache sync.Map
}

//Set 获取缓存
func (d *DefaultCache)Set(key string, value interface{}, expireSeconds int) (err error)  {
	d.defaultCache.Store(key,value)
	return nil
}

//EntryCount 获取实体数量
func (d *DefaultCache)EntryCount() (entryCount int64){
	count:=int64(0)
	d.defaultCache.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	return count
}

//Get 获取缓存
func (d *DefaultCache)Get(key string) (value interface{}, err error){
	v, ok := d.defaultCache.Load(key)
	if !ok{
		return nil,errors.New("load default cache fail")
	}
	return v.([]byte),nil
}

//Range 遍历缓存
func (d *DefaultCache)Range(f func(key, value interface{}) bool){
	d.defaultCache.Range(f)
}

//Del 删除缓存
func (d *DefaultCache)Del(key string) (affected bool) {
	d.defaultCache.Delete(key)
	return true
}

//Clear 清除所有缓存
func (d *DefaultCache)Clear() {
	d.defaultCache=sync.Map{}
}

//DefaultCacheFactory 构造默认缓存组件工厂类
type DefaultCacheFactory struct {

}

//Create 创建默认缓存组件
func (d *DefaultCacheFactory) Create()agcache.CacheInterface {
	return &DefaultCache{}
}


type DefaultLogger struct {
}

func (this *DefaultLogger) Debugf(format string, params ...interface{}) {
	this.Debug(format, params)
}

func (this *DefaultLogger) Infof(format string, params ...interface{}) {
	this.Debug(format, params)
}

func (this *DefaultLogger) Warnf(format string, params ...interface{}) {
	this.Debug(format, params)
}

func (this *DefaultLogger) Errorf(format string, params ...interface{}) {
	this.Debug(format, params)
}

func (this *DefaultLogger) Debug(v ...interface{}) {
	fmt.Println(v)
}
func (this *DefaultLogger) Info(v ...interface{}) {
	this.Debug(v)
}

func (this *DefaultLogger) Warn(v ...interface{}) {
	this.Debug(v)
}

func (this *DefaultLogger) Error(v ...interface{}) {
	this.Debug(v)
}