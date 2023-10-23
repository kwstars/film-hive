module github.com/kwstars/film-hive

go 1.19

require (
	github.com/confluentinc/confluent-kafka-go/v2 v2.2.0
	github.com/envoyproxy/protoc-gen-validate v1.0.0
	github.com/go-kratos/kratos/contrib/config/nacos/v2 v2.0.0-20230309161611-7eca8f803405
	github.com/go-kratos/kratos/contrib/registry/nacos/v2 v2.0.0-20230309161611-7eca8f803405
	github.com/go-kratos/kratos/v2 v2.6.2
	github.com/go-redis/cache/v9 v9.0.0
	github.com/go-redsync/redsync/v4 v4.8.1
	github.com/golang/protobuf v1.5.3
	github.com/google/wire v0.5.0
	github.com/nacos-group/nacos-sdk-go v1.0.9
	github.com/pkg/errors v0.9.1
	github.com/redis/go-redis/extra/redisotel/v9 v9.0.2
	github.com/redis/go-redis/v9 v9.0.4
	github.com/stretchr/testify v1.8.2
	go.opentelemetry.io/otel v1.15.1
	go.opentelemetry.io/otel/exporters/jaeger v1.15.1
	go.opentelemetry.io/otel/sdk v1.15.1
	go.uber.org/automaxprocs v1.5.2
	go.uber.org/zap v1.24.0
	google.golang.org/genproto/googleapis/api v0.0.0-20230525234035-dd9d682886f9
	google.golang.org/grpc v1.57.0
	google.golang.org/protobuf v1.30.0
	gopkg.in/Shopify/sarama.v1 v1.20.1
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/mysql v1.5.2
	gorm.io/gorm v1.25.2-0.20230530020048-26663ab9bf55
	gorm.io/plugin/opentelemetry v0.1.1
)

require (
	github.com/DataDog/zstd v1.5.2 // indirect
	github.com/Shopify/toxiproxy v2.1.4+incompatible // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.18 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/eapache/go-resiliency v1.3.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230111030713-bf00bc1b83b6 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/frankban/quicktest v1.14.4 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/go-errors/errors v1.0.1 // indirect
	github.com/go-kratos/aegis v0.2.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-playground/form/v4 v4.2.0 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/lestrrat/go-file-rotatelogs v0.0.0-20180223000712-d3151e2a480f // indirect
	github.com/lestrrat/go-strftime v0.0.0-20180220042222-ba3bf9c1d042 // indirect
	github.com/lufia/plan9stats v0.0.0-20230110061619-bbe2e5e100de // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/power-devops/perfstat v0.0.0-20221212215047-62379fc7944b // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/redis/go-redis/extra/rediscmd/v9 v9.0.2 // indirect
	github.com/shirou/gopsutil/v3 v3.23.2 // indirect
	github.com/tklauser/go-sysconf v0.3.11 // indirect
	github.com/tklauser/numcpus v0.6.0 // indirect
	github.com/toolkits/concurrent v0.0.0-20150624120057-a4371d70e3e3 // indirect
	github.com/vmihailenco/go-tinylfu v0.2.2 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.4 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	go.opentelemetry.io/otel/metric v0.36.0 // indirect
	go.opentelemetry.io/otel/trace v1.15.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230526161137-0005af68ea54 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230525234030-28d5490b6b19 // indirect
	gopkg.in/ini.v1 v1.51.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/sqlite v1.4.3 // indirect
)
