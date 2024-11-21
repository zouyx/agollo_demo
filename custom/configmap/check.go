package main

import (
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/configmap"
	"github.com/zouyx/agollo_demo/info"
	"strings"
)

func main() {
	// enable k8sManager
	manager, err := configmap.GetK8sManager("default")
	if err != nil {
		panic(err)
	}

	configMapHandler := configmap.NewConfigMapHandler(manager)
	agollo.AddBackupFileHandler(configMapHandler, 1)

	client, err := agollo.Start()

	if err != nil {
		fmt.Println("err:", err)
		panic(err)
	}

	split := strings.Split(info.Namespace, ",")
	for _, n := range split {
		checkKey(n, client)
	}
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
