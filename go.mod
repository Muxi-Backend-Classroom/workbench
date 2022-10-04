module github.com/Muxi-Backend-Classroom/workbench

replace workbench => ./

go 1.16

require (
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/go-micro v1.18.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/smartystreets/goconvey v1.7.2
	github.com/spf13/viper v1.12.0
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	go.uber.org/zap v1.23.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	workbench v0.0.0-00010101000000-000000000000
)

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	golang.org/x/net v0.0.0-20220907135653-1e95f45603a7 // indirect
	golang.org/x/sys v0.0.0-20220908164124-27713097b956 // indirect
	google.golang.org/genproto v0.0.0-20220908141613-51c1cc9bc6d0 // indirect
	google.golang.org/grpc v1.49.0 // indirect
)
