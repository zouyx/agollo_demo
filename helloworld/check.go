package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/zouyx/agollo/v4"
	"github.com/zouyx/agollo/v4/agcache"
	"github.com/zouyx/agollo/v4/env/config"
)

// apolloProvider apollo source
type apolloProvider struct {
	sync.RWMutex
	c         agcache.CacheInterface
	namespace string
	keys      map[string]interface{}
}

func main() {
	p, _ := NewApolloProvider(config.AppConfig{
		AppID:          "hk109",
		Cluster:        "dev",
		IP:             "http://106.54.227.205:8080",
		NamespaceName:  "dubbo",
		IsBackupConfig: false,
		Secret:         "6ce3ff7e96a24335a9634fe9abca6d51",
	})
	for {
		fmt.Printf("p--------------->"+"%+v\n", p)
		time.Sleep(1 * time.Second)
	}
}

// NewApolloProvider get a apollo source singleton, and pull configs at once after init apollo client.
func NewApolloProvider(c config.AppConfig) (*apolloProvider, error) {
	// 初始化apolloProvider实例
	ap := &apolloProvider{
		namespace: c.NamespaceName,
		keys:      make(map[string]interface{}),
	}
	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) { return &c, nil })
	if err != nil {
		return nil, err
	}
	ap.c = client.GetConfigCache(c.NamespaceName)
	return ap, nil
}
