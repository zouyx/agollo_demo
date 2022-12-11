package main

import (
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
	"sync"
)

func main() {
	c := &config.AppConfig{
		AppID:          "agollo-test",
		Cluster:        "dev",
		IP:             "http://81.68.181.139:8080",
		NamespaceName:  "testyml.yml",
		IsBackupConfig: false,
		Secret:         "7c2ddeb1cd344b8b8db185b3d8641e7f",
	}
	c2 := &CustomChangeListener{}
	c2.wg.Add(5)

	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	client.AddChangeListener(c2)

	fmt.Println("err:", err)

	c2.wg.Wait()
	writeConfig(c.NamespaceName, client)
}

func writeConfig(namespace string, client agollo.Client) {
	cache := client.GetConfigCache(namespace)
	cache.Range(func(key, value interface{}) bool {
		fmt.Println("key : ", key, ", value :", value)
		return true
	})
}

type CustomChangeListener struct {
	wg sync.WaitGroup
}

func (c *CustomChangeListener) OnChange(changeEvent *storage.ChangeEvent) {
	//write your code here
	fmt.Println(changeEvent.Changes)
	for key, value := range changeEvent.Changes {
		fmt.Println("change key : ", key, ", value :", value)
	}
	fmt.Println(changeEvent.Namespace)
	c.wg.Done()
}

func (c *CustomChangeListener) OnNewestChange(event *storage.FullChangeEvent) {
	//write your code here
}
