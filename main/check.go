package main

import (
	"bytes"
	"fmt"
	"github.com/zouyx/agollo"
	"net/http"
	"strings"
)

var NamespaceName="application,product.joe"

func main() {
	//agollo.InitCustomConfig(func () (*agollo.AppConfig, error) {
	//	return &agollo.AppConfig{
	//		AppId:         "testApplication_yang",
	//		Cluster:       "dev",
	//		Ip:            "http://106.12.25.204:8080",
	//		NamespaceName: NamespaceName,
	//	}, nil
	//})
	go agollo.StartWithLogger(&DefaultLogger{})

	http.HandleFunc("/check",GetAllConfig)

	http.ListenAndServe("0.0.0.0:9000",nil)
}

func GetAllConfig(rw http.ResponseWriter,req *http.Request)  {
	config:=agollo.GetCurrentApolloConfig()["application"]

	var buffer bytes.Buffer
	buffer.WriteString("<html>")
	buffer.WriteString("<meta http-equiv=\"refresh\" content=\"3\">")

	key:=req.URL.Query().Get("key")
	if key=="" {
		buffer.WriteString(fmt.Sprintf("AppId : %s  <br/>", config.AppId))
		buffer.WriteString(fmt.Sprintf("Cluster : %s <br/>", config.Cluster))
		buffer.WriteString(fmt.Sprintf("ReleaseKey : %s <br/>", config.ReleaseKey))

		namespaces:=strings.Split(NamespaceName,",")
		for _, namespace := range namespaces {
			writeConfig(&buffer,namespace)
		}
	}

	buffer.WriteString("</html>")

	rw.Write(buffer.Bytes())
}

func writeConfig(buffer *bytes.Buffer, namespace string) {
	buffer.WriteString(fmt.Sprintf("NamespaceName : %s <br/>", namespace))
	buffer.WriteString("Configurations: <br/>")
	cache := agollo.GetConfigCache(namespace)
	it := cache.NewIterator()
	for i := 0; i < int(cache.EntryCount()); i++ {
		entry := it.Next()
		if entry == nil {
			continue
		}
		buffer.WriteString(fmt.Sprintf("key : %s , value : %s <br/>", string(entry.Key), string(entry.Value)))
	}
}

type DefaultLogger struct {

}

func (this *DefaultLogger)Debugf(format string, params ...interface{})  {
	this.Debug(format,params)
}

func (this *DefaultLogger)Infof(format string, params ...interface{}) {
	this.Debug(format,params)
}


func (this *DefaultLogger)Warnf(format string, params ...interface{}) error {
	this.Debug(format,params)
	return nil
}

func (this *DefaultLogger)Errorf(format string, params ...interface{}) error {
	this.Debug(format,params)
	return nil
}


func (this *DefaultLogger)Debug(v ...interface{}) {
	fmt.Println(v)
}
func (this *DefaultLogger)Info(v ...interface{}){
	this.Debug(v)
}

func (this *DefaultLogger)Warn(v ...interface{}) error{
	this.Debug(v)
	return nil
}

func (this *DefaultLogger)Error(v ...interface{}) error{
	this.Debug(v)
	return nil
}
