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
		NamespaceName:  "testjson.json",
		IsBackupConfig: false,
		Secret:         "6ce3ff7e96a24335a9634fe9abca6d51",
	}

	agollo.SetBackupFileHandler(&FileHandler{})

	client,err:=agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})

	if err!=nil{
		fmt.Println("err:", err)
		panic(err)
	}

	checkKey(c.NamespaceName,client)
}


func checkKey(namespace string,client *agollo.Client) {
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

// FileHandler 默认备份文件读写
type FileHandler struct {
}

// WriteConfigFile write config to file
func (fileHandler *FileHandler) WriteConfigFile(config *config.ApolloConfig, configPath string) error {
	fmt.Println(config.Configurations)
	return nil
}

// GetConfigFile get real config file
func (fileHandler *FileHandler) GetConfigFile(configDir string, appID string, namespace string) string {
	return ""
}

//LoadConfigFile load config from file
func (fileHandler *FileHandler) LoadConfigFile(configDir string, appID string, namespace string) (*config.ApolloConfig, error) {
	return &config.ApolloConfig{}, nil
}