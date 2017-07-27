package main

import (
	"net/http"
	"github.com/zouyx/agollo"
	"fmt"
	"bytes"
)

func main() {
	go agollo.Start()

	http.HandleFunc("/check",GetAllConfig)

	http.ListenAndServe("0.0.0.0:9000",nil)
}

func GetAllConfig(rw http.ResponseWriter,req *http.Request)  {
	config:=agollo.GetCurrentApolloConfig()

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
		for key, value := range config.Configurations {
			buffer.WriteString(fmt.Sprintf("key : %s , value : %s <br/>", key, value))
		}
	}else{
		if config!=nil&&config.Configurations!=nil{
			buffer.WriteString(fmt.Sprintf("key : %s , value : %s <br/>", key, config.Configurations[key]))
		}
	}

	buffer.WriteString("</html>")

	rw.Write(buffer.Bytes())
}
