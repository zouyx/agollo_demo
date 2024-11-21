自定义configmap缓存组件
------------

* 启动

```bash
go run check.go
```

* 如何使用

## 实现
`github.com/apolloconfig/agollo/v4/env/file/file_handler.go`

```go
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
func (fileHandler *FileHandler) LoadConfigFile(configDir string, appID string, namespace string, cluster string) (*config.ApolloConfig, error) {
	return &config.ApolloConfig{}, nil
}
```

## 添加组件

```go
agollo.AddBackupFileHandler(&FileHandler{}, -1)
```
默认有文件缓存，优先级为0，可以使用AddBackupFileHandler添加configmap缓存组件，并自定义优先级