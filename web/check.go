package main

import (
	"bytes"
	"fmt"
	"github.com/apolloconfig/agollo/v4/env/config"
	"net/http"
	"strings"

	"github.com/apolloconfig/agollo/v4"
)

var namespaces = make(map[string]*struct{}, 0)
var appConfig= &config.AppConfig{
	AppID:          "agollo-test",
	Cluster:        "dev",
	IP:             "http://106.54.227.205:8080",
	NamespaceName:  "dubbo",
	IsBackupConfig: false,
	Secret:         "7c2ddeb1cd344b8b8db185b3d8641e7f",
}

var client agollo.Client

func main() {
	var err error
	agollo.SetLogger(&DefaultLogger{})

	client,err=agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return appConfig, nil
	})

	fmt.Println("err:", err)

	http.HandleFunc("/check", GetAllConfig)

	http.ListenAndServe("0.0.0.0:9000", nil)
}

func GetAllConfig(rw http.ResponseWriter, req *http.Request) {
	ns := strings.Split(appConfig.NamespaceName, ",")
	for _, n := range ns {
		namespaces[n] = &struct{}{}
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
		buffer.WriteString(fmt.Sprintf("AppId : %s  <br/>", appConfig.AppID))
		buffer.WriteString(fmt.Sprintf("Cluster : %s <br/>", appConfig.Cluster))

		namespaces := strings.Split(namespaceName, ",")
		for _, namespace := range namespaces {
			buffer.WriteString(fmt.Sprintf("ReleaseKey : %s <br/>", appConfig.GetCurrentApolloConfig().GetReleaseKey(namespace)))
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
	cache := client.GetConfigCache(namespace)
	if cache==nil{
		return
	}
	cache.Range(func(key, value interface{}) bool {
		buffer.WriteString(fmt.Sprintf("key : %s , value : %s <br/>", key, value))
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

func (this *DefaultLogger) Warnf(format string, params ...interface{}) {
	this.Debug(format, params)
}

func (this *DefaultLogger) Errorf(format string, params ...interface{}) {
	this.Debug(format, params)
}

func (this *DefaultLogger) Debug(v ...interface{}) {
	fmt.Println(v)
}
func (this *DefaultLogger) Info(v ...interface{}) {
	this.Debug(v)
}

func (this *DefaultLogger) Warn(v ...interface{}) {
	this.Debug(v)
}

func (this *DefaultLogger) Error(v ...interface{}) {
	this.Debug(v)
}
