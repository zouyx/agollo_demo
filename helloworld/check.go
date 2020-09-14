package main

import (
	"fmt"
	"github.com/zouyx/agollo/v4"
	"github.com/zouyx/agollo/v4/env/config"
)

func main() {
	c := &config.AppConfig{
		AppID:          "testApplication_yang",
		Cluster:        "dev",
		IP:             "http://106.54.227.205:8080",
		NamespaceName:  "dubbo",
		IsBackupConfig: false,
		Secret:         "6ce3ff7e96a24335a9634fe9abca6d51",
	}

	client,err:=agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})

	if err!=nil{
		fmt.Println("err:", err)
		panic(err)
	}

	writeConfig(c.NamespaceName,client)
}

func writeConfig(namespace string,client *agollo.Client) {
	cache := client.GetConfigCache(namespace)
	cache.Range(func(key, value interface{}) bool {
		fmt.Println("key : ", key, ", value :", value) 
		return true
	})
}
