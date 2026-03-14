module github.com/zouyx/agollo_demo

require (
	github.com/apolloconfig/agollo/v4 v4.4.1-0.20200101000000-067e04afcfcc
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575
)

replace github.com/apolloconfig/agollo/v4 => github.com/apolloconfig/agollo/v4 v4.4.0
replace github.com/golang/protobuf => github.com/golang/protobuf v1.5.2

go 1.13
