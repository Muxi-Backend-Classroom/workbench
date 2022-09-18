module workbench-user

//replace workbench-project => ../project

replace workbench => ../../

go 1.16

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/micro/go-micro => github.com/Lofanmi/go-micro v1.16.1-0.20210804063523-68bbf601cfa4 // to go 1.16

require (
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/kubernetes v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/wrapper/trace/opentracing v0.0.0-20200119172437-4fe21aa238fd
	github.com/opentracing/opentracing-go v1.2.0
	github.com/satori/go.uuid v1.2.0
	github.com/smartystreets/goconvey v1.7.2
	github.com/spf13/viper v1.12.0
	golang.org/x/net v0.0.0-20220826154423-83b083e8dc8b // indirect
	workbench v0.0.0-00010101000000-000000000000
)

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/nats-io/nats.go v1.9.1 // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	google.golang.org/protobuf v1.28.0
)
