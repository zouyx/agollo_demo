package main

import (
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/storage"
	"sync"
)

func main() {
	c2 := &CustomChangeListener{}
	c2.wg.Add(2)

	client, err := agollo.Start()
	client.AddChangeListener(c2)

	fmt.Println("err:", err)

	c2.wg.Wait()
	writeConfig("testyml.yml", client)
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
