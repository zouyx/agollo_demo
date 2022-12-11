package main

import (
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"time"
)

func main() {
	c := &config.AppConfig{
		AppID:          "agollo-test",
		Cluster:        "dev",
		IP:             "http://81.68.181.139:8080",
		NamespaceName:  "dubbo",
		IsBackupConfig: true,
		Secret:         "7c2ddeb1cd344b8b8db185b3d8641e7f",
	}
	agollo.SetLogger(&DefaultLogger{})

	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})

	if err != nil {
		fmt.Println("err:", err)
		panic(err)
	}

	checkKey(c.NamespaceName, client)

	c = &config.AppConfig{
		AppID:          "agollo-test",
		Cluster:        "dev",
		IP:             "http://81.68.181.139:8080",
		NamespaceName:  "dubbo",
		IsBackupConfig: false,
		Secret:         "7c2ddeb1cd344b8b8db185b3d8641e7f",
	}

	client, err = agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})

	if err != nil {
		fmt.Println("err:", err)
		panic(err)
	}

	checkKey(c.NamespaceName, client)

	time.Sleep(5 * time.Second)
}

func checkKey(namespace string, client agollo.Client) {
	cache := client.GetConfigCache(namespace)
	count := 0
	cache.Range(func(key, value interface{}) bool {
		fmt.Println("key : ", key, ", value :", value)
		count++
		return true
	})
	if count < 1 {
		panic("config key can not be null")
	}
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