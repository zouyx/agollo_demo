package main

import (
	"bytes"
	"fmt"
	"github.com/zouyx/agollo/v3/env"
	"github.com/zouyx/agollo/v3/env/config"
	"net/http"
	"strings"

	"github.com/zouyx/agollo/v3"
)

var namespaces = make(map[string]*struct{}, 0)

func main() {
	agollo.InitCustomConfig(func() (*config.AppConfig, error) {
		return &config.AppConfig{
			AppID:         "testApplication_yang",
			Cluster:       "dev",
			IP:            "http://106.54.227.205:8080",
			NamespaceName: "testyml.yml",
			IsBackupConfig:false,
		}, nil
	})
	agollo.SetLogger(&DefaultLogger{})

	error := agollo.Start()

	fmt.Println("err:", error)

	//agollo.StartWithLogger(&DefaultLogger{})

	http.HandleFunc("/check", GetAllConfig)

	http.ListenAndServe("0.0.0.0:9000", nil)
}

func GetAllConfig(rw http.ResponseWriter, req *http.Request) {
	var config *env.ApolloConnConfig
	for k, v := range env.GetCurrentApolloConfig() {
		if config == nil {
			config = v
		}
		namespaces[k] = &struct{}{}
	}
	n := req.URL.Query().Get("namespace")
	if n != "" {
		namespaces[n] = &struct{}{}
	}

	arr := make([]string, 0)
	for k := range namespaces {
		arr = append(arr, k)
	}

	var namespaceName = strings.Join(arr, ",")

	var buffer bytes.Buffer
	buffer.WriteString("<html>")
	buffer.WriteString("<meta http-equiv=\"refresh\" content=\"3\">")

	key := req.URL.Query().Get("key")
	if key == "" {
		buffer.WriteString(fmt.Sprintf("AppId : %s  <br/>", config.AppID))
		buffer.WriteString(fmt.Sprintf("Cluster : %s <br/>", config.Cluster))
		buffer.WriteString(fmt.Sprintf("ReleaseKey : %s <br/>", config.ReleaseKey))

		namespaces := strings.Split(namespaceName, ",")
		for _, namespace := range namespaces {
			writeConfig(&buffer, namespace)
		}
	}

	//buffer.WriteString(fmt.Sprintf("NamespaceName : %s <br/>", "testjson.json"))
	//buffer.WriteString("Configurations: <br/>")
	//cache := agollo.GetConfig("testjson.json")
	buffer.WriteString("</html>")

	rw.Write(buffer.Bytes())
}

func writeConfig(buffer *bytes.Buffer, namespace string) {
	buffer.WriteString(fmt.Sprintf("NamespaceName : %s <br/>", namespace))
	buffer.WriteString("Configurations: <br/>")
	cache := agollo.GetConfigCache(namespace)
	cache.Range(func(key, value interface{}) bool {
		buffer.WriteString(fmt.Sprintf("key : %s , value : %s <br/>", key, string(value.([]byte))))
		return true
	})
}

type DefaultLogger struct {
}

func (this *DefaultLogger) Debugf(format string, params ...interface{}) {
	this.Debug(format, params)
}

func (this *DefaultLogger) Infof(format string, params ...interface{}) {
	this.Debug(format, params)
}

func (this *DefaultLogger) Warnf(format string, params ...interface{}) error {
	this.Debug(format, params)
	return nil
}

func (this *DefaultLogger) Errorf(format string, params ...interface{}) error {
	this.Debug(format, params)
	return nil
}

func (this *DefaultLogger) Debug(v ...interface{}) {
	fmt.Println(v)
}
func (this *DefaultLogger) Info(v ...interface{}) {
	this.Debug(v)
}

func (this *DefaultLogger) Warn(v ...interface{}) error {
	this.Debug(v)
	return nil
}

func (this *DefaultLogger) Error(v ...interface{}) error {
	this.Debug(v)
	return nil
}
