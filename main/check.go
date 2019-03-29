package main

import (
	"net/http"
	"github.com/zouyx/agollo"
	"fmt"
	"bytes"
)

func main() {
	agollo.InitCustomConfig(func () (*agollo.AppConfig, error) {
		return &agollo.AppConfig{
			AppId:         "SampleApp",
			Cluster:       "dev",
			Ip:            "http://localhost:8180",
			NamespaceName: "application",
		}, nil
	})
	go agollo.Start()

	http.HandleFunc("/check",GetAllConfig)

	http.ListenAndServe("0.0.0.0:9000",nil)
}

func GetAllConfig(rw http.ResponseWriter,req *http.Request)  {
	config:=agollo.GetCurrentApolloConfig()
	cache:=agollo.GetApolloConfigCache()

	var buffer bytes.Buffer
	buffer.WriteString("<html>")
	buffer.WriteString("<meta http-equiv=\"refresh\" content=\"3\">")

	key:=req.URL.Query().Get("key")
	if key=="" {
		buffer.WriteString(fmt.Sprintf("AppId : %s  <br/>", config.AppId))
		buffer.WriteString(fmt.Sprintf("Cluster : %s <br/>", config.Cluster))
		buffer.WriteString(fmt.Sprintf("NamespaceName : %s <br/>", config.NamespaceName))
		buffer.WriteString(fmt.Sprintf("ReleaseKey : %s <br/>", config.ReleaseKey))

		buffer.WriteString("Configurations: <br/>")
		it := cache.NewIterator()
		for i := 0; i < int(cache.EntryCount()); i++ {
			entry := it.Next()
			if entry == nil {
				continue
			}
			buffer.WriteString(fmt.Sprintf("key : %s , value : %s <br/>", string(entry.Key), string(entry.Value)))
		}
	}else{
		if config!=nil&&cache.EntryCount()>0{
			value,err:=cache.Get([]byte(key))

			if err!=nil {
				buffer.WriteString(fmt.Sprintf("get key : %s fail, error:%s <br/>", key,err.Error()))
			}else{
				buffer.WriteString(fmt.Sprintf("key : %s , value : %s <br/>", key, string(value)))
			}
		}else{
			buffer.WriteString("no value in cache!")
		}
	}

	buffer.WriteString("</html>")

	rw.Write(buffer.Bytes())
}
