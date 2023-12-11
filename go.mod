module github.com/ahmad-ibra/count-client

go 1.21.3

require (
	buf.build/gen/go/ahmad-ibrahim/count-api/connectrpc/go v1.13.0-20231211181708-5a11761623de.1
	buf.build/gen/go/ahmad-ibrahim/count-api/protocolbuffers/go v1.31.0-20231211181708-5a11761623de.2
	connectrpc.com/connect v1.13.0
)

require google.golang.org/protobuf v1.31.0 // indirect
