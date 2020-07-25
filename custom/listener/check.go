package main

import (
	"fmt"
	"github.com/zouyx/agollo/v3"
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
	c2 := &CustomChangeListener{}
	c2.wg.Add(5)
	agollo.AddChangeListener(c2)

	error := agollo.Start()

	fmt.Println("err:", error)

	c2.wg.Wait()
	writeConfig(c.NamespaceName)
}

func writeConfig(namespace string) {
	cache := agollo.GetConfigCache(namespace)
	cache.Range(func(key, value interface{}) bool {
		fmt.Println("writeConfig key : ", key, ", value :", value)
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