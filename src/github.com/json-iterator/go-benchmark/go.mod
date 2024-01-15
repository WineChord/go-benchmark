module github.com/json-iterator/go-benchmark

go 1.21.2

replace github.com/apache/thrift => git.apache.org/thrift.git v0.19.0

replace git.apache.org/thrift.git => github.com/apache/thrift v0.19.0

require (
	git.apache.org/thrift.git v0.0.0-00010101000000-000000000000
	github.com/buger/jsonparser v1.1.1
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.3
	github.com/json-iterator/go v1.1.12
	github.com/mailru/easyjson v0.7.7
)

require (
	github.com/josharian/intern v1.0.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
)
