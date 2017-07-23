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
	buffer.WriteString(fmt.Sprintf("AppId : %s  \n",config.AppId))
	buffer.WriteString(fmt.Sprintf("Cluster : %s \n",config.Cluster))
	buffer.WriteString(fmt.Sprintf("NamespaceName : %s \n",config.NamespaceName))
	buffer.WriteString(fmt.Sprintf("ReleaseKey : %s \n",config.ReleaseKey))

	buffer.WriteString("Configurations: \n")
	for key,value :=range config.Configurations {
		buffer.WriteString(fmt.Sprintf("key : %s , value : %s \n",key,value))
	}

	rw.Write(buffer.Bytes())
}
