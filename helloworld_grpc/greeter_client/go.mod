module google.golang.org/grpc/examples/helloworld/helloworld/client

go 1.15

replace google.golang.org/grpc/examples/helloworld/helloworld v0.0.0 => ../helloworld

require (
	github.com/golang/protobuf v1.5.2
	google.golang.org/grpc v1.41.0
	google.golang.org/grpc/examples/helloworld/helloworld v0.0.0
	google.golang.org/protobuf v1.27.1 // indirect
)
