package main

import (
	"fmt"
	"testing"
)

// rpc service interface
type RPCService interface {
	Reference() string // rpc service id or reference id
}

// callback interface for async
type AsyncCallBack interface {
	CallBack() // callback
}

type AService struct {
}

func (a *AService) Reference() string {
	return "hello"
}

func (a *AService) CallBack() {
	fmt.Println("kjjj")
	return
}

type BService struct {
}

func (a *BService) Reference() string {
	return "hello"
}

func TestName(t *testing.T) {
	var service RPCService = &AService{}
	validCallback(service)

	var bService RPCService = &BService{}
	validCallback(bService)
}

func validCallback(service RPCService) {
	if sv, ok := service.(AsyncCallBack); ok {
		fmt.Printf("v implements String(): \n") // note: sv, not v
		sv.CallBack()
	}
}
