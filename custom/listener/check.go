package main

import (
	"errors"
	"fmt"
	"github.com/zouyx/agollo/v3"
	"github.com/zouyx/agollo/v3/agcache"
	"github.com/zouyx/agollo/v3/env/config"
	"github.com/zouyx/agollo/v3/storage"
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
	agollo.InitCustomConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	agollo.AddChangeListener(&CustomChangeListener{})

	error := agollo.Start()

	fmt.Println("err:", error)

	writeConfig(c.NamespaceName)
}

func writeConfig(namespace string) {
	cache := agollo.GetConfigCache(namespace)
	cache.Range(func(key, value interface{}) bool {
		fmt.Println("key : ", key, ", value :", string(value.([]byte)))
		return true
	})
}
type CustomChangeListener struct {

}

func (c *CustomChangeListener) OnChange(changeEvent *storage.ChangeEvent) {
	//write your code here
}