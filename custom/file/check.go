package main

import (
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/constant"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/extension"
	"github.com/zouyx/agollo_demo/info"
	"strings"
)

func main() {
	agollo.SetBackupFileHandler(&FileHandler{})
	extension.AddFormatParser(constant.YML, &Parser{})

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

// FileHandler 默认备份文件读写
type FileHandler struct {
}

// WriteConfigFile write config to file
func (fileHandler *FileHandler) WriteConfigFile(config *config.ApolloConfig, configPath string) error {
	//fmt.Println(config.Configurations)
	return nil
}

// GetConfigFile get real config file
func (fileHandler *FileHandler) GetConfigFile(configDir string, appID string, namespace string) string {
	return ""
}

// LoadConfigFile load config from file
func (fileHandler *FileHandler) LoadConfigFile(configDir string, appID string, namespace string, cluster string) (*config.ApolloConfig, error) {
	return &config.ApolloConfig{}, nil
}

// Parser properties转换器
type Parser struct {
}

// Parse 内存内容=>yml文件转换器
func (d *Parser) Parse(configContent interface{}) (map[string]interface{}, error) {
	fmt.Println(configContent)

	m := make(map[string]interface{})
	m["content"] = configContent
	return m, nil
}
